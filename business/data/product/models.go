package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Info represent a individual product information.
type Info struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Seller      string               `bson:"seller,omitempty" json:"seller"`
	Title       string               `bson:"title,omitempty" json:"title"`
	Type        TypeEnum             `bson:"type" json:"type"`
	Price       primitive.Decimal128 `bson:"price,omitempty" json:"price"`
	Currency    string               `bson:"currency,omitempty" json:"currency"`
	Description string               `bson:"description,omitempty" json:"description"`
	SKU         string               `bson:"sku,omitempty" json:"sku"`
	Stock       int                  `bson:"stock,omitempty" json:"stock"`
	Sizes       []string             `bson:"sizes,omitempty" json:"sizes"`
	Colors      []string             `bson:"colors,omitempty" json:"colors"`
	License     string               `bson:"license,omitempty" json:"license"`
	Code        string               `bson:"code,omitempty" json:"code"`
	ExpireAt    time.Time            `bson:"expire_at,omitempty" json:"expire_at"`
	CreatedAt   time.Time            `bson:"created_at,omitempty" json:"created_at"`
}

type Software struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Seller      string               `bson:"seller,omitempty" json:"seller"`
	Title       string               `bson:"title,omitempty" json:"title"`
	Type        TypeEnum             `bson:"type" json:"type"`
	Price       primitive.Decimal128 `bson:"price,omitempty" json:"price"`
	Currency    string               `bson:"currency,omitempty" json:"currency"`
	Description string               `bson:"description,omitempty" json:"description"`
	SKU         string               `bson:"sku,omitempty" json:"sku"`
	Stock       int                  `bson:"stock,omitempty" json:"stock"`
	License     string               `bson:"license,omitempty" json:"license"`
	Code        string               `bson:"code,omitempty" json:"code"`
	CreatedAt   time.Time            `bson:"created_at,omitempty" json:"created_at"`
}

type NewSoftware struct {
	Title       string  `json:"title" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Currency    string  `json:"currency" validate:"required"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Stock       int     `json:"stock"`
	License     string  `json:"license"`
	Code        string  `json:"code"`
}

type Food struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Seller      string               `bson:"seller,omitempty" json:"seller"`
	Title       string               `bson:"title,omitempty" json:"title"`
	Type        TypeEnum             `bson:"type" json:"type"`
	Price       primitive.Decimal128 `bson:"price,omitempty" json:"price"`
	Currency    string               `bson:"currency,omitempty" json:"currency"`
	Description string               `bson:"description,omitempty" json:"description"`
	SKU         string               `bson:"sku,omitempty" json:"sku"`
	Stock       int                  `bson:"stock,omitempty" json:"stock"`
	ExpireAt    time.Time            `bson:"expire_at,omitempty" json:"expire_at"`
	CreatedAt   time.Time            `bson:"created_at,omitempty" json:"created_at"`
}

type NewFood struct {
	Title       string    `json:"title" validate:"required"`
	Price       string    `json:"price" validate:"required"`
	Currency    string    `json:"currency" validate:"required"`
	Description string    `json:"description"`
	SKU         string    `json:"sku"`
	Stock       int       `json:"stock"`
	ExpireAt    time.Time `json:"expire_at"`
}

type Dress struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Seller      string               `bson:"seller,omitempty" json:"seller"`
	Title       string               `bson:"title,omitempty" json:"title"`
	Type        TypeEnum             `bson:"type" json:"type"`
	Price       primitive.Decimal128 `bson:"price,omitempty" json:"price"`
	Currency    string               `bson:"currency,omitempty" json:"currency"`
	Description string               `bson:"description,omitempty" json:"description"`
	SKU         string               `bson:"sku,omitempty" json:"sku"`
	Stock       int                  `bson:"stock,omitempty" json:"stock"`
	Sizes       []string             `bson:"sizes,omitempty" json:"sizes"`
	Colors      []string             `bson:"colors,omitempty" json:"colors"`
	CreatedAt   time.Time            `bson:"created_at,omitempty" json:"created_at"`
}

type NewDress struct {
	Title       string   `json:"title" validate:"required"`
	Price       string   `json:"price" validate:"required"`
	Currency    string   `json:"currency" validate:"required"`
	Description string   `json:"description"`
	SKU         string   `json:"sku"`
	Stock       int      `json:"stock"`
	Sizes       []string `json:"sizes"`
	Colors      []string `json:"colors"`
}
