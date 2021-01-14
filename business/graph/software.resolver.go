package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/thamthee/merchant/business/graph/models"
)

func (r *mutationResolver) CreateSoftware(ctx context.Context, input models.NewSoftware) (*models.Software, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Software(ctx context.Context, id string) (*models.Software, error) {
	panic(fmt.Errorf("not implemented"))
}
