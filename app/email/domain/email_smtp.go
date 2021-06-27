package domain

import (
	"fmt"
	"log"
	"net"
	"net/smtp"
)

// EmailSMTPValidator is validate smtp server
func EmailSMTPValidator(mxs []*net.MX, email string) error {
	mx := mxs[0]
	// Connect to the remote SMTP server.
	host := fmt.Sprintf("%s:25", mx.Host)
	c, err := smtp.Dial(host)
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("info@example.com"); err != nil {
		return err
	}

	if err := c.Rcpt(email); err != nil {
		log.Println(err)
		return err
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		return err
	}

	return nil
}
