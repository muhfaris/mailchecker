package domain

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"

	"github.com/muhfaris/request"
)

// SMTPValidate is wrap smtp validation
type SMTPValidate struct {
	IsValid bool   `json:"is_valid"`
	Host    string `json:"-"`
	Email   string `json:"-"`
}

// EmailSMTPValidator is validate smtp server
func EmailSMTPValidator(mxs []string, e *EmailVerifier) (*SMTPValidate, error) {
	var smtpValidate = &SMTPValidate{Host: e.Domain, Email: e.Email}

	mx := mxs[0]
	// Connect to the remote SMTP server.
	host := fmt.Sprintf("%s:25", mx)
	c, err := smtp.Dial(host)
	if err != nil {
		return smtpValidate, err
	}

	// Set the sender and recipient first
	if err := c.Mail("info@example.com"); err != nil {
		return smtpValidate, err
	}

	if err := c.Rcpt(e.Email); err != nil {
		log.Println(err)
		return smtpValidate, err
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		return smtpValidate, err
	}

	return smtpValidate, nil
}

// MailValidate is other way to validate email
func (smtp *SMTPValidate) MailValidate() error {
	switch smtp.Host {
	case "gmail.com":
		return smtp.GmailValidate()

	default:
		return nil
	}
}

// GmailValidate is validate gmail
func (smtp *SMTPValidate) GmailValidate() error {
	//const baseURL = "https://mail.google.com/mail/gxlu?email=%s"
	const baseURL = "https://mail.google.com/mail/gxlu"

	app := request.ReqApp{
		URL: baseURL,
		QueryString: map[string]string{
			"email": smtp.Email,
		},
	}

	resp, err := app.GET()
	if err != nil {
		return err
	}

	value := resp.HTTP.Header.Get("Set-Cookie")
	if value == "" {
		return errors.New("gmail account not found")
	}

	return nil
}
