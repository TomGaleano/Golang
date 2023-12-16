package routes

import (
	"github.com/TomGaleano/Golang/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
