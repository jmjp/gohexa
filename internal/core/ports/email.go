package ports

type EmailRepository interface {
	SendEmail(to string, subject string, body string) error
}

type EmailService interface {
	SendEmail(user string, message string, subject string) error
}
