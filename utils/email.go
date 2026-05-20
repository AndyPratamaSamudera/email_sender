package utils

import (
	"errors"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

var ErrTimeout = errors.New("timeout when sending email")

// AutoDetectSMTPHost detects the SMTP host based on the email domain.
func AutoDetectSMTPHost(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return ""
	}
	domain := parts[1]

	switch domain {
	case "gmail.com":
		return "smtp.gmail.com"
	case "outlook.com", "hotmail.com":
		return "smtp.office365.com"
	case "yahoo.com":
		return "smtp.mail.yahoo.com"
	default:
		// Return empty string if domain is not known
		return ""
	}
}

// SendEmail sends an HTML email via SMTP with timeout.
func SendEmail(sender, password, recipient, subject, body, smtpHost string, smtpPort int) error {
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, sender, password)

	// Use channel for timeout handling (simple implementation)
	errCh := make(chan error, 1)

	go func() {
		errCh <- d.DialAndSend(m)
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(15 * time.Second): // 15 seconds timeout
		return ErrTimeout
	}
}
