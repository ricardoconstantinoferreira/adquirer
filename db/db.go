package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("erro ao carregar .env: %w", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao MySQL: %w", err)
	}

	return conn, nil
}
