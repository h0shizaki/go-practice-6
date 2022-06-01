package config

import (
	"database/sql"
	"fmt"
	"server/utils/env"

	_ "github.com/lib/pq"
)

type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

var DB *sql.DB

func init() {

	env.CheckENV()

	dbConfig := &Database{
		env.MustGet("DBHOST"),
		env.MustGet("DBPORT"),
		env.MustGet("DBUSER"),
		env.MustGet("DBNAME"),
		env.MustGet("DBPASS"),
	}

	DB, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.DB),
	)

	if err != nil {
		panic("Error connecting to database")
	}

	err = DB.Ping()

	if err != nil {
		panic("Error connecting to database")
	}

}
