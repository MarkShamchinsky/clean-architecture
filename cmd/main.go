package main

import (
	"ca-library-go/internal/composites"
	"ca-library-go/internal/config"
	"ca-library-go/pkg/logging"
	"context"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("Config initializing")
	cfg := config.GetConfig()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("mongodb composite initializing")
	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, "", "", "", "")
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}
	logger.Info("author composite initializing")
	authorComposite, err := composites.NewAuthorComposite(mongoDBC)
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	authorComposite.Handler.Register(router)

	logger.Info("book composite initializing")
	bookComposite, err := composites.NewBookComposite(mongoDBC)
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	bookComposite.Handler.Register(router)

}
