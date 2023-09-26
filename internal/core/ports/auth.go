package ports

import "deluze/internal/core/domain"

type AuthService interface {
	Login(email string) error
	Verify(email string, code string) (*domain.User, *string, error)
	Refresh(token string) (*string, error)
}
