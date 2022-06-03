package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/config"
	"server/models"
)

type application struct {
	config  config.Config
	models  models.Models
	logger  *log.Logger
	handler http.Handler
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	cfg := config.LoadConfig()
	fmt.Println(cfg)

	db, err := config.OpenDB(cfg)

	if err != nil {
		logger.Println("Failed connecting to database")
	}

	fmt.Println(db)
	logger.Println("Success connecting to database")

	model := models.NewModels(db)

	logger.Println(model)
}
