// Defines utilities and functionality for password security

package security

import (
	"golang.org/x/crypto/bcrypt"
)

type passwordValidator func(string) bool

type passwordPolicy struct {
	Validations []passwordValidator
}

// Runs all password validations on the given password
// Returns true if all validations pass, false otherwise
func (policy *passwordPolicy) PasswordValid(password string) bool {
	for _, validator := range policy.Validations {
		if !validator(password) {
			return false
		}
	}
	return true
}

var (
	minimumLength  = 6
	PasswordPolicy = &passwordPolicy{
		Validations: []passwordValidator{
			meetsMinLength,
		},
	}
)

/*
 * Password Validation Functions
 */

func meetsMinLength(password string) bool {
	return len(password) >= minimumLength
}

/*
 * Functions for securely handling/storing passwords
 */

// Returns a cryptographically secure hash of the given password
func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err // TODO: Better error handling
	}
	return string(hash[:]), nil
}

// Confirms whether the given password matches the expected password
// for the user
func ConfirmPassword(passwordHash string, password string) bool {
	passwordBytes := []byte(password)
	storedHash := []byte(passwordHash)
	return bcrypt.CompareHashAndPassword(storedHash, passwordBytes) == nil
}
