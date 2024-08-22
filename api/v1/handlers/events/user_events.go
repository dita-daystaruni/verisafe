package events

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/models"
)

func EmitUserCreated(user *models.User, cfg *configs.Config) {
	userData, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error: Failed to marshal user data: %s\n", err.Error())
		return
	}

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	for _, url := range cfg.EventConfig.UserCreateEvent {
		wg.Add(1) // Increment the wait group counter
		go func(url string) {
			defer wg.Done() // Decrement the counter when the goroutine completes

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(userData))
			if err != nil {
				log.Printf("Error: Failed to create request for %s: %s\n", url, err.Error())
				return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("api-key", cfg.APISecrets.EventApiSecret)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				bodyString := string(bodyBytes)
				log.Printf("Error: Failed to send request to %s: %s\n\n%s\n", url, err.Error(), bodyString)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				log.Printf("Error: Received non-Created response from %s: %s\n", url, resp.Status)
			} else {
				log.Printf("Successfully sent request to %s\n", url)
			}
		}(url)
	}

	wg.Wait() // Wait for all goroutines to finish
}
