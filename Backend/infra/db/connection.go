package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection() string {
	// user -> admin
	// password -> 1234
	// dbname -> ecommerce
	// host -> localhost
	// port -> 5432
	return "user=admin password=1234 host=localhost dbname=ecommerce port=5432 sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnection()
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}
