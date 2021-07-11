package domain

import (
	"errors"
	"fmt"
	"net/smtp"

	"github.com/muhfaris/request"
	"golang.org/x/net/proxy"
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

	var valid bool
	var errorSTMP error
	for _, mx := range mxs {
		if valid {
			continue
		}

		// create a socks5 dialer
		dialer, err := proxy.SOCKS5("tcp", "88.198.24.108:1080", nil, proxy.Direct)
		if err != nil {
			return &SMTPValidate{}, err
		}

		host := fmt.Sprintf("%s:25", mx)
		conn, err := dialer.Dial("tcp", host)
		if err != nil {
			return &SMTPValidate{}, err
		}

		// Connect to the remote SMTP server.
		//c, err := smtp.Dial(host)
		c, err := smtp.NewClient(conn, mx)
		if err != nil {
			errorSTMP = err
			continue
		}

		// Set the sender and recipient first
		if err = c.Mail("info@example.com"); err != nil {
			errorSTMP = err
			continue
		}

		if err = c.Rcpt(e.Email); err != nil {
			errorSTMP = err
			continue
		}

		// Send the QUIT command and close the connection.
		err = c.Quit()
		if err != nil {
			errorSTMP = err
			continue
		}

		valid = true
		errorSTMP = nil
	}

	if errorSTMP != nil {
		return smtpValidate, errorSTMP
	}

	smtpValidate.IsValid = valid
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
