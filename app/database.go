package app

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB = nil

func ConnectToDatabase(user string, password string, host string, port string, dbName string) error {
	sourceStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable host=%s password=%s port=%s", user, dbName, host, password, port)
	db, err := sqlx.Connect("postgres", sourceStr)
	if err != nil {
		return err
	}

	DB = db
	return nil
}