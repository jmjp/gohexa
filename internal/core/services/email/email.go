package EmailService

import "deluze/internal/core/ports"

type EmailService struct {
	userRepo  ports.UserRepository
	emailRepo ports.EmailRepository
}

func New(userRepo ports.UserRepository, emailRepo ports.EmailRepository) *EmailService {
	return &EmailService{
		userRepo:  userRepo,
		emailRepo: emailRepo,
	}
}

func (s *EmailService) SendEmail(user string, message string, subject string) error {
	usr, err := s.userRepo.Find(user)
	if err != nil {
		return err
	}
	return s.emailRepo.SendEmail(usr.Email, message, subject)
}
