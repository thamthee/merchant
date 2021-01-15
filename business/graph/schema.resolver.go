package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/thamthee/merchant/business/graph/generated"
	"github.com/thamthee/merchant/business/graph/models"
)

func (r *queryResolver) Product(ctx context.Context, id string) (models.SearchResult, error) {
	return nil, errors.New("not found")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
