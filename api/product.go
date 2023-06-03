package api

import (
	"fmt"
	"net/http"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/product"
	"rarefinds-backend/internal/product/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StartProducts(r *gin.RouterGroup) {
	productHandler := NewProductsHandler(product.NewService(product.NewRepository()))


	r.POST("/new", productHandler.CreateProduct)
	r.GET("/", productHandler.GetAll)
	r.GET("/:id", productHandler.GetProduct)
	r.GET("/search", productHandler.SearchProducts)
	r.GET("/ping", productHandler.Ping)
}

type ProductsHandler interface {
	CreateProduct(*gin.Context)
	GetAll(*gin.Context)
	GetProduct(*gin.Context)
	SearchProducts(*gin.Context)
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

func (h *productsHandler) GetProduct(c *gin.Context) {
	ctx := c.Request.Context()

	productId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot convert product id")
		return
	}

	product, errr := h.service.GetProduct(productId, ctx)
	if errr != nil {
		c.JSON(http.StatusNotFound, "Product not found")
		return
	}

	fmt.Println(&product)

	c.JSON(http.StatusFound, &product)
}

func (h *productsHandler) SearchProducts(c *gin.Context) {
	ctx := c.Request.Context()
	search := c.Query("q")

	products, err := h.service.SearchProducts(search, ctx)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusFound, products)
}



func (h *productsHandler) Ping(c *gin.Context) {
	c.String(200, "Pong!")
}