package controllers

import (
	"strconv"

	"github.com/Rajil1213/go-admin/database"
	"github.com/Rajil1213/go-admin/models"
	"github.com/gofiber/fiber/v2"
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

	user.SetPassword(DefaultPassword)
	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Find(&user)

	return c.JSON(user)
}
