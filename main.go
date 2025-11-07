package main

import (
	"fiber-rest/dal"
	"fiber-rest/database"
	"fiber-rest/services"
	"log"

	"github.com/gofiber/fiber/v3"
)

func healthz(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func main() {
	database.Connect()
	database.DB.AutoMigrate(&dal.Todo{})

	app := fiber.New()

	app.Get("/", healthz)
	app.Post("/todos", services.CreateTodo)
	app.Get("/todos", services.GetTodos)
	app.Get("/todos/:todoID", services.GetTodoByID)
	app.Put("/todos/:todoID", services.UpdateTodo)
	app.Delete("/todos/:todoID", services.DeleteTodo)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
