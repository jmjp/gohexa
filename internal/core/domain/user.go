package domain

import "time"

type User struct {
	Id        string    `json:"id" bson:"_id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    *string   `json:"avatar"`
	Blocked   bool      `json:"blocked"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
