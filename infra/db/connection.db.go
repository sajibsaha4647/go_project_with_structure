package db

import (
	"ecommerce/config"
	"fmt"

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
