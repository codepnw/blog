package main

import (
	"github.com/codepnw/blog/dbs"
	"github.com/gofiber/fiber/v2"
)

func init() {
	dbs.ConnectDB()
}

func main() {
	sqlDB, err := dbs.DBCoonn.DB()
	if err != nil {
		panic("Error in sql connection.")
	}
	defer sqlDB.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to My Blog"})
	})

	app.Listen(":8080")
}