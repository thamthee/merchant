package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/thamthee/merchant/app/merchant-graph/handlers"
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
	if err != nil {
		return errors.Wrap(err, "unable to connect mongo db")
	}

	defer func() {
		log.Printf("main: Database shutting down: %v", config.Mongo.Hosts)
		db.Disconnect(context.Background())
	}()

	// =========================================
	// Start API Service

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	serverError := make(chan error, 1)

	go func() {
		log.Printf("main: API listening on %s", config.Web.APIHost)
		serverError <- handlers.API(log, db, config.Web.APIHost)
	}()

	// =========================================
	// Shutdown

	select {
	case err := <-serverError:
		return errors.Wrap(err, "server error")
	case sig := <-shutdown:
		log.Printf("main: %v : Shuting down", sig)
		// TODO: Teardown graphql here.
	}

	return nil
}
