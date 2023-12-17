package controller

import (
	"fmt"

	"github.com/TomGaleano/Golang/database"
	"github.com/TomGaleano/Golang/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body.")
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Unable to create post.",
		})

	}
	return c.JSON(fiber.Map{
		"message": "Post created succesfully.",
	})

}
