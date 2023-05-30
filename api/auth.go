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
	r.POST("/login", authHandler.Login)
}

type AuthHandler interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
}

type authHandler struct {
	service auth.UsersService
}

func NewAuthHandler(service auth.UsersService) AuthHandler {
	return &authHandler{
		service: service,
	}
}

func (h *authHandler) SignUp(c *gin.Context) {
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

func (h *authHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var payload *domain.SignInInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(errors.NewBadRequestError(err.Error()).Status, errors.NewBadRequestError("invalid json body"))
		return
	}

	token, err := h.service.LoginUser(payload, ctx)
	if err != nil {
		c.JSON(err.Status, err)
		return 
	}

	c.SetSameSite(http.SameSiteNoneMode)
	// c.SetCookie("token", token, 60*60, "/", "localhost", false, true)
	c.SetCookie("token", token, 60*60, "/", "127.0.0.1:5050", true, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}