package helper

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(hashedPassword, inputPassword string) error {
	fmt.Println("here", hashedPassword, inputPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		return err // Passwords don't match
	}
	return nil // Passwords match
}

func GenerateSecurePassword(length int) (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>/?`~"

	password := make([]byte, length)
	_, err := rand.Read(password)
	if err != nil {
		return "", err
	}

	// Map the random bytes to the characters defined in 'chars'
	for i := range password {
		password[i] = chars[int(password[i])%len(chars)]
	}
	fmt.Println("GP", password)
	return string(password), nil
}