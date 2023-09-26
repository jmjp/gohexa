package OtpRepository

import (
	"context"
	"deluze/internal/core/domain"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type OtpMongoRepository struct {
	db *mongo.Collection
}

func NewOtpMongoRepository() *OtpMongoRepository {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		log.Fatalf("failed opening connection to mongodb: %v", err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return &OtpMongoRepository{
		db: client.Database("deluze").Collection("otps"),
	}
}

// Create implements ports.OtpRepository.
func (r *OtpMongoRepository) Create(code string, user string, expiresAt time.Time) (*domain.Otp, error) {
	id, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return nil, err
	}
	insertedDocument, err := r.db.InsertOne(context.Background(), bson.M{"code": code, "user": id, "expires_at": expiresAt, "created_at": time.Now(), "updated_at": time.Now()})
	if err != nil {
		return nil, err
	}
	var otp *domain.Otp
	err = r.db.FindOne(context.Background(), bson.M{"_id": insertedDocument.InsertedID}).Decode(&otp)
	if err != nil {
		return nil, err
	}
	return otp, err
}

// Delete implements ports.OtpRepository.
func (r *OtpMongoRepository) Delete(id string) error {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := r.db.DeleteOne(context.Background(), bson.M{"_id": idHex})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("not document found")
	}
	return nil
}

// Find implements ports.OtpRepository.
func (r *OtpMongoRepository) Find(code string, user string) (*domain.Otp, error) {
	var otp *domain.Otp
	id, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return nil, err
	}
	err = r.db.FindOne(context.Background(), bson.M{"user": id, "code": code}).Decode(&otp)
	if err != nil {
		return nil, err
	}
	return otp, err
}
