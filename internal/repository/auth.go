package repository

import (
	"context"

	"github.com/ursulgwopp/azamon/internal/errors"
	"github.com/ursulgwopp/azamon/internal/models"
)

func (r *PostgresRepository) SignUp(req models.SignUpRequest) (models.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var profile models.Profile

	query := `INSERT INTO users (username, email, hash_password) VALUES ($1, $2, $3) RETURNING id, username, email, balance, items_list`
	if err := r.db.QueryRowContext(ctx, query, req.Username, req.Email, req.Password).Scan(&profile.Id, &profile.Username, &profile.Email, &profile.Balance, &profile.ItemsList); err != nil {
		return models.Profile{}, err
	}

	return profile, nil
}

func (r *PostgresRepository) SignIn(req models.SignInRequest) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var id int

	query := `SELECT id FROM users WHERE username = $1 AND hash_password = $2`
	if err := r.db.QueryRowContext(ctx, query, req.Username, req.Password).Scan(&id); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return -1, errors.ErrInvalidUsernameOrPassword
		}

		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) SignOut(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `INSERT INTO blacklist (token) VALUES ($1)`
	_, err := r.db.ExecContext(ctx, query, token)

	return err
}
