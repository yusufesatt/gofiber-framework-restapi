package services

import (
	"errors"
	"fiber-rest/dal"
	"fiber-rest/types"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var validate = validator.New()

func CreateTodo(c fiber.Ctx) error {
	t := new(types.TodoCreateDTO)

	if err := c.Bind().Body(t); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	if err := validate.Struct(t); err != nil {
		valErr := err.(validator.ValidationErrors)[0]
		message := fmt.Sprintf("Field: '%s', failed on '%s' with your value '%s'", valErr.Field(), valErr.Tag(), valErr.Value())

		return c.Status(400).JSON(fiber.Map{
			"message": message,
		})
	}

	newTodo := dal.Todo{
		Title: t.Title,
	}

	if res := dal.CreateTodo(&newTodo); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create todo",
			"success": false,
		})
	}

	return c.JSON(fiber.Map{"message": "Todo added successfully", "success": true})
}

func GetTodos(c fiber.Ctx) error {
	todos := []types.TodoResponse{}

	res := dal.GetTodos(&todos)

	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to fetch todos",
			"success": false,
		})
	}

	return c.JSON(todos)
}

func GetTodoByID(c fiber.Ctx) error {
	todoID := c.Params("todoID")

	d := types.TodoResponse{}

	res := dal.GetTodoByID(todoID, &d)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Todo not found",
				"success": false,
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to fetch todo",
			"success": false,
		})
	}
	return c.JSON(d)
}

func UpdateTodo(c fiber.Ctx) error {
	todoID := c.Params("todoID")

	t := new(types.TodoUpdateDTO)

	if err := c.Bind().Body(t); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	if err := validate.Struct(t); err != nil {
		valErr := err.(validator.ValidationErrors)[0]
		message := fmt.Sprintf("Field: '%s', failed on '%s' with your value '%s'", valErr.Field(), valErr.Tag(), valErr.Value())

		return c.Status(400).JSON(fiber.Map{
			"message": message,
		})
	}

	res := dal.UpdateTodo(todoID, t)

	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update todo",
			"success": false,
		})
	}

	if res.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
			"success": false,
		})
	}

	return c.JSON(fiber.Map{"message": "Todo updated successfully", "success": true})
}

func DeleteTodo(c fiber.Ctx) error {
	todoID := c.Params("todoID")

	res := dal.DeleteTodo(todoID)

	if res.Error != nil || res.RowsAffected == 0 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete todo",
			"success": false,
		})
	}

	return c.JSON(fiber.Map{"message": "Todo deleted successfully", "success": true})

}
