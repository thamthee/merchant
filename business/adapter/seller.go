package adapter

import (
	"context"

	"github.com/thamthee/merchant/business/data/seller"
	"github.com/thamthee/merchant/business/graph/models"
)

func NewSellerGraphToDB(ctx context.Context, ns models.NewSeller) seller.NewSeller {
	return seller.NewSeller{
		Title:       ns.Title,
		Slug:        pointerStringToString(ns.Slug),
		Description: ns.Description,
	}
}

func SellerDBToGraph(ctx context.Context, info seller.Info) models.Seller {
	return models.Seller{
		ID:          info.ID.Hex(),
		Title:       info.Title,
		Slug:        info.Slug,
		Description: info.Description,
	}
}
