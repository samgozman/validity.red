package main

// Mailer is an interface for sending emails.
type Mailer interface {
	SendEmailVerification(email, tokenURL string) error
	// SendPasswordReset(email, token string) error
	// TODO: Send email with "how to use" instructions
}
