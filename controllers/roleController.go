package controllers

import (
	"strconv"

	"github.com/Rajil1213/go-admin/database"
	"github.com/Rajil1213/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Preload("Permissions").Find(&roles)

	return c.JSON(roles)
}

// Data Transfer Object
type RoleCreateDTO struct {
	Name        string `json:"name"`
	Permissions []uint `json:"permissions"`
}

func CreateRole(c *fiber.Ctx) error {
	/* This function is for when a registered user creates another user */
	var roleDto RoleCreateDTO

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	permissions := make([]models.Permission, len(roleDto.Permissions))

	for i, permId := range roleDto.Permissions {
		permissions[i] = models.Permission{
			Id: permId,
		}
	}

	role := models.Role{
		Name:        roleDto.Name,
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permission").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto RoleCreateDTO

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	permissions := make([]models.Permission, len(roleDto.Permissions))
	for i, permId := range roleDto.Permissions {
		permissions[i] = models.Permission{
			Id: permId,
		}
	}

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto.Name,
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
