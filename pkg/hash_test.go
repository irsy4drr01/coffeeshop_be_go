package pkg

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "password123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if hashedPassword == password {
		t.Fatalf("Expected hashed password to be different from plain password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		t.Fatalf("Expected hashed password to match the original password, but got error: %v", err)
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "password123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = VerifyPassword(hashedPassword, password)
	if err != nil {
		t.Fatalf("Expected correct password verification to pass, but got error: %v", err)
	}

	incorrectPassword := "wrongpassword"
	err = VerifyPassword(hashedPassword, incorrectPassword)
	if err == nil {
		t.Fatalf("Expected incorrect password verification to fail, but it passed")
	}
}
