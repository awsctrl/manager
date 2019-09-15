package token

import (
	uuid "github.com/satori/go.uuid"
)

type Token interface {
	// Generate creates a random string
	Generate() string
}

type token struct{}

// New initialized a new token
func New() token {
	return token{}
}

// Generate will generate a UUID V4 string
func (t token) Generate() string {
	return uuid.NewV4().String()
}
