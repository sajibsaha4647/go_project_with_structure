package db

import (
	"ecommerce/config"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
func getConnectionString(cnf *config.Config) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		cnf.DBUser,
		cnf.DBPass,
		cnf.DbName,
		cnf.DBPort,
		cnf.DBHost,
	)
}

func ConnectDB(cnf *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", getConnectionString(cnf))
	return db, err
}

func RunMigrations(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DbName,
	)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}