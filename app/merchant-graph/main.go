package main

import (
	"context"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/thamthee/merchant/business/data/product"
	"github.com/thamthee/merchant/business/data/seller"
	"github.com/thamthee/merchant/business/graph"
	"github.com/thamthee/merchant/business/graph/generated"
	"github.com/thamthee/merchant/configs"
	"github.com/thamthee/merchant/foundation/database"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	if err := run(log); err != nil {
		log.Fatalln(err)
	}
}

func run(log *logrus.Logger) error {

	// =========================================
	// Configuration

	file, err := os.Open("./configs/config.yml")
	if err != nil {
		return errors.Wrap(err, "read file config")
	}

	config, err := configs.ParseFrom(file, "yaml")
	if err != nil {
		return errors.Wrap(err, "parse from file")
	}

	// =========================================
	// Start Database
	db, err := database.Open(context.Background(), database.MGConfig{
		AuthMechanism: config.Mongo.AuthMechanism,
		AuthSource:    config.Mongo.AuthSource,
		Username:      config.Mongo.Username,
		Password:      config.Mongo.Password,
		Hosts:         config.Mongo.Hosts,
		ReplicaName:   config.Mongo.ReplicaSetName,
		SSL: func() *database.MGSSL {
			if config.Mongo.SSL != nil {
				return &database.MGSSL{
					Type: config.Mongo.SSL.Type,
					Cert: config.Mongo.SSL.Cert,
				}
			}
			return nil
		}(),
		IsDirect: config.Mongo.IsDirect,
		ReadPref: readpref.Nearest(),
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.New(
			product.New(log, db.Database(database.Merchant)),
			seller.New(log, db.Database(database.Merchant)),
		),
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}))

	http.Handle("/", playground.Handler("Playground", "/query"))
	http.Handle("/query", srv)

	return http.ListenAndServe(":3000", nil)
}
