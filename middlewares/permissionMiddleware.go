package middlewares

import (
	"errors"
	"strconv"

	"github.com/Rajil1213/go-admin/database"
	"github.com/Rajil1213/go-admin/models"
	"github.com/Rajil1213/go-admin/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")

	id, err := util.ParseJwt(cookie)

	if err != nil {
		return err
	}

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	database.DB.Preload("Permissions").Find(&role)

	if c.Method() == "GET" {
		for _, perm := range role.Permissions {
			if perm.Name == "view_"+page || perm.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, perm := range role.Permissions {
			if perm.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("Unauthorized")
}
