package main

import (
	"backend/database"
	"backend/routes"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	var app = fiber.New(fiber.Config{AppName: "Fiber/Vue Todo App"})

	database.ConnectDb()

	// CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173 , http://127.0.0.1:4000",
	}))

	// Auth api
	auth := app.Group("/auth")

	auth.Post("/register", routes.Register)
	auth.Post("/login", routes.Login)

	// Todo api

	// Secret key is only for test and this isn't secure!
	todos := app.Group("/todos", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret_key")},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Bad Token",
			})
		},
	}))

	todos.Post("/add", routes.AddTodo)
	todos.Get("/get", routes.GetTodo)
	todos.Delete("/delete/:id", routes.DeleteTodo)
	todos.Post("/update", routes.UpdateTodo)
	todos.Delete("/delete-account", routes.DeleteAccount)
	todos.Post("/change-password", routes.ChangePassword)
	todos.Get("/check-token", routes.CheckToken)

	app.Listen(":4000")
}
