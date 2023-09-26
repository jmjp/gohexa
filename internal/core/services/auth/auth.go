package AuthService

import (
	"deluze/internal/core/domain"
	"deluze/internal/core/ports"
	UserService "deluze/internal/core/services/user"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type AuthService struct {
	userRepo  ports.UserRepository
	otpRepo   ports.OtpRepository
	emailRepo ports.EmailRepository
}

// Refresh implements ports.AuthService.
func (*AuthService) Refresh(token string) (*string, error) {
	panic("unimplemented")
}

func New(userRepo ports.UserRepository, otpRepo ports.OtpRepository, emailRepo ports.EmailRepository) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		otpRepo:   otpRepo,
		emailRepo: emailRepo,
	}
}

func (s *AuthService) Login(email string) error {
	user, err := s.userRepo.FindByEmail(email)
	if user == nil {
		usr, err := s.userRepo.Create(email, UserService.GenerateAnimalsUsername())
		if err != nil {
			return err
		}
		user = usr
	}
	otp, err := s.otpRepo.Create(generateOtp(), user.Id, time.Now().Add(time.Minute*1))
	if err != nil {
		return err
	}
	err = s.emailRepo.SendEmail(user.Email, "otp code", otp.Code)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Verify(email string, code string) (*domain.User, *string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, nil, err
	}
	otp, err := s.otpRepo.Find(code, user.Id)
	if err != nil {
		return nil, nil, err
	}
	if otp.ExpiredAt.After(time.Now()) {
		return nil, nil, errors.New("verification code expired")
	}
	token, err := GenerateToken(user.Id, time.Hour*24)
	if err != nil {
		return nil, nil, err
	}
	err = s.otpRepo.Delete(otp.Id)
	if err != nil {
		return nil, nil, err
	}
	return user, token, nil
}

func generateOtp() string {
	rand.NewSource(time.Now().UnixNano())
	// Gerar um número de 6 dígitos aleatório
	codigo := fmt.Sprintf("%06d", rand.Intn(1000000))
	return codigo
}
