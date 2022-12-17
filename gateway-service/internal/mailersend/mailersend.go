// Package mailersend is used to send emails via MailerSend 3rd party service.
package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	ms "github.com/mailersend/mailersend-go"
)

// MailerSend integration.
type MailerSend struct {
	APIKey string
}

// SendEmailVerification sends email verification email to the user with the tokenURL
// via MailerSend. The tokenURL is a link to the frontend with the token as a query param.
//
// TokenURL should be a full URL, e.g. https://validity.red/verify?token=123
func (m *MailerSend) SendEmailVerification(email, tokenURL string) error {
	const requestTimeout = 5 * time.Second

	client := ms.NewMailersend(m.APIKey)
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	subject := "Confirm your email | Validity.Red"
	recipientName := strings.Split(email, "@")[0]

	from := ms.From{
		Name:  "Validity.Red",
		Email: "noreply@validity.red",
	}

	recipients := []ms.Recipient{
		{
			Name:  recipientName,
			Email: email,
		},
	}

	personalization := []ms.Personalization{
		{
			Email: email,
			Data: map[string]interface{}{
				"name":             recipientName,
				"confirmation_url": tokenURL,
			},
		},
	}

	message := client.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetTemplateID("3zxk54vm7ox4jy6v")
	message.SetPersonalization(personalization)

	res, err := client.Email.Send(ctx, message)
	if err != nil {
		ef := fmt.Errorf("error sending email to '%s': %w", email, err)
		sentry.CaptureException(ef)

		return ef
	}

	if res.StatusCode != http.StatusAccepted {
		ef := fmt.Errorf("error sending email to '%s': %w", email, err)
		sentry.CaptureException(ef)

		return ef
	}

	return nil
}
