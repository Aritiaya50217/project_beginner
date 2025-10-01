package repository

import (
	"context"
	"log"
	"los-service/internal/domain"

	"github.com/jackc/pgx/v5"
)

type PostgresApplicationRepo struct {
	Conn *pgx.Conn
}

func NewPostgresApplication(conn *pgx.Conn) *PostgresApplicationRepo {
	return &PostgresApplicationRepo{Conn: conn}
}

func (r *PostgresApplicationRepo) Insert(app *domain.Application) (int, error) {
	var id int
	err := r.Conn.QueryRow(context.Background(),
		`INSERT INTO applications (customer_id, amount, term, status)
         VALUES ($1, $2, $3, $4) RETURNING id`,
		app.CustomerID, app.Amount, app.Term, app.Status).Scan(&id)
	if err != nil {
		log.Printf("failed to insert application: %v", err)
		return 0, err
	}
	return id, nil
}

func (r *PostgresApplicationRepo) FindByID(id int64) (*domain.Application, error) {
	row := r.Conn.QueryRow(context.Background(), "SELECT id,customer_id,status FROM applications WHERE id=$1", id)
	var app domain.Application
	if err := row.Scan(&app.ID, *&app.CustomerID, &app.Status); err != nil {
		return nil, err
	}
	return &app, nil
}
