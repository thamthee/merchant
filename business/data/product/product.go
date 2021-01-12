package product

import (
	"context"
	"time"

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

// Create inserts a new product into the database.
func (p Product) Create(ctx context.Context, seller string, np NewProduct, date time.Time) (Info, error) {
	price, err := primitive.ParseDecimal128(np.Price)
	if err != nil {
		return Info{}, errors.Wrap(err, "parse price")
	}

	info := Info{
		ID:          primitive.NewObjectID(),
		Title:       np.Title,
		Price:       price,
		Currency:    np.Currency,
		Description: np.Description,
		SKU:         np.SKU,
		Stock:       np.Stock,
		Sizes:       np.Sizes,
		Colors:      np.Colors,
		Seller:      seller,
		CreatedAt:   date.UTC(),
	}

	col := p.db.Collection(Collection)

	if _, err := col.InsertOne(ctx, &info); err != nil {
		return Info{}, errors.Wrap(err, "inserting product")
	}

	return info, nil
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
