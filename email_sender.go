package main

// EmailSender defines the interface for sending emails
type EmailSender interface {
	SendEmail(recipient string, subject string, body string) error
}

// SMTPSender implements the EmailSender interface using SMTP
type SMTPSender struct{}

func (s SMTPSender) SendEmail(recipient string, subject string, body string) error {
	return nil
}
