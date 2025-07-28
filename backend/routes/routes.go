package routes

import (
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
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	//user_record.Username = username
	bytes, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	hashed_password := string(bytes)

	var userRecord = models.UserModel{
		Username: userData.Username,
		Password: hashed_password,
	}

	result := database.DB.Create(&userRecord)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	return c.JSON(fiber.Map{"message": "Account created successfully"})
}

func Login(c *fiber.Ctx) error {
	//username := c.FormValue("username")
	//pass := c.FormValue("password")

	var userData models.UserModel

	err := c.BodyParser(&userData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	var userRecord models.UserModel

	result := database.DB.Where("username = ?", userData.Username).First(&userRecord)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Username or password is not correct"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRecord.Password), []byte(userData.Password))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Username or password is not correct"})
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
		return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
	}

	var userRecord models.UserModel

	if user := c.Locals("user").(*jwt.Token); user == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	err = database.DB.Where("username = ?", name).First(&userRecord).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Account not found"})
	}
	todo.TargetUserID = userRecord.ID

	database.DB.Create(&todo)

	return c.Status(200).JSON(fiber.Map{"message": "todo created successfully"})
}

func GetTodo(c *fiber.Ctx) error {
	var todo []models.Todo

	if user := c.Locals("user").(*jwt.Token); user == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	var userRecord models.UserModel
	result := database.DB.Where("username = ?", name).First(&userRecord)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Account not found"})
	}

	result = database.DB.Where("target_user_id = ?", userRecord.ID).Find(&todo)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Account not found"})
	}

	return c.Status(200).JSON(fiber.Map{"todos": todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	var targetID = c.Params("id")
	var todo models.Todo

	if user := c.Locals("user").(*jwt.Token); user == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	var userRecord models.UserModel

	result := database.DB.Where("username = ?", name).First(&userRecord)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Account not found"})
	}

	todo.TargetUserID = userRecord.ID

	result = database.DB.First(&todo, targetID).Delete(targetID)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Todo deleted successfully"})
}

func UpdateTodo(c *fiber.Ctx) error {
	var reqTodo models.Todo
	var todo models.Todo
	err := c.BodyParser(&reqTodo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
	}

	if user := c.Locals("user").(*jwt.Token); user == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	var userRecord models.UserModel
	result := database.DB.Where("username = ?", name).First(&userRecord)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Account not found"})
	}

	database.DB.First(&todo, reqTodo.ID)
	if todo.TargetUserID != userRecord.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	todo.ID = reqTodo.ID
	todo.Title = reqTodo.Title
	todo.Done = reqTodo.Done
	todo.TargetUserID = userRecord.ID

	database.DB.Save(&todo)

	return c.Status(200).JSON(fiber.Map{"message": "Todo updated successfully"})
}

func DeleteAccount(c *fiber.Ctx) error {
	if user := c.Locals("user").(*jwt.Token); user == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	var userRecord models.UserModel

	result := database.DB.Where("username = ?", name).Delete(&userRecord)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Account not found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Account deleted successfully"})
}

func ChangePassword(c *fiber.Ctx) error {
	var reqBody models.ChangePassword
	err := c.BodyParser(&reqBody)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	if user := c.Locals("user").(*jwt.Token); user == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized no token"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	var userRecord models.UserModel

	result := database.DB.Where("username = ?", name).First(&userRecord)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Account not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRecord.Password), []byte(reqBody.CurrentPassword))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized password"})
	} else {
		bytes, err := bcrypt.GenerateFromPassword([]byte(reqBody.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		hashed_password := string(bytes)
		userRecord.Password = hashed_password
		database.DB.Save(&userRecord)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Password changed"})
}

func CheckToken(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok || user == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	var userRecord models.UserModel
	if err := database.DB.Where("username = ?", name).First(&userRecord).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Token is valid"})
}
