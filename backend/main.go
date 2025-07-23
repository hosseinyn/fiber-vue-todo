package main

import (
	"backend/database"
	"backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var app = fiber.New(fiber.Config{AppName: "Fiber/Vue Todo App"})

	database.ConnectDb()

	// Todo api

	todos := app.Group("/todos")

	todos.Post("/add", routes.AddTodo)
	todos.Get("/get", routes.GetTodo)
	todos.Delete("/delete/:id", routes.DeleteTodo)
	todos.Post("/update", routes.UpdateTodo)

	app.Listen(":4000")
}
