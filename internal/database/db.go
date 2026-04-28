package database

import (
	"context"
	"database/sql"
)

func InitDB() (*sql.DB, error) {
	return sql.Open("sqlite", "./cotacoes.db")
}

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	return err
}

func SaveCotacao(ctx context.Context, db *sql.DB, bid string) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO cotacoes (bid) VALUES (?)", bid)
	return err
}