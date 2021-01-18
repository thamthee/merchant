package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	"github.com/thamthee/merchant/business/adapter"
	"github.com/thamthee/merchant/business/auth"
	"github.com/thamthee/merchant/business/graph/models"
)

func (r *mutationResolver) CreateSoftware(ctx context.Context, input models.NewSoftware) (*models.Software, error) {
	claims, ok := ctx.Value(auth.Key).(*auth.Claims)
	if !ok {
		return nil, errors.New("claims missing from context")
	}

	ns := adapter.NewSoftwareToDB(ctx, input)

	sf, err := r.product.CreateSoftware(ctx, claims.Subject, ns, time.Now())
	if err != nil {
		return nil, err
	}

	sl, err := r.seller.QueryByID(ctx, claims.Subject)
	if err != nil {
		return nil, err
	}

	sg := adapter.ProductDBToSoftwareGraph(ctx, sf, sl)

	return &sg, nil
}

func (r *queryResolver) Software(ctx context.Context, id string) (*models.Software, error) {
	claims, ok := ctx.Value(auth.Key).(*auth.Claims)

	if !ok {
		return nil, errors.New("claims missing from context")
	}

	sf, err := r.product.QuerySoftwareByID(ctx, id)
	if err != nil {
		return nil, err
	}

	sl, err := r.seller.QueryByID(ctx, claims.Subject)
	if err != nil {
		return nil, err
	}

	sg := adapter.SoftwareDBToGraph(ctx, sf, sl)

	return &sg, nil
}
