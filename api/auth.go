package api

import (
	"net/http"
	"rarefinds-backend/common/errors"
	"rarefinds-backend/internal/auth"
	"rarefinds-backend/internal/auth/domain"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartAuth(r *gin.RouterGroup) {
	authHandler := NewAuthHandler(auth.NewService(auth.NewRepository()))

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: 	true,
		// AllowOrigins: 		[]string{"http://localhost:8080", "http://localhost:5173", "http://127.0.0.1:5500"},
		AllowMethods: 		[]string{"PUT","PATCH","GET","DELETE","POST","OPTIONS"},
		AllowHeaders: 		[]string{"Origin","Content-type","Authorization","Content-Length","Content-Language",
										"Content-Disposition","User-Agent","Referrer","Host","Access-Control-Allow-Origin","sentry-trace"},
		ExposeHeaders: 		[]string{"Authorization","Content-Length"},
		AllowCredentials: 	true,
		MaxAge: 			12*time.Hour,	
	}))

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