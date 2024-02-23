package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PgConfig struct {
	user     string
	password string
	dbname   string
	host     string
	port     string
}

func isDev() bool {
	env, exists := os.LookupEnv("ENV")
	return exists && env == "dev"
}

func GetDbConfig() *PgConfig {
	err := godotenv.Load("pg.env")
	if err != nil {
		log.Println("no pg.env provided")
	}

	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		user = "postgres"
	}
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		password = "postgres"
	}
	dbname, exists := os.LookupEnv("POSTGRES_DB_NAME")
	if !exists {
		dbname = "postgres"
	}
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = "localhost"
	}
	port, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		port = "5432"
	}

	return &PgConfig{
		user:     user,
		password: password,
		dbname:   dbname,
		host:     host,
		port:     port,
	}
}

func GetTgToken() string {
	err := godotenv.Load("telegram-token.env")
	if err != nil {
		log.Println("no telegram-token.env provided", err)
	}
	token, ok := os.LookupEnv("TELEBOTTOKEN")
	if !ok {
		log.Fatal("TELEBOTTOKEN env must be provided")
	}
	return token
}
