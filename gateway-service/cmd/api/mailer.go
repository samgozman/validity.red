package main

// Emails driver
type Mailer interface {
	SendEmailVerification(email, tokenUrl string) error
	// SendPasswordReset(email, token string) error
	// TODO: Send email with "how to use" instructions
}
