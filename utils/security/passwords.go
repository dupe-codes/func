// Defines utilities and functionality for password security

package security

import (
    "golang.org/x/crypto/bcrypt"

    "github.com/njdup/serve/users"
)

// Returns a cryptographically secure hash of the given password
func HashPassword(password string) string {
    passwordBytes := []byte(password)
    hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
    if err != nil {
        panic(err) // TODO: Better error handling
    }
    return string(hash[:])
}

// Confirms whether the given password matches the expected password
// for the user
func ConfirmPassword(user *users.User, password string) bool {
    passwordBytes := []byte(password)
    storedHash := []byte(user.PasswordHash)
    return bcrypt.CompareHashAndPassword(storedHash, passwordBytes) == nil
}
