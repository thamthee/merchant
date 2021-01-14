package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/pkg/errors"

	"github.com/thamthee/merchant/business/graph/generated"
	"github.com/thamthee/merchant/business/graph/models"
)

func (r *mutationResolver) CreateFood(ctx context.Context, input models.NewFood) (*models.Food, error) {
	return nil, errors.New("waiting for implement")
}

func (r *mutationResolver) CreateDress(ctx context.Context, input models.NewDress) (*models.Dress, error) {
	return nil, errors.New("waiting for implement")
}

func (r *queryResolver) Product(ctx context.Context, id string) (models.SearchResult, error) {
	return nil, errors.New("not found")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
