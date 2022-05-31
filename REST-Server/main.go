package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"rest-server/config"
	"rest-server/data/database"
	"rest-server/routers"
)

type Application struct {
	Config   *config.Config
	Database *sql.DB
}

func main() {

	appConfig, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(*appConfig.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// app := &Application{
	// 	Config:   appConfig,
	// 	Database: db,
	// }

	router := routers.CreateRoutes()

	fmt.Println("Server is running on port", appConfig.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", appConfig.Port), router)

	if err != nil {
		panic(err)
	}
}
