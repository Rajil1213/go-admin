package controllers

import (
	"github.com/Rajil1213/go-admin/database"
	"github.com/Rajil1213/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}
