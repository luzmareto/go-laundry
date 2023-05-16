package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	DbConfig
}

func NewConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	cfg := Config{}
	cfg.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
	return cfg
}
