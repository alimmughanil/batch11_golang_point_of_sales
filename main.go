package main

import (
	"batch11_golang_pos/configs"
	"batch11_golang_pos/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	port := os.Getenv("PORT")
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "local" {
		port = "8000"
	}

	e.Start(":" + port)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
