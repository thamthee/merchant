package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/thamthee/merchant/business/adapter"
	"github.com/thamthee/merchant/business/graph/models"
)

func (r *mutationResolver) CreateSeller(ctx context.Context, input models.NewSeller) (*models.Seller, error) {
	ns := adapter.NewSellerGraphToDB(ctx, input)

	seller, err := r.seller.Create(ctx, ns, time.Now())
	if err != nil {
		return nil, err
	}

	sg := adapter.SellerDBToGraph(ctx, seller)

	return &sg, nil
}

func (r *queryResolver) Seller(ctx context.Context, id string) (*models.Seller, error) {
	seller, err := r.seller.QueryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	sg := adapter.SellerDBToGraph(ctx, seller)

	return &sg, nil
}

func (r *queryResolver) Sellers(ctx context.Context, limit int, offer int) ([]*models.Seller, error) {
	panic(fmt.Errorf("not implemented"))
}
