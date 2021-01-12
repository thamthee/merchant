package seller

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Info represent a individual seller information.
type Info struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title,omitempty" json:"title"`
	Slug        string             `bson:"slug,omitempty" json:"slug"`
	Description string             `bson:"description,omitempty" json:"description"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at"`
}

// NewSeller what we required for insert data into the database.
type NewSeller struct {
	Title       string `json:"title" validate:"required"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}
