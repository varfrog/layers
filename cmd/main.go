package main

import (
	"layers/internal/appservice"
	"layers/internal/rest"
	"layers/internal/storagemysql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("create zap logger", err)
		return
	}

	dbPort, found := os.LookupEnv("DBPORT")
	if !found {
		log.Fatal("env var DBPORT not set", err)
		return
	}

	dbHost, found := os.LookupEnv("DBHOST")
	if !found {
		log.Fatal("env var DBHOST not set", err)
		return
	}

	db, err := storagemysql.Connect("app", "app", "pass", dbHost, dbPort)
	if err != nil {
		logger.Fatal("connect to mysql", zap.Error(err))
		return
	}

	router := gin.Default()
	err = router.SetTrustedProxies(nil)
	if err != nil {
		logger.Fatal("SetTrustedProxies", zap.Error(err))
		return
	}

	rest.CreateRoutes(
		router,
		rest.NewRequestHandler(
			appservice.NewApp(storagemysql.NewStorageMysql(db), appservice.NewItemTransformer()),
			rest.NewItemTransformer(),
			logger))

	logger.Fatal("run router", zap.Error(router.Run()))
}
