package repository

import (
	"context"

	"github.com/lib/pq"
	"github.com/ursulgwopp/azamon/internal/errors"
	"github.com/ursulgwopp/azamon/internal/models"
)

func (r *PostgresRepository) SignUp(req models.SignUpRequest) (models.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var profile models.Profile

	query := `INSERT INTO users (username, email, hash_password) VALUES ($1, $2, $3) RETURNING id, username, email, balance, items_list`
	if err := r.db.QueryRowContext(ctx, query, req.Username, req.Email, req.Password).Scan(&profile.Id, &profile.Username, &profile.Email, &profile.Balance, pq.Array(&profile.ItemsList)); err != nil {
		return models.Profile{}, err
	}

	return profile, nil
}

func (r *PostgresRepository) SignIn(req models.SignInRequest) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var username string

	query := `SELECT username FROM users WHERE username = $1 AND hash_password = $2`
	if err := r.db.QueryRowContext(ctx, query, req.Username, req.Password).Scan(&username); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", errors.ErrInvalidUsernameOrPassword
		}

		return "", err
	}

	return username, nil
}

func (r *PostgresRepository) SignOut(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `INSERT INTO blacklist (token) VALUES ($1)`
	_, err := r.db.ExecContext(ctx, query, token)

	return err
}

func (r *PostgresRepository) ValidateToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM blacklist WHERE token = $1)`
	if err := r.db.QueryRowContext(ctx, query, token).Scan(&exists); err != nil {
		return err
	}

	if exists {
		return errors.ErrInvalidToken
	}

	return nil
}
