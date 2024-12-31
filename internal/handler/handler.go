package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ursulgwopp/azamon/internal/models"
)

type Service interface {
	SignUp(req models.SignUpRequest) (models.Profile, error)
	SignIn(req models.SignInRequest) (string, error)
	SignOut(token string) error
	ValidateToken(token string) error

	GetItem(itemId uuid.UUID) (models.Item, error)
	ListItems(username string) ([]models.Item, error)
	CreateItem(username string, req models.ItemRequest) (models.Item, error)
	UpdateItem(username string, itemId uuid.UUID, req models.ItemRequest) (models.Item, error)
	DeleteItem(username string, itemId uuid.UUID) error
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

	items := router.Group("/items", h.userIdentity)
	{
		items.GET("/:itemId", h.getItem)
		items.GET("/all", h.listAllItems)
		items.GET("/all/my", h.listAllMyItems)
		items.GET("/all/:username", h.listAllUserItems)
		items.POST("/add", h.createItem)
		items.PUT("/update/:productId", h.updateItem)
		items.DELETE("/remove/:productId", h.deleteItem)
	}

	return router
}
