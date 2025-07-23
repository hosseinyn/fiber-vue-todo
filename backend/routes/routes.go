package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)
import "golang.org/x/crypto/bcrypt"
import "backend/database"
import "backend/models"

// Register Authentication routes
func Register(c *fiber.Ctx) error {
	//username := c.FormValue("username")
	//pass := c.FormValue("password")

	var userData models.UserModel

	err := c.BodyParser(&userData)
	if err != nil {
		fmt.Println(err)
	}

	//user_record.Username = username
	bytes, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	hashed_password := string(bytes)

	var userRecord = models.UserModel{
		Username: userData.Username,
		Password: hashed_password,
	}

	result := database.DB.Create(&userRecord).Error
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return c.JSON(fiber.Map{"message": "Account created successfully"})
}

func Login(c *fiber.Ctx) error {
	//username := c.FormValue("username")
	//pass := c.FormValue("password")

	var userRecord models.UserModel

	err := c.BodyParser(&userRecord)
	if err != nil {
		fmt.Println(err)
	}

	result := database.DB.Where("username = ?", userRecord.Username).First(&userRecord)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Username or password is not correct"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRecord.Password), []byte(userRecord.Password))
	if err != nil {
		fmt.Println(err)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": userRecord.Username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

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
