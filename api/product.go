package api

import (
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/product"
	"rarefinds-backend/internal/product/domain"
	"rarefinds-backend/internal/product/repository"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartProducts() http.Handler {
	var router = gin.Default()
	productHandler := NewHandler(product.NewService(repository.NewRepository()))

	router.Use(cors.Default())

	router.POST("/new-product", productHandler.CreateProduct)
	router.GET("/products", productHandler.GetAll)
	router.GET("/ping", productHandler.Ping)

	return router
}

type ProductsHandler interface {
	CreateProduct(*gin.Context)
	GetAll(*gin.Context)
	Ping(*gin.Context)
}

type productsHandler struct {
	service product.ProductsService
}

func NewHandler(service product.ProductsService) ProductsHandler {
	return &productsHandler{
		service: service,
	}
}

func(h *productsHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		fmt.Println(err)
		c.JSON(errors.NewBadRequestError(err.Error()).Status, errors.NewBadRequestError("invalid json body"))
		return
	}

	result, err := h.service.CreateProduct(product)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *productsHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	clients, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(errors.NewNotFoundError(err.Error).Status, errors.NewNotFoundError(err.Message))
		return
	}

	c.JSON(http.StatusAccepted, clients)
}

func (h *productsHandler) Ping(c *gin.Context) {
	c.String(200, "Pong!")
}