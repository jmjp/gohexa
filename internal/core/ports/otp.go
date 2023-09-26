package ports

import (
	"deluze/internal/core/domain"
	"time"
)

type OtpRepository interface {
	Create(code string, user string, expiresAt time.Time) (*domain.Otp, error)
	Find(code string, user string) (*domain.Otp, error)
	Delete(id string) error
}
