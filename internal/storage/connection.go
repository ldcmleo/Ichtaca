package storage

import (
	"database/sql"
	"fmt"

	"github.com/ldcmleo/Ichtaca/internal/config"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PgHost,
		cfg.PgPort,
		cfg.PgUser,
		cfg.PgPassword,
		cfg.PgDatabase,
		cfg.PgSSLMode,
	)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return database, nil
}
