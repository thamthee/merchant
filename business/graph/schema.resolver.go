package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/thamthee/merchant/business/graph/generated"
	"github.com/thamthee/merchant/business/graph/models"
)

func (r *mutationResolver) CreateFood(ctx context.Context, input models.NewFood) (*models.Food, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDress(ctx context.Context, input models.NewDress) (*models.Dress, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id string) (models.SearchResult, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
