package db

import (
	"log"

	"github.com/rndpk/distributed-task-queue/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=rndpk dbname=taskqueue sslmode=disable"

	database, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}

	err = database.AutoMigrate(
		&models.Task{},
	)

	if err != nil {
		log.Fatal(err)
	}

	DB = database

	log.Println("Connected to PostgreSQL")
}
