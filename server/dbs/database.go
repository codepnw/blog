package dbs

import (
	"log"

	"github.com/codepnw/blog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBCoonn *gorm.DB

func ConnectDB() {
	dsn := "postgres://postgres:123456@localhost:4444/blog"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connection failed.")
	}

	log.Println("Database connection successful.")

	db.AutoMigrate(new(models.Blog))

	DBCoonn = db
}