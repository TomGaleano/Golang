package routes

import (
	"github.com/TomGaleano/Golang/controller"
	"github.com/TomGaleano/Golang/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuthenticated)
	app.Post("/api/createpost", controller.CreatePost)
}
