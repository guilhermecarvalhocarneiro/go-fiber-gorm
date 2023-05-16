package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-gorm/database"
	"github.com/sixfwa/fiber-gorm/models"
)

func CreateAuditoria(c *fiber.Ctx) error {
	var auditoria models.Auditoria
	if err := c.BodyParser(&auditoria); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.Database.Db.Create(&auditoria)
	return c.Status(200).JSON(auditoria)
}

func GetAuditorias(c *fiber.Ctx) error {
	auditorias := []models.Auditoria{}
	database.Database.Db.Find(&auditorias)
	return c.Status(200).JSON(auditorias)
}

func GetAuditoria(c *fiber.Ctx) error {
	id := c.Params("id")
	auditoria := models.Auditoria{}
	database.Database.Db.Find(&auditoria, id)
	return c.Status(200).JSON(auditoria)
}
