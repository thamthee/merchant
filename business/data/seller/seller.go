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
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s Seller) QueryByIDs(ctx context.Context, ids []string) ([]Info, error) {
	var hexs []primitive.ObjectID
	for _, id := range ids {
		hex, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, ErrInvalidID
		}

		hexs = append(hexs, hex)
	}

	filter := bson.D{
		primitive.E{
			Key: "_id",
			Value: bson.D{
				primitive.E{
					Key:   "$in",
					Value: hexs,
				},
			},
		},
	}

	col := s.db.Collection(Collection)

	res, err := col.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "selecting sellers")
	}
	var sellers []Info
	if err := res.All(ctx, &sellers); err != nil {
		return nil, errors.Wrap(err, "unable to decode")
	}

	return sellers, nil
}

func (s Seller) QueryAllByPaginate(ctx context.Context, limit, offer int) ([]Info, error) {
	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offer))

	col := s.db.Collection(Collection)

	var sellers []Info
	cursor, err := col.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "selecting")
	}

	if err := cursor.All(ctx, &sellers); err != nil {
		return nil, errors.Wrap(err, "unable to decode payload")
	}

	return sellers, nil
}
