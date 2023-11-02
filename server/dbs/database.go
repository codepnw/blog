package dbs

import (
	"fmt"
	"log"
	"os"

	"github.com/codepnw/blog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBCoonn *gorm.DB

func ConnectDB() {
	dbUser := os.Getenv("db_user")
	dbPass := os.Getenv("db_password")
	dbPort := os.Getenv("db_port")
	dbName := os.Getenv("db_name")

	dsn := fmt.Sprintf(`postgres://%s:%s@localhost:%s/%s`, dbUser, dbPass, dbPort, dbName)

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