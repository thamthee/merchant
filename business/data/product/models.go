package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Info represent a individual product information.
type Info struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title       string               `bson:"title,omitempty" json:"title"`
	Price       primitive.Decimal128 `bson:"price,omitempty" json:"price"`
	Currency    string               `bson:"currency,omitempty" json:"currency"`
	Description string               `bson:"description,omitempty" json:"description"`
	SKU         string               `bson:"sku,omitempty" json:"sku"`
	Stock       int                  `bson:"stock,omitempty" json:"stock"`
	Sizes       []string             `bson:"sizes,omitempty" json:"sizes"`
	Colors      []string             `bson:"colors,omitempty" json:"colors"`
	Vendor      string               `bson:"vendor,omitempty" json:"vendor"`
	CreatedAt   time.Time            `bson:"created_at,omitempty" json:"created_at"`
}

type NewProduct struct {
	Title       string   `json:"title" validate:"required"`
	Price       string   `json:"price" validate:"required"`
	Currency    string   `json:"currency" validate:"required"`
	Description string   `json:"description"`
	SKU         string   `json:"sku"`
	Stock       int      `json:"stock"`
	Sizes       []string `json:"sizes"`
	Colors      []string `json:"colors"`
}
