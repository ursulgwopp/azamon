package service

import "github.com/ursulgwopp/azamon/internal/models"

type Repository interface {
	SignUp(req models.SignUpRequest) (models.Profile, error)
	SignIn(req models.SignInRequest) (int, error)
	SignOut(token string) error
	ValidateToken(token string) error

	CheckUsernameExists(username string) error
	CheckEmailExists(email string) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
