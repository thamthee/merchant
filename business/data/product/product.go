package product

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound = errors.New("not found")

	ErrInvalidID = errors.New("invalid id")
)

const (
	Collection = "products"
)

// Product manage set of API's for product access.
type Product struct {
	log *logrus.Logger
	db  *mongo.Database
}

// New constructs a product for access api.
func New(log *logrus.Logger, db *mongo.Database) Product {
	return Product{
		log: log,
		db:  db,
	}
}

// QueryByID gets a specific product from the database by id.
func (p Product) QueryByID(ctx context.Context, id string) (Info, error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Info{}, ErrInvalidID
	}

	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: hex,
		},
	}

	col := p.db.Collection(Collection)

	var info Info
	if err := col.FindOne(ctx, filter).Decode(&info); err != nil {
		if err == mongo.ErrNoDocuments {
			return Info{}, ErrNotFound
		}
		return Info{}, errors.Wrapf(err, "finding product : %q", id)
	}

	return info, nil
}

func (p Product) QuerySoftwareByID(ctx context.Context, id string) (Software, error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Software{}, ErrInvalidID
	}

	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: hex,
		},
	}

	col := p.db.Collection(Collection)

	var sw Software
	if err := col.FindOne(ctx, filter).Decode(&sw); err != nil {
		if err == mongo.ErrNoDocuments {
			return Software{}, ErrNotFound
		}
		return Software{}, errors.Wrapf(err, "finding software : %q", id)
	}

	return sw, nil
}

func (p Product) QueryFoodByID(ctx context.Context, id string) (Food, error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Food{}, ErrInvalidID
	}

	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: hex,
		},
	}

	col := p.db.Collection(Collection)

	var food Food
	if err := col.FindOne(ctx, filter).Decode(&food); err != nil {
		if err == mongo.ErrNoDocuments {
			return Food{}, ErrNotFound
		}
		return Food{}, errors.Wrapf(err, "finding food : %q", id)
	}

	return food, nil
}

func (p Product) QueryDressByID(ctx context.Context, id string) (Dress, error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Dress{}, ErrInvalidID
	}

	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: hex,
		},
	}

	col := p.db.Collection(Collection)

	var dress Dress
	if err := col.FindOne(ctx, filter).Decode(&dress); err != nil {
		if err == mongo.ErrNoDocuments {
			return Dress{}, ErrNotFound
		}
		return Dress{}, errors.Wrapf(err, "finding dress : %q", id)
	}

	return dress, nil
}
