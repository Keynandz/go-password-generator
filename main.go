package main

import (
	"fmt"
	"go-password/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("error loading file .env")
	}

	format := ":%s"
	port := fmt.Sprintf(format, os.Getenv("PORT"))

	e.Static("/", "views")
	e.File("/index", "views/index.html")
	routes.ApiRoutes(e)
	e.Logger.Fatal(e.Start(port))
}
