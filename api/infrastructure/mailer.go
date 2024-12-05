package infrastructure

import "api/domain"

type Mailer struct{}

func NewMailer() domain.MailerInterface {
	return &Mailer{}
}

func (m *Mailer) SendEmail(to string, subject string, body string) error {
	// TODO: メール送信処理を実装する
	return nil
}
