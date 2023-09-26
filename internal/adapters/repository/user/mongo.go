package UserRepository

import (
	"context"
	"deluze/internal/core/domain"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UserMongoRepository struct {
	db *mongo.Collection
}

// Find implements ports.UserRepository.
func (r *UserMongoRepository) Find(id string) (*domain.User, error) {
	panic("unimplemented")
}

// FindByEmail implements ports.UserRepository.
func (r *UserMongoRepository) FindByEmail(email string) (*domain.User, error) {
	var result *domain.User
	err := r.db.FindOne(context.Background(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMany implements ports.UserRepository.
func (r *UserMongoRepository) FindMany(page int, limit int) ([]*domain.User, error) {
	var result []*domain.User
	usr, err := r.db.Find(context.Background(), bson.M{}, options.Find().SetSkip(int64(page)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	usr.All(context.Background(), &result)
	return result, nil
}

// Update implements ports.UserRepository.
func (*UserMongoRepository) Update(id string, username *string) (*domain.User, error) {
	panic("unimplemented")
}

func NewUserMongoRepository() *UserMongoRepository {
	//criar conexão com banco de dados
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		log.Fatalf("failed opening connection to mongodb: %v", err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return &UserMongoRepository{
		db: client.Database("deluze").Collection("users"),
	}
}

func (r *UserMongoRepository) Create(email string, username string) (*domain.User, error) {
	usr, err := r.db.InsertOne(context.Background(), bson.M{"username": username, "email": email, "created_at": time.Now(), "updated_at": time.Now()})
	if err != nil {
		return nil, err
	}
	var user *domain.User
	err = r.db.FindOne(context.Background(), bson.M{"_id": usr.InsertedID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
