package handlers

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/thamthee/merchant/business/data/product"
	"github.com/thamthee/merchant/business/data/seller"
	"github.com/thamthee/merchant/business/graph"
	"github.com/thamthee/merchant/business/graph/generated"
	"github.com/thamthee/merchant/foundation/database"
)

func API(log *logrus.Logger, mClient *mongo.Client, addr string) error {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.New(
			product.New(log, mClient.Database(database.Merchant)),
			seller.New(log, mClient.Database(database.Merchant)),
		),
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}))

	http.Handle("/", playground.Handler("Playground", "/query"))
	http.Handle("/query", srv)

	return http.ListenAndServe(addr, nil)
}
