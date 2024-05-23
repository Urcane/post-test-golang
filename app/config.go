package app

import (
	"github.com/joho/godotenv"
	"github.com/urcane/post-test-golang/helper"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewConfig() *Config {
	err := godotenv.Load()
	helper.PanicIfError(err)

	return &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
