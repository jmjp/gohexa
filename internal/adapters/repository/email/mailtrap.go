package EmailRepository

import (
	"net/smtp"
	"os"
)

type EmailMailTrapRepository struct {
	smtp smtp.Auth
}

func NewEmailTrapRepository() *EmailMailTrapRepository {
	auth := smtp.PlainAuth("deluze", os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_HOST"))
	return &EmailMailTrapRepository{
		smtp: auth,
	}
}

func (r *EmailMailTrapRepository) SendEmail(to string, subject string, body string) error {
	msg := []byte("To:" + to + "\r\n" + "Subject:" + subject + "\r\n" + "\r\n" + body + "\r\n")
	err := smtp.SendMail(os.Getenv("EMAIL_HOST")+":587", r.smtp, "teste@teste.com", []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}
