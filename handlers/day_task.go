package handlers

import (
	"github.com/cledupe/api-generic-go/database"
	"github.com/cledupe/api-generic-go/models"
	"github.com/gofiber/fiber/v2"
)

func ListAllFact(c *fiber.Ctx) error {
	facts := []models.DayTask{}
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}

func CreatFact(c *fiber.Ctx) error {
	fact := new(models.DayTask)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}
