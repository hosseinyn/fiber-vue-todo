package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)
import "backend/database"
import "backend/models"

func AddTodo(c *fiber.Ctx) error {
	var todo models.Todo

	err := c.BodyParser(&todo)
	if err != nil {
		fmt.Println(err)
		return err
	}

	database.DB.Create(&todo)

	return c.Status(200).JSON(fiber.Map{"message": "todo created successfully"})
}

func GetTodo(c *fiber.Ctx) error {
	var todo []models.Todo

	result := database.DB.Find(&todo)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return c.Status(200).JSON(fiber.Map{"todos": todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	var targetID = c.Params("id")
	var todo models.Todo

	result := database.DB.First(&todo, targetID).Delete(targetID)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todo deleted successfully"})
}

func UpdateTodo(c *fiber.Ctx) error {
	var reqTodo models.Todo
	var todo models.Todo
	err := c.BodyParser(&reqTodo)
	if err != nil {
		fmt.Println(err)
	}

	database.DB.First(&todo, reqTodo.ID)
	todo.ID = reqTodo.ID
	todo.Title = reqTodo.Title
	todo.Done = reqTodo.Done

	database.DB.Save(&todo)

	return c.Status(200).JSON(fiber.Map{"message": "Todo updated successfully"})
}
