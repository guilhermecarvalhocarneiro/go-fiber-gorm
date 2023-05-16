package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	fiber_log "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sixfwa/fiber-gorm/database"
	"github.com/sixfwa/fiber-gorm/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.Status(200).JSON("OK")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	// Agrupando pela versão da API
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/healtcheck", welcome)
	// User endpoints
	v1.Post("/users", routes.CreateUser)
	v1.Get("/users", routes.GetUsers)
	v1.Get("/users/:id", routes.GetUser)
	v1.Delete("/users/:id", routes.DeleteUser)
	// Product endpoints
	v1.Post("/products", routes.CreateProduct)
	v1.Get("/products", routes.GetProducts)
	v1.Get("/products/:id", routes.GetProduct)
	v1.Put("/products/:id", routes.UpdateProduct)

	// Order endpoints
	v1.Post("/orders", routes.CreateOrder)
	v1.Get("/orders", routes.GetOrders)
	v1.Get("/orders/:id", routes.GetOrder)

	// Auditoria endpoints
	v1.Post("/auditoria", routes.CreateAuditoria)
	v1.Get("/auditoria", routes.GetAuditorias)
	v1.Get("/auditoria/:id", routes.GetAuditoria)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Use(fiber_log.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
