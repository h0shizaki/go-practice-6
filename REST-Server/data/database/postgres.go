package database

import (
	"database/sql"
	"fmt"
	"rest-server/config"

	_ "github.com/lib/pq"
)

//DB connection

func Connect(dbConfig config.Database) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.DB),
	)

	if err != nil {
		fmt.Println("Error connecting to database err", err)
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Error connecting to database err", err)
		return nil, err
	}

	return db, nil
}
