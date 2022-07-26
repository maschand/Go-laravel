package tests

import (
	"github.com/chand19-af/digitels-template/platform/database"
	"github.com/chand19-af/digitels-template/platform/migrations"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func boot() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Config Cant load", err)
	}
}

func TestConnection(t *testing.T) {
	boot()
	db, err := database.GormMysqlConnection()
	if err != nil {
		panic(err)
	}
	print(db)
	print(err)
}

func TestRunMigration(t *testing.T) {
	boot()
	db, err := database.GormMysqlConnection()
	if err != nil {
		panic(err)
	}
	migrations.AutoMigration(*db)

	print("migration success")
}
