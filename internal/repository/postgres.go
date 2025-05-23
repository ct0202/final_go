package repository

import (
	"context"
	"database/sql"

	"github.com/ct0202/bicycle-shop/internal/domain"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(connStr string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		"INSERT INTO orders (id, user_id, bicycle_id, total_price, status) VALUES ($1, $2, $3, $4, $5)",
		order.ID, order.UserID, order.BicycleID, order.TotalPrice, order.Status,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		"UPDATE bicycles SET available = false WHERE id = $1",
		order.BicycleID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
