package auth

import "errors"

// Errors
var (
	ErrTokenNotFound = errors.New("auth: token not found")
)

// Token entity
type Token struct {
	ID        string
	AccountID int64
	TokenID   string
}

// Repository is the auth storage
type Repository interface {
	GetToken(tokenID string) (*Token, error)
}
