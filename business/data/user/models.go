package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Info struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name,omitempty" json:"name"`
	Slug         string             `bson:"slug,omitempty" json:"slug"`
	Roles        []string           `bson:"roles" json:"roles"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash []byte             `bson:"password_hash" json:"-"`
	Description  string             `bson:"description,omitempty" json:"description"`
	CreatedAt    time.Time          `bson:"created_at,omitempty" json:"created_at"`
}

type NewUser struct {
	Name            string   `json:"name" validate:"required"`
	Description     string   `json:"description"`
	Email           string   `json:"email" validate:"required,email"`
	Roles           []string `json:"roles" validate:"required"`
	Password        string   `json:"password" validate:"required"`
	PasswordConfirm string   `json:"password_confirm" validate:"eqfield=Password"`
}
