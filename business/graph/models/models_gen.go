// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type Product interface {
	IsProduct()
}

type SearchResult interface {
	IsSearchResult()
}

type Dress struct {
	ID          string     `json:"id"`
	Title       *string    `json:"title"`
	Price       *float64   `json:"price"`
	Currency    *string    `json:"currency"`
	Description *string    `json:"description"`
	Sku         *string    `json:"sku"`
	Stock       *int       `json:"stock"`
	Owner       *Seller    `json:"owner"`
	CreateAt    *time.Time `json:"createAt"`
	Sizes       []*string  `json:"sizes"`
	Colors      []*string  `json:"colors"`
}

func (Dress) IsSearchResult() {}
func (Dress) IsProduct()      {}

type Food struct {
	ID          string     `json:"id"`
	Title       *string    `json:"title"`
	Price       *float64   `json:"price"`
	Currency    *string    `json:"currency"`
	Description *string    `json:"description"`
	Sku         *string    `json:"sku"`
	Stock       *int       `json:"stock"`
	Owner       *Seller    `json:"owner"`
	CreateAt    *time.Time `json:"createAt"`
	ExpireAt    *time.Time `json:"expireAt"`
}

func (Food) IsSearchResult() {}
func (Food) IsProduct()      {}

type NewDress struct {
	Title       *string   `json:"title"`
	Price       *float64  `json:"price"`
	Currency    *string   `json:"currency"`
	Description *string   `json:"description"`
	Sku         *string   `json:"sku"`
	Stock       *int      `json:"stock"`
	Sizes       []*string `json:"sizes"`
	Colors      []*string `json:"colors"`
}

type NewFood struct {
	Title       *string    `json:"title"`
	Price       *float64   `json:"price"`
	Currency    *string    `json:"currency"`
	Description *string    `json:"description"`
	Sku         *string    `json:"sku"`
	Stock       *int       `json:"stock"`
	ExpireAt    *time.Time `json:"expireAt"`
}

type NewSeller struct {
	Title       *string `json:"title"`
	Slug        *string `json:"slug"`
	Description *string `json:"description"`
}

type NewSoftware struct {
	Title       *string  `json:"title"`
	Price       *float64 `json:"price"`
	Currency    *string  `json:"currency"`
	Description *string  `json:"description"`
	Sku         *string  `json:"sku"`
	Stock       *int     `json:"stock"`
	License     *string  `json:"license"`
	Code        *string  `json:"code"`
}

type Seller struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	Slug        *string `json:"slug"`
	Description *string `json:"description"`
}

type Software struct {
	ID          string     `json:"id"`
	Title       *string    `json:"title"`
	Price       *float64   `json:"price"`
	Currency    *string    `json:"currency"`
	Description *string    `json:"description"`
	Sku         *string    `json:"sku"`
	Stock       *int       `json:"stock"`
	Owner       *Seller    `json:"owner"`
	CreateAt    *time.Time `json:"createAt"`
	License     *string    `json:"license"`
	Code        *string    `json:"code"`
}

func (Software) IsSearchResult() {}
func (Software) IsProduct()      {}
