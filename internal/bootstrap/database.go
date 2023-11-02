package bootstrap

import (
	"database/sql"
	"fmt"
	"simple-checkout-app/internal/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase(config *config.Config) (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", DbStringURL(config))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected")

	return db, nil
}

func GetDB() *sql.DB {
	return db
}

func DbStringURL(config *config.Config) string {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Databasename,
		config.Password,
	)

	return connectionString
}
