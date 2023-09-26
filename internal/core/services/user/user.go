package UserService

import (
	"deluze/internal/core/domain"
	"deluze/internal/core/ports"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
	repo ports.UserRepository
}

func New(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(email string) (*domain.User, error) {
	return s.repo.Create(email, GenerateAnimalsUsername())
}

func (s *UserService) Find(id string) (*domain.User, error) {
	return s.repo.Find(id)
}

func (s *UserService) FindByEmail(email string) (*domain.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) FindMany(page *int, limit *int) ([]*domain.User, error) {
	if page == nil {
		page = new(int)
		*page = 0
	}
	if limit == nil {
		limit = new(int)
		*limit = 10
	}
	return s.repo.FindMany(*page, *limit)
}

func (s *UserService) Update(id string, username *string) (*domain.User, error) {
	if username == nil {
		return s.Find(id)
	}
	return s.repo.Update(id, username)
}

func GenerateAnimalsUsername() string {
	animals := []string{"dog", "cat", "elephant", "lion", "tiger", "giraffe", "monkey", "dolphin", "penguin", "kangaroo"}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(animals))
	randomAnimal := animals[randomIndex]
	suffix := randomSuffix(6)
	username := fmt.Sprintf("%s_%s", randomAnimal, suffix)
	return username

}

func randomSuffix(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	suffix := make([]byte, length)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range suffix {
		suffix[i] = charset[rand.Intn(len(charset))]
	}
	return string(suffix)
}
