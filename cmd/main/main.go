package main

import (
	"io"
	"log"
	"os"
	"text/template"

	"github.com/Ckefa/gotham/db"
	"github.com/Ckefa/gotham/handlers"
	"github.com/Ckefa/gotham/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	println("Starting app .....")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load environment variables")
	}

	addr := os.Getenv("laddr")
	if addr == "" {
		log.Fatal("Environment variable 'laddr' not set")
	}

	dsn := os.Getenv("dsn")
	if dsn == "" {
		log.Fatal("Environment variable 'dsn' not set")
	}

	err = db.Init(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	log.Println("Database connection initialized successfully")

	if db.DB == nil {
		log.Fatalf("DB is nil after initialization")
	}

	log.Println("Starting AutoMigrate...")
	err = db.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
	log.Println("AutoMigrate completed")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Static("/", "static")

	e.Renderer = newTemplate()
	e.GET("/", handlers.Index)
	e.POST("/users", handlers.CreateUser)
	e.DELETE("/users", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(addr))
}
