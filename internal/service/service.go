package service

import (
	"github.com/google/uuid"
	"github.com/ursulgwopp/azamon/internal/models"
)

type Repository interface {
	SignUp(req models.SignUpRequest) (models.Profile, error)
	SignIn(req models.SignInRequest) (string, error)
	SignOut(token string) error
	ValidateToken(token string) error

	GetItem(itemId uuid.UUID) (models.Item, error)
	ListItems(username string) ([]models.Item, error)
	CreateItem(username string, req models.ItemRequest) (models.Item, error)
	UpdateItem(itemId uuid.UUID, req models.ItemRequest) (models.Item, error)
	DeleteItem(itemId uuid.UUID) error

	CheckUsernameExists(username string) error
	CheckEmailExists(email string) error
	CheckItemIdExists(itemId uuid.UUID) error
	CheckItemSeller(itemId uuid.UUID) (string, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
