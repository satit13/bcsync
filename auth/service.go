package auth

import "context"

type (
	keyToken struct{}
)

// Service is the auth service
type Service interface {
	GetToken(tokenID string) (*Token, error)
}

// NewService creates new auth service
func NewService(auths Repository) (Service, error) {
	s := service{auths}
	return &s, nil
}

// GetAccountID returns account id from context
func GetAccountID(ctx context.Context) int64 {
	x, ok := ctx.Value(keyToken{}).(*Token)
	if !ok {
		return -1
	}
	return x.AccountID
}

// GetVendingID returns vending id from context

func GetTokenID(ctx context.Context) string {
	x, ok := ctx.Value(keyToken{}).(*Token)
	if !ok {
		return ""
	}
	return x.TokenID
}

type service struct {
	auths Repository
}

func (s *service) GetToken(tokenID string) (*Token, error) {
	tk, err := s.auths.GetToken(tokenID)

	if err != nil {
		return nil, err
	}
	return tk, nil
}
