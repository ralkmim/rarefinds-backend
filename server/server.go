package server

import (
	"log"
	"net/http"
	"os"
	"rarefinds-backend/api"
	"rarefinds-backend/common/logger"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func StartServer() {
	productsServer := &http.Server{
		Addr: os.Getenv("HOST") + ":" + os.Getenv("PORT"),
		// Addr: ":9090",
		Handler: api.StartProducts(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	authServer := &http.Server{
		Addr: os.Getenv("HOST") + ":" + os.Getenv("PORT"),
		// Addr: ":9091",
		Handler: api.StartAuth(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return productsServer.ListenAndServe()
	})

	g.Go(func() error {
		return authServer.ListenAndServe()
	})

	logger.Info("products api listening and serving on " + productsServer.Addr)
	logger.Info("auth api listening and serving on " + authServer.Addr)

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}