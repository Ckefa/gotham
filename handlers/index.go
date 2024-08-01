package handlers

import (
	"log"
	"net/http"

	"github.com/Ckefa/gotham/db"
	"github.com/Ckefa/gotham/models"
	"github.com/labstack/echo/v4"
)

var data = make(map[string]interface{})

func Index(c echo.Context) error {
	var users []models.User

	var user models.User
	db.DB.Find(&users)

	log.Println("user found", user.ID, user.Name)

	data["Users"] = users

	return c.Render(http.StatusOK, "index", data)
}
