package services

import (
	"context"

	"github.com/muhfaris/mailchecker/app/email/domain"
	"github.com/muhfaris/mailchecker/gateway/structures"
)

// EmailVerifier is email verification service
type EmailVerifier struct{}

// NewEmailVerifier is create new service email verification
func NewEmailVerifier() *EmailVerifier {
	return &EmailVerifier{}
}

// Validate is validate email
func (service *EmailVerifier) Validate(c context.Context, p structures.EmailVerifierRead) (*domain.EmailVerifier, error) {
	email, err := domain.CreateEmailVerifier(p)
	if err != nil {
		return &domain.EmailVerifier{}, err
	}

	if err := email.Valid(); err != nil {
		return &domain.EmailVerifier{}, err
	}

	return email, nil
}
