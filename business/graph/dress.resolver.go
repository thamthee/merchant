package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/thamthee/merchant/business/graph/generated"
	"github.com/thamthee/merchant/business/graph/models"
)

func (r *mutationResolver) CreateDress(ctx context.Context, input models.NewDress) (*models.Dress, error) {
	return nil, errors.New("waiting for implement")
}

func (r *queryResolver) Dress(ctx context.Context, id string) (*models.Dress, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
