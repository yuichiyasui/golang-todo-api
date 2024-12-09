package infrastructure

import (
	"api/domain"
	"fmt"
)

type Mailer struct{}

func NewMailer() domain.MailerInterface {
	return &Mailer{}
}

func (m *Mailer) SendEmail(to string, subject string, body string) error {
	// MEMO: 実際にはメールを送信する
	fmt.Printf("Send email to: %s\n", to)
	fmt.Printf("Subject: %s\n", subject)
	fmt.Printf("Body: %s\n", body)
	return nil
}
