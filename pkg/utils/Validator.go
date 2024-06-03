package utils

import (
	"errors"
	"regexp"
)

// The ValidateAdmissionNumber ensures that the supplied input matches
// the daystar admission number
func ValidateAdmissionNumber(input string) error {
	pattern := `^\d{2}-\d{4}$`

	re := regexp.MustCompile(pattern)

	if !re.MatchString(input) {
		return errors.New("Please check your admission number and try that again")
	}
	return nil
}
