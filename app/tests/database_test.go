package tests

import (
	"fmt"
	"github.com/chand19-af/digitels-template/app/models"
	"github.com/chand19-af/digitels-template/platform/database"
	"github.com/chand19-af/digitels-template/platform/migrations"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

func boot() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Config Cant load", err)
	}
}

func TestConnection(t *testing.T) {
	boot()
	db := database.GormMysqlConnection()
	if db.Error != nil {
		panic(db.Error)
	}
	print(db)
}

func TestRunMigration(t *testing.T) {
	boot()
	db := database.GormMysqlConnection()
	if db.Error != nil {
		panic(db.Error)
	}
	migrations.AutoMigration(*db)

	print("migration success")
}

func TestCreate(t *testing.T) {
	boot()
	db := database.GormMysqlConnection()
	if db.Error != nil {
		panic(db.Error)
	}
	res := db.Create(&models.Hotel{
		Model:    gorm.Model{},
		Name:     "First Hotel",
		CheckIn:  time.Time{},
		CheckOut: time.Time{},
		About:    "-",
		Logo:     "-",
		Status:   "0",
		Timezone: "Makassar",
	})

	if res.Error != nil {
		panic("cant create data!")
	}

	fmt.Println(res)
}

func TestUpdate(t *testing.T) {
	boot()
	db := database.GormMysqlConnection()
	if db.Error != nil {
		panic(db.Error)
	}

	res := db.First(&models.Hotel{
		Model: gorm.Model{
			ID: 1,
		},
	}).Updates(&models.Hotel{
		Name: "Siti Kotijah",
	})

	if res.Error != nil {
		panic("cant update data")
	}
}

func TestGet(t *testing.T) {
	boot()
	db := database.GormMysqlConnection()
	if db.Error != nil {
		panic(db.Error)
	}

	res := db.First(&models.Hotel{}, "id = ?", 1)

	fmt.Print(res.Get("*"))
}
