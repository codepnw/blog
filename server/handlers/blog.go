package handlers

import (
	"log"

	"github.com/codepnw/blog/dbs"
	"github.com/codepnw/blog/models"
	"github.com/gofiber/fiber/v2"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Blog List",
	}

	db := dbs.DBCoonn
	var records []models.Blog

	db.Find(&records)
	context["blog_records"] = records
	return c.Status(fiber.StatusOK).JSON(context)
}

func BlogDetail(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")
	var record models.Blog

	dbs.DBCoonn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."
		return c.Status(fiber.StatusNotFound).JSON(context)
	}

	context["record"] = record
	context["statusText"] = "OK"
	context["msg"] = "Blog Detail"
	return c.Status(200).JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Create Blog",
	}

	record := new(models.Blog)

	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error in file upload.", err)
	}

	if file.Size > 0 {
		filename := "./static/uploads/" + file.Filename

		if err := c.SaveFile(file, filename); err != nil {
			log.Println("Error in file uploading...", err)
		}

		record.Image = filename
	}

	result := dbs.DBCoonn.Create(record)

	if result.Error != nil {
		log.Println("Error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	context["msg"] = "Record is saved successfully."
	context["data"] = record
	return c.Status(fiber.StatusCreated).JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Update Blog",
	}

	id := c.Params("id")
	var record models.Blog

	dbs.DBCoonn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		context["statusText"] = ""
		context["msg"] = "Record not Found."
		return c.Status(fiber.StatusNotFound).JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Record bad request."
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	result := dbs.DBCoonn.Save(&record)
	if result.Error != nil {
		log.Println("Error in saving data.")
	}

	context["msg"] = "Record updated successfully."
	context["data"] = record
	return c.Status(fiber.StatusOK).JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")
	var record models.Blog

	dbs.DBCoonn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."
		return c.Status(fiber.StatusNotFound).JSON(context)
	}

	result := dbs.DBCoonn.Delete(&record)
	if result.Error != nil {
		context["msg"] = "Something went wrong."
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	context["statusText"] = "OK"
	context["msg"] = "Record deleted successfully."
	return c.Status(fiber.StatusOK).JSON(context)
}
