package routes

import (
	"github.com/Rajil1213/go-admin/controllers"
	"github.com/Rajil1213/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// public routes
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	// private routes, authenticatd by middleware
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/", controllers.Hello)

	// User CRUD
	app.Post("/api/users", controllers.CreateUser)       // C
	app.Get("/api/users", controllers.AllUsers)          // R
	app.Get("/api/users/:id", controllers.GetUser)       // R
	app.Put("/api/users/:id", controllers.UpdateUser)    // U
	app.Delete("/api/users/:id", controllers.DeleteUser) // D
}
