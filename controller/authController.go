package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/TomGaleano/Golang/database"
	"github.com/TomGaleano/Golang/models"
	"github.com/TomGaleano/Golang/util"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body.")
	}
	//Check if password is less than 6 characters
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 characters.",
		})
	}
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid email address.",
		})
	}

	//Check if email already exists
	database.DB.Where("email = ?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already in use.",
		})

	}

	//Create user
	user := models.User{
		Firstname: data["first_name"].(string),
		Lastname:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}
	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(201)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "User created succesfully.",
	})

}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	var user models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body.")
	}
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User not found.",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect password.",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Logged in succesfully.",
		"user":    user,
	})

}
