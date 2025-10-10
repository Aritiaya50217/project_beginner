package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func ConnectPostgres() *pgx.Conn {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")

	if host == "" || port == "" || user == "" || pass == "" || db == "" {
		log.Fatal("Postgres environment variables are not set")
	}

	// hide password in log
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, db)
	// dsnSafe := fmt.Sprintf("postgres://%s:****@%s:%s/%s", user, host, port, db)
	fmt.Println(dsn)

	var conn *pgx.Conn
	var err error
	for i := 0; i < 10; i++ {
		// log.Printf("Trying to connect to Postgres: %s", dsnSafe)
		conn, err = pgx.Connect(context.Background(), dsn)
		if err == nil {
			log.Println("Connected to Postgres!")
			return conn
		}
		log.Printf("Retry Postgres connection... (%d/10), error: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatalf("failed to connect to postgres after retries: %v", err)
	return nil
}
