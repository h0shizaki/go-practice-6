package config

import (
	"context"
	"database/sql"
	"fmt"
	"server/log"
	"server/utils/env"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {

	env.CheckENV()

	DB, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			env.MustGet("DBHOST"),
			env.MustGet("DBPORT"),
			env.MustGet("DBUSER"),
			env.MustGet("DBPASS"),
			env.MustGet("DBNAME")),
	)

	if err != nil {
		panic("Error connecting to database")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = DB.PingContext(ctx)

	if err != nil {
		panic("Error connecting to database")
	}

	log.Log.LogInfo("Connected to database")

}
