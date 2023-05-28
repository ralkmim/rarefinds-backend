package server

import (
	"rarefinds-backend/api"
	"rarefinds-backend/common/logger"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func StartServer() {
	productsServer := &http.Server{
		Addr: ":9090",
		Handler: api.StartProducts(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return productsServer.ListenAndServe()
	})

	logger.Info("products api listening and serving on " + productsServer.Addr)

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}