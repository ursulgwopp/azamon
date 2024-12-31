package repository

import (
	"context"

	"github.com/ursulgwopp/azamon/internal/errors"
)

func (r *PostgresRepository) CheckUsernameExists(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	if err := r.db.QueryRowContext(ctx, query, username).Scan(&exists); err != nil {
		return err
	}

	if exists {
		return errors.ErrUsernameExists
	}

	return nil
}

func (r *PostgresRepository) CheckEmailExists(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	if err := r.db.QueryRowContext(ctx, query, email).Scan(&exists); err != nil {
		return err
	}

	if exists {
		return errors.ErrEmailExists
	}

	return nil
}
