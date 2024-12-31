package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/azamon/internal/models"
)

type Service interface {
	SignUp(req models.SignUpRequest) (models.Profile, error)
	SignIn(req models.SignInRequest) (string, error)
	SignOut(token string) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-out", h.signOut)
	}

	return router
}
