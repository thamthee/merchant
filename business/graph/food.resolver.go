package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/thamthee/merchant/business/graph/models"
)

func (r *mutationResolver) CreateFood(ctx context.Context, input models.NewFood) (*models.Food, error) {
	return nil, errors.New("waiting for implement")
}

func (r *queryResolver) Food(ctx context.Context, id string) (*models.Food, error) {
	panic(fmt.Errorf("not implemented"))
}
