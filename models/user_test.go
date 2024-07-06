package models

import (
	"testing"
	"time"
)

// Tests that the password hashing mechanism works
// as expected
func TestPasswordHashing(t *testing.T) {
	var password string = "hello"
	user := User{Password: password}

	// hash the password
	if err := user.HashPassword(); err != nil {
		t.Fatal("Failed to hash password")
	}

	if user.Password == password {
		t.Fatalf("Expected user password to be some random strings instead got %s\n", user.Password)
	}

	// Passord comparison
	ok, err := user.ComparePassword(password)
	if err != nil {
		t.Fatalf("An error occured while comparing passwords: %v\n", err)
	}

	if !ok {
		t.Fatalf("Expected passwords to match instead got %v\n", ok)
	}
}

func TestUserValidation(t *testing.T) {
	user := User{
		Username:    "user1",
		FirstName:   "John",
		LastName:    "Doe",
		Gender:      "male",
		Address:     "123 Main St",
		NationalID:  "1234567890",
		Email:       "john.doe@example.com",
		DateOfBirth: time.Now().AddDate(-25, 0, 0), // 25 years ago
		Password:    "testpass",
	}

	// Test a valid user
	err := user.Validate()
	if err != nil {
		t.Errorf("Validation failed for valid user: %v", err)
	}

	// Test invalid cases

	// Invalid username
	user.Username = "us"
	err = user.Validate()
	if err == nil {
		t.Error("Expected error for invalid username")
	}

	// Invalid email
	user.Username = "user1"
	user.Email = "invalidemail"
	err = user.Validate()
	if err == nil {
		t.Error("Expected error for invalid email")
	}

	// Invalid password
	user.Email = "john.doe@example.com"
	user.Password = "12345"
	err = user.Validate()
	if err == nil {
		t.Error("Expected error for invalid password")
	}

	// Invalid gender
	user.Password = "testpass"
	user.Gender = "unknown"
	err = user.Validate()
	if err == nil {
		t.Error("Expected error for invalid gender")
	}
}
