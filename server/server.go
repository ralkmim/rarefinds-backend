package server

import (
	"log"
	"net/http"
	"os"

	// "os"
	"rarefinds-backend/api"
	"rarefinds-backend/common/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func StartServer() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// AllowAllOrigins: 	true,
		AllowOrigins: 		[]string{"http://localhost:8080", "http://localhost:5173", "http://127.0.0.1:5500", "https://rarefinds.herokuapp.com"},
		AllowMethods: 		[]string{"PUT","PATCH","GET","DELETE","POST","OPTIONS"},
		AllowHeaders: 		[]string{"Origin","Content-type","Authorization","Content-Length","Content-Language",
										"Content-Disposition","User-Agent","Referrer","Host","Access-Control-Allow-Origin","sentry-trace"},
		ExposeHeaders: 		[]string{"Authorization","Content-Length"},
		AllowCredentials: 	true,
		MaxAge: 			12*time.Hour,	
	}))

	api.StartAuth(router.Group("/auth"))
	api.StartProducts(router.Group("/products"))

	server := &http.Server{
		Addr: os.Getenv("HOST") + ":" + os.Getenv("PORT"),
		// Addr: ":9090",
		Handler: router,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("Server listening and serving on " + server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
	// productsServer := &http.Server{
	// 	// Addr: os.Getenv("HOST") + ":" + os.Getenv("PORT"),
	// 	Addr: ":9090",
	// 	Handler: api.StartProducts(),
	// 	ReadTimeout: 5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	// authServer := &http.Server{
	// 	// Addr: os.Getenv("HOST") + ":" + os.Getenv("PORT"),
	// 	Addr: ":9091",
	// 	Handler: api.StartAuth(),
	// 	ReadTimeout: 5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	// g.Go(func() error {
	// 	return productsServer.ListenAndServe()
	// })

	// g.Go(func() error {
	// 	return authServer.ListenAndServe()
	// })

	// logger.Info("products api listening and serving on " + productsServer.Addr)
	// logger.Info("auth api listening and serving on " + authServer.Addr)

	// if err := g.Wait(); err != nil {
	// 	log.Fatal(err)
	// }
// }