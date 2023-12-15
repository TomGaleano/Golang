package controller

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TomGaleano/Distophy/blob/main/golang-test/models/"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+-]+@[a-z0-9._%+]+\.[a-z0-9._%+]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body.")
	}
	//Check if password is less than6 characters
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 characters",
		})
	}
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {

	}
}
