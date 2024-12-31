package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/azamon/internal/errors"
	"github.com/ursulgwopp/azamon/internal/models"
)

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}

func validateUsername(s *Service, username string) error {
	if len(username) < 2 || len(username) > 30 {
		return errors.ErrInvalidUsername
	}

	if matched, _ := regexp.MatchString(`[a-zA-Z0-9-]+`, username); !matched {
		return errors.ErrInvalidUsername
	}

	if err := s.repo.CheckUsernameExists(username); err != nil {
		return err
	}

	return nil
}

func validateEmail(s *Service, email string) error {
	if len(email) < 1 || len(email) > 50 {
		return errors.ErrInvalidEmail
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email); !matched {
		return errors.ErrInvalidEmail
	}

	if err := s.repo.CheckEmailExists(email); err != nil {
		return err
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 6 || len(password) > 100 {
		return errors.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
		return errors.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
		return errors.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
		return errors.ErrInvalidPassword
	}

	return nil
}

func (s *Service) SignUp(req models.SignUpRequest) (models.Profile, error) {
	if err := validateUsername(s, req.Username); err != nil {
		return models.Profile{}, err
	}

	if err := validateEmail(s, req.Email); err != nil {
		return models.Profile{}, err
	}

	if err := validatePassword(req.Password); err != nil {
		return models.Profile{}, err
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.SignUp(req)
}

func (s *Service) SignIn(req models.SignInRequest) (string, error) {
	req.Password = generatePasswordHash(req.Password)

	id, err := s.repo.SignIn(req)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id: id,
	})

	return jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s *Service) SignOut(token string) error {
	return s.repo.SignOut(token)
}
