package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/thamthee/merchant/business/graph/models"
)

func (r *mutationResolver) CreateSeller(ctx context.Context, input models.NewSeller) (*models.Seller, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Seller(ctx context.Context) (*models.Seller, error) {
	panic(fmt.Errorf("not implemented"))
}
