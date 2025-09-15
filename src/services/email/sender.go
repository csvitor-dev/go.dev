package email

import (
	"fmt"
	"log"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	"github.com/csvitor-dev/go.dev/src/views"
	"github.com/csvitor-dev/go.dev/types"
	"github.com/resend/resend-go/v2"
)

func SendEmailForPasswordReset(email types.Email, token string) error {
	resetLink := fmt.Sprintf("%s/auth/reset-password?token=%s", env.ApiEnv.WEB_URL, token)

	content, err := views.Get(
		views.ViewOptions{
			View: "email.recover-password",
			Data: map[string]any{
				"Link": resetLink,
			},
		},
	)

	if err != nil {
		return err
	}
	email.Body = content

	client := resend.NewClient(env.EmailEnv.API_KEY)
	params := &resend.SendEmailRequest{
		From:    env.EmailEnv.EMAIL,
		To:      []string{email.To},
		Html:    email.Body,
		Subject: email.Subject,
	}
	response, err := client.Emails.Send(params)

	if err != nil {
		return err
	}
	log.Printf("Email sent: %s", response.Id)

	return nil
}
