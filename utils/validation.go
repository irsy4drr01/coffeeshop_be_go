package utils

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

var (
	upperRegex   = regexp.MustCompile(`[A-Z]`)
	lowerRegex   = regexp.MustCompile(`[a-z]`)
	digitRegex   = regexp.MustCompile(`[0-9]`)
	specialRegex = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-\=\[\]\{\};':"\\|,.<>\/?]`)
)

// ValidateEmailFormat uses net/mail for robust check
func ValidateEmailFormat(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

// ValidatePasswordStrength checks for upper, lower, digit, special
func ValidatePasswordStrength(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if !upperRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !lowerRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !digitRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one number")
	}
	if !specialRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one special character")
	}
	return nil
}

func ConvertValidatorError(raw string) string {
	raw = strings.ToLower(raw)
	switch {
	case strings.Contains(raw, "email"):
		return "email is required"
	case strings.Contains(raw, "password"):
		return "password is required"
	case strings.Contains(raw, "fullname"):
		return "fullname is required"
	default:
		return "validation failed. please check your input"
	}
}
