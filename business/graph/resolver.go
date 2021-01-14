package graph

import (
	"github.com/thamthee/merchant/business/data/product"
	"github.com/thamthee/merchant/business/data/seller"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	product product.Product
	seller  seller.Seller
}

func New(p product.Product, s seller.Seller) *Resolver {
	return &Resolver{
		product: p,
		seller:  s,
	}
}
