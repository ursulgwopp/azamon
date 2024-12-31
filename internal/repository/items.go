package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ursulgwopp/azamon/internal/models"
)

func (r *PostgresRepository) GetItem(productId uuid.UUID) (models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var item models.Item

	query := `SELECT id, seller, name, description, quantity, price FROM items WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, productId).Scan(&item.Id, &item.Seller, &item.Name, &item.Description, &item.Quantity, &item.Price); err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (r *PostgresRepository) ListItems(username string) ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var items []models.Item

	query := `SELECT id, seller, name, description, quantity, price FROM items WHERE username = $1`
	rows, err := r.db.QueryContext(ctx, query, username)
	if err != nil {
		return []models.Item{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if rows.Scan(&item.Id, &item.Seller, &item.Name, &item.Description, &item.Quantity, &item.Price); err != nil {
			return []models.Item{}, err
		}

		items = append(items, item)
	}

	if rows.Err() != nil {
		return []models.Item{}, err
	}

	return items, nil
}

func (r *PostgresRepository) CreateItem(username string, req models.ItemRequest) (models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var item models.Item

	query := `INSERT INTO items (id, seller, name, description, quantity, price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, seller, name, description, quantity, price`
	if err := r.db.QueryRowContext(ctx, query, uuid.New(), username, req.Name, req.Description, req.Quantity, req.Price).Scan(&item.Id, &item.Seller, &item.Name, &item.Description, &item.Quantity, &item.Price); err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (r *PostgresRepository) UpdateItem(productId uuid.UUID, req models.ItemRequest) (models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var item models.Item

	query := `UPDATE items SET name = $1, description = $2, quantity = $3, price = $4 WHERE id = $5 RETURNING id, seller, name, description, quantity, price`
	if err := r.db.QueryRowContext(ctx, query, req.Name, req.Description, req.Quantity, req.Price, productId).Scan(&item.Id, &item.Seller, &item.Name, &item.Description, &item.Quantity, &item.Price); err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (r *PostgresRepository) DeleteItem(productId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `DELETE FROM items WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, productId)

	return err
}
