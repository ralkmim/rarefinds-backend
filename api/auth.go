package api

import (
	"net/http"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/auth"
	"rarefinds-backend/internal/auth/domain"
	"rarefinds-backend/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StartAuth(r *gin.RouterGroup) {
	authHandler := NewAuthHandler(auth.NewService(auth.NewRepository()))

	r.POST("/signup", authHandler.SignUp)
	r.POST("/login", authHandler.Login)
	r.GET("/user/:user_id", authHandler.GetUserById)

	r.Use(middleware.JWTAuthMiddleware()) 
	{
		r.GET("/user", authHandler.GetUserByToken)
	}
}

type AuthHandler interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	GetUserByToken(c *gin.Context)
	GetUserById(c *gin.Context)
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

func (h *authHandler) GetUserByToken(c *gin.Context) {
	ctx := c.Request.Context()

	sub, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusBadRequest, "No subject claim found")
		return
	}

	subString, ok := sub.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "Subject claim is not a string")
		return
	}

	userId, err := primitive.ObjectIDFromHex(subString)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid subject claim")
		return
	}

	user, errr := h.service.GetUser(userId, ctx)
	if errr != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusFound, &user)
}

func (h *authHandler) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot convert user id")
		return
	}

	user, errr := h.service.GetUser(userId, ctx)
	if errr != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusFound, &user)
}

