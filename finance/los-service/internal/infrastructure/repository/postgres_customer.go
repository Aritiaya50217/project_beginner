package repository

import (
	"context"
	"los-service/internal/domain"

	"github.com/jackc/pgx/v5"
)

type PostgresCustomer struct {
	Conn *pgx.Conn
}

func NewPostgresCustomer(conn *pgx.Conn) *PostgresCustomer {
	return &PostgresCustomer{Conn: conn}
}

func (r *PostgresCustomer) Save(c *domain.Customer) error {
	_, err := r.Conn.Exec(context.Background(), "INSERT INTO customers (id,name,email) VALUES ($1,$2,$3)", c.ID, c.Name, c.Email)
	return err
}

func (r *PostgresCustomer) FindByID(id int64) (*domain.Customer, error) {
	row := r.Conn.QueryRow(context.Background(), "SELECT id,name,email FROM customers WHERE id=$1", id)
	var c domain.Customer
	if err := row.Scan(&c.ID, &c.Name, &c.Email); err != nil {
		return nil, err
	}
	return &c, nil
}
