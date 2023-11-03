package main

import (
	"log"

	"github.com/codepnw/blog/dbs"
	"github.com/codepnw/blog/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file.")
	}

	dbs.ConnectDB()
}

func main() {
	sqlDB, err := dbs.DBCoonn.DB()
	if err != nil {
		panic("Error in sql connection.")
	}
	defer sqlDB.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())
	routes.SetupRoutes(app)

	app.Listen(":8080")
}
