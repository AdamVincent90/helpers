package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDatabaseCreds() map[string]string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	m := map[string]string{
		"DB_USER":    os.Getenv("DB_USER"),
		"DB_PASS":    os.Getenv("DB_PASS"),
		"DB_NAME":    os.Getenv("DB_NAME"),
		"DB_ADDRESS": os.Getenv("DB_ADDRESS"),
		"DB_PORT":    os.Getenv("DB_PORT"),
		"DB":         os.Getenv("DB"),
	}

	return m
}
