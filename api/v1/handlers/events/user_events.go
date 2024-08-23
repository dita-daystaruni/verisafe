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
				cfg.Logger.Errorf("Error: Failed to create request for %s: %s\n", url, err.Error())
				return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("api-key", cfg.APISecrets.EventApiSecret)

			// Perform Request
			client := &http.Client{}
			resp, err := client.Do(req)
			// handle sending errors
			if err != nil {
				cfg.Logger.Errorf("Failed to send request with error: %s\n", err.Error())
				return
			}

			// Check the reponse
			if resp.StatusCode != http.StatusCreated {
				// Read the body bytes for logging
				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					cfg.Logger.Error(err)
				}
				bodyString := string(bodyBytes)
				cfg.Logger.Error("Error: Failed to send request to %s: %s\n\n%s\n", url, err.Error(), bodyString)

			} else {
				cfg.Logger.Infof("Successfully sent user created event to %s for user %s\n", url, user.Username)
			}

		}(url)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func EmitUserUpdated(user *models.User, cfg *configs.Config) {
	userData, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error: Failed to marshal user data: %s\n", err.Error())
		return
	}

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	for _, url := range cfg.EventConfig.UserUpdatedEvent {
		wg.Add(1) // Increment the wait group counter
		go func(url string) {
			defer wg.Done() // Decrement the counter when the goroutine completes

			req, err := http.NewRequest("PATCH", url+user.ID.String(), bytes.NewBuffer(userData))
			if err != nil {
				cfg.Logger.Errorf("Error: Failed to create request for %s: %s\n", url, err.Error())
				return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("api-key", cfg.APISecrets.EventApiSecret)

			client := &http.Client{}
			resp, err := client.Do(req)

			if err != nil {
				cfg.Logger.Errorf(err.Error())
			}

			if resp.StatusCode != http.StatusOK {
				cfg.Logger.Error("Error: Received non-Created response from %s: %s\n", url, resp.Status)
			} else {
				cfg.Logger.Info("Successfully sent request to %s to update student\n", url)
			}
		}(url)
	}

	wg.Wait() // Wait for all goroutines to finish

}

func EmitUserDeleted(user *models.User, cfg *configs.Config) {
	userData, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error: Failed to marshal user data: %s\n", err.Error())
		return
	}

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	for _, url := range cfg.EventConfig.UserUpdatedEvent {
		wg.Add(1) // Increment the wait group counter
		go func(url string) {
			defer wg.Done() // Decrement the counter when the goroutine completes

			req, err := http.NewRequest("DELETE", url+user.ID.String(), bytes.NewBuffer(userData))
			if err != nil {
				cfg.Logger.Errorf("Error: Failed to create request for %s: %s\n", url, err.Error())
				return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("api-key", cfg.APISecrets.EventApiSecret)

			client := &http.Client{}
			resp, err := client.Do(req)

			if err != nil {
				cfg.Logger.Errorf(err.Error())
			}

			if resp.StatusCode != http.StatusOK {
				cfg.Logger.Error("Error: Received non-Created response from %s: %s\n", url, resp.Status)
			} else {
				cfg.Logger.Info("Successfully sent request to %s to update student\n", url)
			}
		}(url)
	}

	wg.Wait() // Wait for all goroutines to finish

}
