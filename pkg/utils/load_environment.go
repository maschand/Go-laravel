package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func boot() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Config Cant load", err)
	}
}

func LoadEnvironment() {
	boot()
}
