package config

import (
	"os"
)

const (
	DB_NAME      string = "DB_NAME"
	DB_HOST      string = "DB_HOST"
	DB_USRERNAME string = "DB_USRERNAME"
	DB_PASSWORD  string = "DB_PASSWORD"
	DB_DIALECT   string = "DB_DIALECT"
	DB_PORT      string = "DB_PORT"
	GRPC_CLIENT  string = "PRODUCT_GRPC_CLIENT"
)

type Config struct {
	DBName     string
	DBUser     string
	Host       string
	Password   string
	Dialect    string
	Port       string
	GrpcClient string
}

func LoadEnv() *Config {
	dbName     := os.Getenv(DB_NAME)
	dbHost     := os.Getenv(DB_HOST)
	dbUsername := os.Getenv(DB_USRERNAME)
	dbPassword := os.Getenv(DB_PASSWORD)
	dbPort     := os.Getenv(DB_PORT)
	dbDialect  := os.Getenv(DB_DIALECT)
	grpcClient := os.Getenv(GRPC_CLIENT)

	return &Config{
		DBName:     dbName,
		DBUser:     dbUsername,
		Host:       dbHost,
		Password:   dbPassword,
		Dialect:    dbDialect,
		Port:       dbPort,
		GrpcClient: grpcClient,
	}
}
