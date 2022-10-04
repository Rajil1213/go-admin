package controllers

import (
	"strconv"

	"github.com/Rajil1213/go-admin/database"
	"github.com/Rajil1213/go-admin/middlewares"
	"github.com/Rajil1213/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

const DefaultPassword = "changeme"

func AllUsers(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

/* This function is for when a registered user creates another user */
func CreateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}
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

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
