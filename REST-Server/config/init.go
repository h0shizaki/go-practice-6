package config

import (
	"flag"
	"server/utils/env"
)

type Config struct {
	port string
	env  string
	db   struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
	}
}

func LoadConfig() Config {
	var cfg Config

	env.CheckENV()

	port := env.MustGet("PORT")

	if port == "" {
		port = "4040"
	}

	flag.StringVar(&cfg.port, "port", port, "Server port will be listen on")
	flag.StringVar(&cfg.env, "environment", env.MustGet("ENV"), "Application environment")

	flag.StringVar(&cfg.db.Host, "database host", env.MustGet("DBHOST"), "localhost")
	flag.StringVar(&cfg.db.Port, "database port", env.MustGet("DBPORT"), "5432")
	flag.StringVar(&cfg.db.User, "database user", env.MustGet("DBUSER"), "user")
	flag.StringVar(&cfg.db.Password, "database password", env.MustGet("DBPASS"), "password")
	flag.StringVar(&cfg.db.Dbname, "database name", env.MustGet("DBNAME"), "name")
	return cfg
}
