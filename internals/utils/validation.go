package utils

import (
	"errors"
	"regexp"
)

func ValidateEmail(str string) error {
	if len(str) < 5 {
		return errors.New("Email should be at least 5 characters long.")
	}
	// Basic email format validation using regex
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailPattern, str)
	if err != nil {
		return errors.New("Error while validating email format.")
	}
	if !matched {
		return errors.New("Enter a valid email address.")
	}
	return nil
}

func ValidatePassword(str string) error {
	if len(str) < 8 {
		return errors.New("Password should be at least 8 characters long.")
	}

	// Additional password complexity checks
	// hasNumber := regexp.MustCompile(`[0-9]`).MatchString
	// hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	// hasLower := regexp.MustCompile(`[a-z]`).MatchString
	// hasSpecial := regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString
	//
	// if !hasNumber(str) {
	// 	return errors.New("Password must contain at least one number.")
	// }
	// if !hasUpper(str) {
	// 	return errors.New("Password must contain at least one uppercase letter.")
	// }
	// if !hasLower(str) {
	// 	return errors.New("Password must contain at least one lowercase letter.")
	// }
	// if !hasSpecial(str) {
	// 	return errors.New("Password must contain at least one special character.")
	// }

	return nil
}

type ValidateFunc func(string) error

func ValidatePasswordMatch(password string) ValidateFunc {
	return func(s string) error {
		if password != s {
			return errors.New("Password don't match")
		}
		return ValidatePassword(s)
	}
}
