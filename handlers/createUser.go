package handlers

import (
	"github.com/Ckefa/gotham/db"
	"github.com/Ckefa/gotham/models"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	phone := c.FormValue("phone")

	user := models.NewUser(name, phone)
	db.DB.Create(user)

	var users []models.User
	db.DB.Find(&users)

	data["Users"] = users

	return c.Render(200, "users", data)
}
