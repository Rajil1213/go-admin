package controllers

import (
	"github.com/Rajil1213/go-admin/database"
	"github.com/Rajil1213/go-admin/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const DefaultPassword = "changeme"

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	/* This function is for when a registered user creates another user */
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// default user password
	password, _ := bcrypt.GenerateFromPassword([]byte(DefaultPassword), 14)
	user.Password = password

	database.DB.Create(&user)

	return c.JSON(user)
}
