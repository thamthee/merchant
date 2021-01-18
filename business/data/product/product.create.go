package product

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateSoftware inserts a new product software into the database.
func (p Product) CreateSoftware(ctx context.Context, seller string, ns NewSoftware, date time.Time) (Info, error) {
	price, err := primitive.ParseDecimal128(fmt.Sprintf("%v", ns.Price))

	if err != nil {
		return Info{}, errors.Wrap(err, "parse price")
	}

	info := Info{
		ID:          primitive.NewObjectID(),
		Title:       ns.Title,
		Type:        SoftwareType,
		Price:       price,
		Currency:    ns.Currency,
		Description: ns.Description,
		SKU:         ns.SKU,
		Stock:       ns.Stock,
		Seller:      seller,
		License:     ns.License,
		Code:        ns.Code,
		CreatedAt:   date.UTC(),
	}

	col := p.db.Collection(Collection)

	if _, err := col.InsertOne(ctx, &info); err != nil {
		return Info{}, errors.Wrap(err, "inserting software")
	}

	return info, nil
}

// CreateFood inserts a new product food into the database.
func (p Product) CreateFood(ctx context.Context, seller string, nf NewFood, date time.Time) (Info, error) {
	price, err := primitive.ParseDecimal128(nf.Price)
	if err != nil {
		return Info{}, errors.Wrap(err, "parse price")
	}

	info := Info{
		ID:          primitive.NewObjectID(),
		Title:       nf.Title,
		Type:        FoodType,
		Price:       price,
		Currency:    nf.Currency,
		Description: nf.Description,
		SKU:         nf.SKU,
		Stock:       nf.Stock,
		Seller:      seller,
		ExpireAt:    nf.ExpireAt.UTC(),
		CreatedAt:   date.UTC(),
	}

	col := p.db.Collection(Collection)

	if _, err := col.InsertOne(ctx, &info); err != nil {
		return Info{}, errors.Wrap(err, "inserting food")
	}

	return info, nil
}

// CreateDress inserts a new product dress into the database.
func (p Product) CreateDress(ctx context.Context, seller string, nd NewDress, date time.Time) (Info, error) {
	price, err := primitive.ParseDecimal128(nd.Price)
	if err != nil {
		return Info{}, errors.Wrap(err, "parse price")
	}

	info := Info{
		ID:          primitive.NewObjectID(),
		Title:       nd.Title,
		Type:        DressType,
		Price:       price,
		Currency:    nd.Currency,
		Description: nd.Description,
		SKU:         nd.SKU,
		Stock:       nd.Stock,
		Seller:      seller,
		Sizes:       nd.Sizes,
		Colors:      nd.Colors,
		CreatedAt:   date.UTC(),
	}

	col := p.db.Collection(Collection)

	if _, err := col.InsertOne(ctx, &info); err != nil {
		return Info{}, errors.Wrap(err, "inserting dress")
	}

	return info, nil
}
