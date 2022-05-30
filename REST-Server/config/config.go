package config

import "rest-server/utils/env"

type Config struct {
	Environment string
	Port        string
	Database    *Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

func NewConfig() (*Config, error) {
	env.CheckENV()
	port := env.MustGet("PORT")
	if port == "" {
		port = "4040"
	}

	return &Config{
		Environment: env.MustGet("ENV"),
		Port:        port,
		Database: &Database{
			env.MustGet("DBHOST"),
			env.MustGet("DBPORT"),
			env.MustGet("DBUSER"),
			env.MustGet("DBNAME"),
			env.MustGet("DBPASS"),
		},
	}, nil

}
