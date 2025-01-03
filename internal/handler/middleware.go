package handler

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/azamon/internal/errors"
	"github.com/ursulgwopp/azamon/internal/models"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		// models.NewErrorResponse(c, http.StatusUnauthorized, errors.ErrEmptyAuthHeader.Error())
		return
	}

	token, _ := strings.CutPrefix(header, "Bearer ")

	if err := h.service.ValidateToken(token); err != nil {
		// models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	username, err := parseToken(token)
	if err != nil {
		// models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("token", token)
	c.Set("username", username)
}

func parseToken(token string) (string, error) {
	token_, err := jwt.ParseWithClaims(token, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidSigningMethod
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token_.Claims.(*models.TokenClaims)
	if !ok {
		return "", errors.ErrInvalidTokenClaims
	}

	return claims.Username, nil
}
