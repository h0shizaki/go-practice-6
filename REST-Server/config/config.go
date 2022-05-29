package config

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

//NewConfig()
