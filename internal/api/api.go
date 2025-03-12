package api

import (
	"ApiService/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	Service service.Service
}

func NewRouter(r *Router, token string) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET, POST, PUT, DELETE",
		AllowHeaders:  "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-ID",
		ExposeHeaders: "Link",
		MaxAge:        300,
	}))
	app.Post("/task", r.Service.CreateTask)
	app.Get("/tasks", r.Service.GetTasks)
	app.Get("/task/:id", r.Service.GetTaskById)
	app.Put("/task/:id", r.Service.UpdateTask)
	app.Delete("/task/:id", r.Service.DeleteTask)
	return app
}
