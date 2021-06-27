package domain

import (
	"fmt"

	"github.com/muhfaris/mailchecker/gateway/structures"
)

const (
	ValidEmailStatus   = "valid"
	InvalidEmailStatus = "invalid"
	UnknownEmailStatus = "unknown"
)

// EmailStatus is available email status
var EmailStatus = map[string]string{
	"valid":   ValidEmailStatus,
	"invalid": InvalidEmailStatus,
	"unknown": UnknownEmailStatus,
}

// EmailVerifier is wrap data of email verifier
type EmailVerifier struct {
	Email      string `json:"email"`
	Status     string `json:"status"`
	Username   string `json:"username"`
	Domain     string `json:"domain"`
	Disposable bool   `json:"is_disposable"`
	DNSMX      bool   `json:"is_dns_mx"`
	Message    string `json:"message"`
}

// CreateEmailVerifier is create new object of email verifier
func CreateEmailVerifier(p structures.EmailVerifierRead) (*EmailVerifier, error) {
	if p.Email == "" {
		return &EmailVerifier{}, fmt.Errorf(ErrorMissingParam, "email")
	}

	return &EmailVerifier{
		Email:  p.Email,
		Status: InvalidEmailStatus,
	}, nil
}

// ChangeStatus is change status email
func (e *EmailVerifier) ChangeStatus(status string) error {
	if status == "" {
		return fmt.Errorf(ErrorMissingParam, "status")
	}

	s, ok := EmailStatus[status]
	if !ok {
		return fmt.Errorf(ErrorInvalidParam, "status")
	}

	e.Status = s
	return nil
}

// ChangeUsername is change username
func (e *EmailVerifier) ChangeUsername(username string) error {
	if username == "" {
		return fmt.Errorf(ErrorMissingParam, "username")
	}

	e.Username = username
	return nil
}

// ChangeDomain is change domain email
func (e *EmailVerifier) ChangeDomain(domain string) error {
	if domain == "" {
		return fmt.Errorf(ErrorMissingParam, "domain")
	}

	e.Domain = domain
	return nil
}

// ChangeDisposable is change disposable status
func (e *EmailVerifier) ChangeDisposable(status bool) {
	e.Disposable = status
}

// ChangeDNSMX is change dns mx status
func (e *EmailVerifier) ChangeDNSMX(status bool) {
	e.DNSMX = status
}

// ChangeMessage is change message email
func (e *EmailVerifier) ChangeMessage(message string) error {
	if message == "" {
		return fmt.Errorf(ErrorMissingParam, "message")
	}

	e.Message = message
	return nil
}

// Valid is validate email
func (e *EmailVerifier) Valid() error {
	e.ChangeStatus(InvalidEmailStatus)
	emailSyntax, err := EmailSyntaxValidator(e.Email)
	if err != nil {
		e.ChangeMessage("email addres in invalid format")
		return nil
	}

	e.ChangeUsername(emailSyntax.User)
	e.ChangeDomain(emailSyntax.Domain)

	isDisposable := EmailDisposableValidator(e.Email)
	if isDisposable {
		e.ChangeDisposable(true)
		e.ChangeMessage("disposable email address (temporary email address)")
	}

	mxs, isMX := EmailMXValidator(e.Domain)
	if isMX {
		e.ChangeDNSMX(true)
	}

	if !isMX {
		e.ChangeStatus(InvalidEmailStatus)
		e.ChangeMessage("mail delivery MX host not found")
	}

	if !isDisposable && isMX {
		e.ChangeStatus(ValidEmailStatus)
		e.ChangeMessage("email address is valid")
	}

	if err := EmailSMTPValidator(mxs, e.Email); err != nil {
		e.ChangeStatus(InvalidEmailStatus)
		e.ChangeMessage("email address doesn't exists")
	}

	return nil
}
