package handlers

import (
	"github.com/Ckefa/gotham/db"
	"github.com/Ckefa/gotham/models"
	"github.com/labstack/echo/v4"
)

func DeleteUser(c echo.Context) error {
	userId := c.QueryParam("userId")

	var user models.User
	db.DB.First(&user, userId)
	db.DB.Delete(&user)

	var users []models.User
	db.DB.Find(&users)

	data["Users"] = users

	return c.Render(200, "users", data)
}
