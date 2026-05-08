package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection(cnf *config.DbConfig) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.DbName,
		cnf.SslMode,
	)
	return connectionString
}

func NewConnection(cnf *config.DbConfig) (*sqlx.DB, error) {
	dbSource := GetConnection(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}
