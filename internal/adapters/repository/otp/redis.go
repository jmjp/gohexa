package OtpRepository

import (
	"context"
	"deluze/internal/core/domain"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type OtpRedisRepository struct {
	db *redis.Client
}

func NewOtpRedisRepository() *OtpRedisRepository {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_DB"))
	client := redis.NewClient(opt)
	return &OtpRedisRepository{
		db: client,
	}
}

func (r *OtpRedisRepository) Create(code string, user string, expiresAt time.Time) (*domain.Otp, error) {
	str, err := r.db.Set(context.Background(), user, code, time.Until(expiresAt)).Result()
	if err != nil {
		return nil, err
	}
	return &domain.Otp{
		Id:        str,
		Code:      code,
		User:      user,
		ExpiredAt: expiresAt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Delete implements ports.OtpRepository.
func (r *OtpRedisRepository) Delete(id string) error {
	//ignore because redis already delete when time expires
	return nil
}

// Find implements ports.OtpRepository.
func (r *OtpRedisRepository) Find(code string, user string) (*domain.Otp, error) {
	val, err := r.db.Get(context.Background(), user).Result()
	if err != nil {
		return nil, err
	}
	return &domain.Otp{
		Id:        val,
		Code:      code,
		User:      user,
		ExpiredAt: time.Now(), //set now because redis already delete when time expires
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, err
}
