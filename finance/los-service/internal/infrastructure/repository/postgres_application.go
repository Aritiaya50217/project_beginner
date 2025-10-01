package repository

import (
	"context"
	"los-service/internal/domain"

	"github.com/jackc/pgx/v5"
)

type PostgresApplicationRepo struct {
	Conn *pgx.Conn
}

func NewPostgresApplication(conn *pgx.Conn) *PostgresApplicationRepo {
	return &PostgresApplicationRepo{Conn: conn}
}

func (r *PostgresApplicationRepo) Save(app *domain.Application) error {
	_, err := r.Conn.Exec(context.Background(), "INSERT INTO applications (id,customer_id,status) VALUES ($1,$2,$3)", app.ID, app.CustomerID, app.Status)
	return err
}

func (r *PostgresApplicationRepo) FindByID(id int64) (*domain.Application, error) {
	row := r.Conn.QueryRow(context.Background(), "SELECT id,customer_id,status FROM applications WHERE id=$1", id)
	var app domain.Application
	if err := row.Scan(&app.ID, *&app.CustomerID, &app.Status); err != nil {
		return nil, err
	}
	return &app, nil
}
