package api

import (
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/product"
	"rarefinds-backend/internal/product/domain"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartProducts(r *gin.RouterGroup) {
	productHandler := NewProductsHandler(product.NewService(product.NewRepository()))

	r.Use(cors.Default())

	r.POST("/new", productHandler.CreateProduct)
	r.GET("/", productHandler.GetAll)
	r.GET("/ping", productHandler.Ping)
}

type ProductsHandler interface {
	CreateProduct(*gin.Context)
	GetAll(*gin.Context)
	Ping(*gin.Context)
}

type productsHandler struct {
	service product.ProductsService
}

func NewProductsHandler(service product.ProductsService) ProductsHandler {
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