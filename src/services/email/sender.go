package email

import (
	"fmt"

	"github.com/csvitor-dev/social-media/internal/config"
	"github.com/csvitor-dev/social-media/types"
	"github.com/resend/resend-go/v2"
)

func SendEmailForPasswordReset(email types.Email, token string) error {
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", config.ApiEnv.WEB_URL, token)

	email.Body = fmt.Sprintf(`
		<h2>Recuperação de senha</h2>
		<p>Você solicitou a redefinição de senha. Clique no botão abaixo para continuar:</p>
		<p><a href="%s" style="display:inline-block;padding:10px 20px;background:#4CAF50;color:#fff;text-decoration:none;border-radius:5px;">Redefinir Senha</a></p>
		<p>Se não foi você que solicitou, apenas ignore este email.</p>
	`, resetLink)

	client := resend.NewClient(config.EmailEnv.API_KEY)
	params := &resend.SendEmailRequest{
		From:    config.EmailEnv.EMAIL,
		To:      []string{email.To},
		Html:    email.Body,
		Subject: email.Subject,
	}
	_, err := client.Emails.Send(params)

	if err != nil {
		return err
	}
	return nil
}
