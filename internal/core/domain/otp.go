package domain

import "time"

type Otp struct {
	Id        string    `json:"id" bson:"_id,omitempty"`
	Code      string    `json:"code"`
	User      string    `json:"user"`
	ExpiredAt time.Time `json:"expired_at" bson:"expired_at"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
