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

	// edit own user info
	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	// User CRUD
	app.Post("/api/users", controllers.CreateUser)       // C
	app.Get("/api/users", controllers.AllUsers)          // R
	app.Get("/api/users/:id", controllers.GetUser)       // R
	app.Put("/api/users/:id", controllers.UpdateUser)    // U
	app.Delete("/api/users/:id", controllers.DeleteUser) // D

	// Role CRUD
	app.Post("/api/roles", controllers.CreateRole)       // C
	app.Get("/api/roles", controllers.AllRoles)          // R
	app.Get("/api/roles/:id", controllers.GetRole)       // R
	app.Put("/api/roles/:id", controllers.UpdateRole)    // U
	app.Delete("/api/roles/:id", controllers.DeleteRole) // D

	// Product CRUD
	app.Post("/api/products", controllers.CreateProduct)       // C
	app.Get("/api/products", controllers.AllProducts)          // R
	app.Get("/api/products/:id", controllers.GetProduct)       // R
	app.Put("/api/products/:id", controllers.UpdateProduct)    // U
	app.Delete("/api/products/:id", controllers.DeleteProduct) // D

	app.Get("/api/permissions", controllers.AllPermissions)

	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")
}
