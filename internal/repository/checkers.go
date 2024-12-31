package repository

import (
	"context"

	"github.com/google/uuid"
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

func (r *PostgresRepository) CheckItemIdExists(itemId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM items WHERE id = $1)`
	if err := r.db.QueryRowContext(ctx, query, itemId).Scan(&exists); err != nil {
		return err
	}

	if exists {
		return errors.ErrItemIdNotFound
	}

	return nil
}

func (r *PostgresRepository) CheckItemSeller(itemId uuid.UUID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var seller string

	query := `SELECT seller FROM items WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, itemId).Scan(&seller); err != nil {
		return "", err
	}

	return seller, nil
}
