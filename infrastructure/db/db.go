package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

// New создает новое подключение к базе данных PostgreSQL
func New() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/news_ai?sslmode=disable"
	}

	config, err := pgx.ParseConnectionString(dsn)
	if err != nil {
		return nil, fmt.Errorf("невозможно разобрать строку подключения: %w", err)
	}

	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 10,
		AcquireTimeout: 30 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("невозможно создать пул соединений: %w", err)
	}

	db := stdlib.OpenDBFromPool(connPool)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("невозможно проверить соединение с базой данных: %w", err)
	}

	return db, nil
}
