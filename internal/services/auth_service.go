package services

import (
	"crypto/rand"
	"fmt"
)

type AuthService struct{}

func (s *AuthService) GenerateSecret() (string, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", randomBytes), nil
}
