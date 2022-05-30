package main

import (
	"fmt"
	"net/http"
	"rest-server/config"
	"rest-server/data/database"
	"rest-server/routers"
)

func main() {
	fmt.Println("Hi")

	appConfig, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(*appConfig.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := routers.CreateRoutes()

	fmt.Println("Server is running on port", appConfig.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", appConfig.Port), router)

	if err != nil {
		panic(err)
	}
}
