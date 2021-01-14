package seller

import (
	"context"
	"strings"
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
	Collection = "sellers"
)

// Seller manage set of API's seller access.
type Seller struct {
	log *logrus.Logger
	db  *mongo.Database
}

// New constructs a seller for api access.
func New(log *logrus.Logger, db *mongo.Database) Seller {
	return Seller{
		log: log,
		db:  db,
	}
}

// Create inserts a specific seller into the database.
func (s Seller) Create(ctx context.Context, ns NewSeller, date time.Time) (Info, error) {
	slug := ns.Slug
	if ns.Slug == "" {
		slug = strings.ToLower(strings.ReplaceAll(ns.Title, " ", "-"))
	}

	info := Info{
		ID:          primitive.NewObjectID(),
		Title:       ns.Title,
		Slug:        slug,
		Description: ns.Description,
		CreatedAt:   date.UTC(),
	}

	col := s.db.Collection(Collection)

	if _, err := col.InsertOne(ctx, &info); err != nil {
		return Info{}, errors.Wrap(err, "inserting seller")
	}

	return info, nil
}

// QueryByID gets a specific seller information from the database.
func (s Seller) QueryByID(ctx context.Context, id string) (Info, error) {
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

	col := s.db.Collection(Collection)

	var info Info
	if err := col.FindOne(ctx, filter).Decode(&info); err != nil {
		if err == mongo.ErrNoDocuments {
			return Info{}, ErrNotFound
		}
		return Info{}, errors.Wrap(err, "selecting seller")
	}

	return info, nil
}