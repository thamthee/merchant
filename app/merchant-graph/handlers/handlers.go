package handlers

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/thamthee/merchant/business/data/product"
	"github.com/thamthee/merchant/business/data/seller"
	"github.com/thamthee/merchant/business/graph"
	"github.com/thamthee/merchant/business/graph/generated"
	"github.com/thamthee/merchant/business/mid"
	"github.com/thamthee/merchant/foundation/database"
)

func API(log *logrus.Logger, mClient *mongo.Client, addr string) error {
	router := chi.NewRouter()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.New(
			product.New(log, mClient.Database(database.Merchant)),
			seller.New(log, mClient.Database(database.Merchant)),
		),
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}))

	router.Handle("/", playground.Handler("Playground", "/query"))

	router.Group(func(r chi.Router) {
		r.Use(mid.BypassToken)
		r.Handle("/query", srv)
	})

	return http.ListenAndServe(addr, router)
}
