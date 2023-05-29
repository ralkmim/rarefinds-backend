package api

import (
	"net/http"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/auth"
	"rarefinds-backend/internal/auth/domain"

	"github.com/gin-gonic/gin"
)

func StartAuth(r *gin.RouterGroup) {
	authHandler := NewAuthHandler(auth.NewService(auth.NewRepository()))

	r.POST("/signup", authHandler.SignUp)
}

type UsersHandler interface {
	SignUp(c *gin.Context)
}

type usersHandler struct {
	service auth.UsersService
}

func NewAuthHandler(service auth.UsersService) UsersHandler {
	return &usersHandler{
		service: service,
	}
}

func (h *usersHandler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	var payload *domain.SignUpInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(errors.NewBadRequestError(err.Error()).Status, errors.NewBadRequestError("invalid json body"))
		return
	}

	err := h.service.CreateUser(payload, ctx)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}