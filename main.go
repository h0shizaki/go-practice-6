package main

import (
	"flag"
	"fmt"
	"go-gql/gql"
	"go-gql/postgres"
	"go-gql/server"
	"go-gql/utils/env"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

type Config struct {
	Env  string
	Port string
}

func main() {

	var config Config
	env.CheckENV()
	flag.StringVar(&config.Env, "environment", env.MustGet("ENV"), "ENV")
	flag.StringVar(&config.Port, "port", env.MustGet("PORT"), "Port")

	db, router := initAPI()
	defer db.Close()

	fmt.Println(config.Env)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), router))
}

func initAPI() (*postgres.Db, *chi.Mux) {

	router := chi.NewRouter()

	dbPort, _ := strconv.Atoi(env.MustGet("DBPORT"))
	db, err := postgres.New(
		postgres.ConnString(
			env.MustGet("DBHOST"),
			dbPort,
			env.MustGet("DBUSER"),
			env.MustGet("DBPASS"),
			env.MustGet("DBNAME"),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)

	if err != nil {
		log.Println("Error creating schema ", err)
	}

	s := server.Server{GqlSchema: &sc}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		// middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", s.GraphQL())

	return db, router
}
