package user

import (
	"context"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/thamthee/merchant/business/auth"
)

var (
	// ErrNotFound is used when a specific User is requested but does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("authentication failed")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("attempted action is not allowed")
)

const (
	Collection = "users"
)

type User struct {
	log *logrus.Logger
	mDB *mongo.Database
}

// New constructs a seller for api access.
func New(log *logrus.Logger, db *mongo.Database) User {
	return User{
		log: log,
		mDB: db,
	}
}

func (u User) Create(ctx context.Context, nu NewUser, now time.Time) (Info, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return Info{}, errors.Wrap(err, "generating password hash")
	}

	usr := Info{
		ID:           primitive.NewObjectID(),
		Name:         nu.Name,
		Slug:         strings.ToLower(strings.ReplaceAll(nu.Name, " ", "-")),
		Roles:        nu.Roles,
		Email:        nu.Email,
		PasswordHash: hash,
		Description:  nu.Description,
		CreatedAt:    now.UTC(),
	}

	col := u.mDB.Collection(Collection)

	if _, err := col.InsertOne(ctx, &usr); err != nil {
		return Info{}, errors.Wrap(err, "inserting user")
	}

	return usr, nil
}

func (u User) Authenticate(ctx context.Context, now time.Time, email, password string) (auth.Claims, error) {
	filter := bson.D{
		primitive.E{
			Key:   "email",
			Value: email,
		},
	}

	var usr Info

	col := u.mDB.Collection(Collection)
	if err := col.FindOne(ctx, filter).Decode(&usr); err != nil {
		if err == mongo.ErrNoDocuments {
			return auth.Claims{}, ErrAuthenticationFailure
		}
		return auth.Claims{}, errors.Wrap(err, "selecting user")
	}

	if err := bcrypt.CompareHashAndPassword(usr.PasswordHash, []byte(password)); err != nil {
		return auth.Claims{}, ErrAuthenticationFailure
	}

	claims := auth.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "service project",
			Subject:   usr.ID.Hex(),
			ExpiresAt: jwt.At(now.Add(time.Hour)),
			IssuedAt:  jwt.At(now),
		},
		Roles: usr.Roles,
	}

	return claims, nil
}
