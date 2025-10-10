package repository

import (
	"context"
	"log"
	"los-service/internal/domain"

	"github.com/jackc/pgx/v5"
)

type PostgresCustomer struct {
	Conn *pgx.Conn
}

func NewPostgresCustomer(conn *pgx.Conn) *PostgresCustomer {
	return &PostgresCustomer{Conn: conn}
}

func (r *PostgresCustomer) Insert(c *domain.Customer) (int, error) {
	query := `
		INSERT INTO customers (name, email, phone)
		VALUES ($1, $2, $3) RETURNING id, created_at
	`
	err := r.Conn.QueryRow(
		context.Background(),
		query,
		c.Name,
		c.Email,
		c.Phone,
	).Scan(&c.ID, &c.CreatedAt)

	if err != nil {
		log.Printf("failed to insert customer: %v", err)
		return 0, err
	}
	return int(c.ID), err
}

func (r *PostgresCustomer) FindByID(id int64) (*domain.Customer, error) {
	row := r.Conn.QueryRow(context.Background(), "SELECT id,name,email FROM customers WHERE id=$1", id)
	var c domain.Customer
	if err := row.Scan(&c.ID, &c.Name, &c.Email); err != nil {
		return nil, err
	}
	return &c, nil
}
