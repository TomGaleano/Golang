package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/TomGaleano/Golang/database"
	"github.com/TomGaleano/Golang/models"
	"github.com/TomGaleano/Golang/util"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
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

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 6
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id = ?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body.")
	}
	blogpost.Id = uint(id)
	database.DB.Model(&blogpost).Updates(blogpost)
	return c.JSON(fiber.Map{
		"message": "Post updated succesfully.",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id = ?", id).Preload("User").Find(&blog)
	return c.JSON(fiber.Map{
		"data": blog,
	})
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blogpost := models.Blog{
		Id: uint(id),
	}
	deleteQuery := database.DB.Delete(&blogpost)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Post not found.",
		})
	} else if deleteQuery.RowsAffected == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Post not found.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post deleted succesfully.",
	})

}
