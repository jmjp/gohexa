package UserRepository

import (
	"context"
	"deluze/ent"
	"deluze/internal/core/domain"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type UserPostgresRepository struct {
	db *ent.Client
}

// Find implements ports.UserRepository.
func (*UserPostgresRepository) Find(id string) (*domain.User, error) {
	panic("unimplemented")
}

// FindByEmail implements ports.UserRepository.
func (*UserPostgresRepository) FindByEmail(email string) (*domain.User, error) {
	panic("unimplemented")
}

// FindMany implements ports.UserRepository.
func (r *UserPostgresRepository) FindMany(page int, limit int) ([]*domain.User, error) {
	usrs, err := r.db.User.Query().Limit(limit).Offset(page).All(context.Background())
	if err != nil {
		return nil, err
	}
	var usrsDomain []*domain.User
	for _, usr := range usrs {
		user := &domain.User{
			Id:        usr.ID.String(),
			Username:  usr.Username,
			Email:     usr.Email,
			Blocked:   usr.Blocked,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: usr.UpdatedAt,
		}
		usrsDomain = append(usrsDomain, user)
	}
	return usrsDomain, nil
}

// Update implements ports.UserRepository.
func (*UserPostgresRepository) Update(id string, username *string) (*domain.User, error) {
	panic("unimplemented")
}

func NewUserPostgresRepository() *UserPostgresRepository {
	//criar conexão com banco de dados
	client, err := ent.Open("postgres", os.Getenv("PG_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &UserPostgresRepository{
		db: client,
	}
}

func (r *UserPostgresRepository) Create(email string, username string) (*domain.User, error) {
	usr, err := r.db.User.Create().SetEmail(email).SetUsername(username).Save(context.Background())
	defer r.db.Close()
	if err != nil {
		return nil, err
	}
	return &domain.User{
		Id:        string(usr.ID.String()),
		Username:  usr.Username,
		Email:     usr.Email,
		Blocked:   usr.Blocked,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}, nil
}
