package mailersend

import (
	"context"
	"fmt"
	"strings"
	"time"

	ms "github.com/mailersend/mailersend-go"
)

// MailerSend integration

type MailerSend struct {
	APIKey string
}

func (m *MailerSend) SendEmailVerification(email, tokenUrl string) error {
	client := ms.NewMailersend(m.APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
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
				"confirmation_url": tokenUrl,
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
		return fmt.Errorf("error sending email to '%s': %s", email, err)
	}
	if res.StatusCode != 202 {
		return fmt.Errorf("error sending email to '%s': %s", email, res.Body)
	}

	return nil
}
