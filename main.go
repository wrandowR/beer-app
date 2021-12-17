package main

import (
	"ZachIgarz/test-beer/config"
	"ZachIgarz/test-beer/infrastructure/datastore"
	"ZachIgarz/test-beer/infrastructure/router"
	"ZachIgarz/test-beer/interface/controller"
	"ZachIgarz/test-beer/usecase/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ansel1/merry"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func main() {

	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = datastore.NewDBConn()
	if err != nil {
		log.Fatal(err)
	}

	if config.EnableMigrations() {
		datastore.DoMigration()
	}

	//initialize currency layer service
	service.NewCurrencyLayer()

	appController := controller.AppController{
		Beers: controller.BeerController,
	}

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())

	e.GET("/healthz", healthHandler)

	e = router.NewRouter(e, appController)

	go func() {
		if err := e.Start(config.HTTPListener()); err != nil {
			logrus.WithError(err).Error("shutting down server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Panic(merry.Wrap(err))
	}
}

func healthHandler(c echo.Context) error {
	row := datastore.SQLDB.QueryRow("SELECT 1")
	var val int
	if err := row.Scan(&val); err != nil {
		return merry.Wrap(err)
	}
	if val != 1 {
		return merry.New("error query pg")
	}
	return c.NoContent(http.StatusOK)
}
