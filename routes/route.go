package routes

import (
	"github.com/TomGaleano/Golang/controller"
	"github.com/TomGaleano/Golang/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, https://your-deployed-frontend.com",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticated)

	app.Post("/api/createpost", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/detailpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Get("/api/uniquepost", controller.UniquePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Post("/api/upload", controller.Upload)
	app.Static("/api/uploads", "./uploads")
}
