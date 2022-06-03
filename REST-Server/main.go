package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/config"
	"server/models"
	"server/routers"
	"time"
)

type Application struct {
	config config.Config
	models models.Models

	handler http.Handler
}

func main() {
	log.New(os.Stdout, "", log.Ldate|log.Ltime)
	log.Println("Prepare to start")

	cfg := config.LoadConfig()
	log.Println("Loaded config")

	db, err := config.OpenDB(cfg)

	if err != nil {
		log.Println("Failed connecting to database")
	}

	log.Println("Connected to database")

	model := models.NewModels(db)

	router := routers.CreateRoute(model)

	log.Println("Server is running on port", cfg.Port)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic("Failed to run server")
	}

}
