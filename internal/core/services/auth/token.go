package AuthService

import (
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

func GenerateToken(userId string, expires time.Duration) (*string, error) {
	var key = []byte(os.Getenv("PASSETO_KEY"))
	now := time.Now()
	exp := now.Add(expires)
	nbt := now
	jsonToken := paseto.JSONToken{
		Audience:   os.Getenv("APP_HOST"),
		Issuer:     os.Getenv("APP_HOST"),
		Jti:        uuid.New().String(),
		Subject:    userId,
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	jsonToken.Set("user", userId)
	token, err := paseto.NewV2().Encrypt(key, jsonToken, nil)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func Validate(token string) (*string, error) {
	var key = []byte(os.Getenv("PASSETO_KEY"))
	var newJsonToken paseto.JSONToken
	var newFooter string
	err := paseto.NewV2().Decrypt(token, key, &newJsonToken, &newFooter)
	if err != nil {
		return nil, err
	}
	if newJsonToken.Expiration.Before(time.Now()) {
		return nil, errors.New("token expired")
	}
	if newJsonToken.Issuer != os.Getenv("APP_HOST") {
		return nil, errors.New("invalid token claims")
	}
	usr := newJsonToken.Get("user")
	if usr == "" {
		return nil, errors.New("invalid token claims")
	}
	return &usr, nil
}
