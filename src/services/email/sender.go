package email

import (
	"fmt"
	"log"

	"github.com/csvitor-dev/social-media/internal/config/env"
	"github.com/csvitor-dev/social-media/types"
	"github.com/resend/resend-go/v2"
)

func SendEmailForPasswordReset(email types.Email, token string) error {
	resetLink := fmt.Sprintf("%s/auth/reset-password?token=%s", env.ApiEnv.WEB_URL, token)

	email.Body = fmt.Sprintf(`
		<h2>Recuperação de senha</h2>
		<p>Você solicitou a redefinição de senha. Clique no botão abaixo para continuar:</p>
		<p>
			<a
				href="%s"
				style="cursor:pointer,display:inline-block;padding:10px 20px;background:#4CAF50;color:#fff;text-decoration:none;border-radius:5px;"
			>Redefinir Senha</a>
		</p>
		<p>Se não foi você que solicitou, apenas ignore este email.</p>
	`, resetLink)

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

	return err
}
