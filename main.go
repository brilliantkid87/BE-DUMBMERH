package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	mysql.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	e.Static("/uploads", "./uploads")

	var port = os.Getenv("PORT")
	fmt.Println("server running localhost:" + port)
	e.Logger.Fatal(e.Start(":" + port))
}
