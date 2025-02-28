package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PgConfig struct {
	user     string
	password string
	dbname   string
	host     string
	port     string
}

func IsDev() bool {
	env, exists := os.LookupEnv("ENV")
	return exists && env == "dev"
}

func GetDbConfig() *PgConfig {
	err := godotenv.Load("pg.env")
	if err != nil {
		log.Println("no pg.env provided, getting from ENV vars")
	}

	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		user = "postgres"
	}
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		password = "postgres"
	}
	dbname, exists := os.LookupEnv("POSTGRES_DB")
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
		log.Println("no telegram-token.env provided, getting token from ENV vars")
	}
	token, ok := os.LookupEnv("TELEBOTTOKEN")
	if !ok {
		log.Fatal("TELEBOTTOKEN env must be provided")
	}
	return token
}

func GetPGConnectionString() string {
	conf := GetDbConfig()
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		conf.user,
		conf.password,
		conf.host,
		conf.port,
		conf.dbname,
	)
	// urlExample := "postgres://username:password@localhost:5432/database_name"
}

type ServerPortsCfg struct {
	GrpcPort int
	HttpPort int
}

func GetServerPorts() *ServerPortsCfg {
	err := godotenv.Load("server.env")
	if err != nil {
		log.Println("no server.env provided, getting from ENV vars")
	}
	grpcportStr, exists := os.LookupEnv("grpcport")
	if !exists {
		grpcportStr = "5555"
	}
	httpportStr, exists := os.LookupEnv("httpport")
	if !exists {
		httpportStr = "8080"
	}

	grpcport, err := strconv.Atoi(grpcportStr)
	if err != nil {
		log.Fatal("Error reading grpc port from env")
	}

	httpport, err := strconv.Atoi(httpportStr)
	if err != nil {
		log.Fatal("Error reading http port from env")
	}

	return &ServerPortsCfg{
		GrpcPort: grpcport,
		HttpPort: httpport,
	}
}
