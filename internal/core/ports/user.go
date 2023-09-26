package ports

import "deluze/internal/core/domain"

type UserRepository interface {
	Create(email string, username string) (*domain.User, error)
	Find(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindMany(page int, limit int) ([]*domain.User, error)
	Update(id string, username *string) (*domain.User, error)
}

type UserService interface {
	Create(email string) (*domain.User, error)
	Find(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindMany(page *int, limit *int) ([]*domain.User, error)
	Update(id string, username *string) (*domain.User, error)
}
