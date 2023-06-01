package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost:5000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH},
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

	fmt.Println("Running on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}

// func FindDataUsers(c echo.Context) error {
// 	c.Response().Header().Set("Content-type", "application/json")
// 	c.Response().WriteHeader(http.StatusOK)

// 	return json.NewEncoder(c.Response()).Encode(Talent)
// }

// func AddDataUser(c echo.Context) error {
// 	var data DataUser

// 	json.NewDecoder(c.Request().Body).Decode(&data)

// 	Talent = append(Talent, data)
// 	c.Response().Header().Set("Content-type", "application/json")
// 	c.Response().WriteHeader(http.StatusOK)

// 	return json.NewEncoder(c.Response()).Encode(Talent)
// }

// func GetDataUser(c echo.Context) error {
// 	c.Response().Header().Set("Content-type", "application/json")
// 	id := c.Param("id")
// 	var Data DataUser
// 	var cekId = false

// 	for _, talent := range Talent {
// 		if id == talent.Id {
// 			cekId = true
// 			Data = talent
// 		}
// 	}

// 	if !cekId {
// 		c.Response().WriteHeader(http.StatusNotFound)
// 		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
// 	}

// 	c.Response().WriteHeader(http.StatusOK)
// 	return json.NewEncoder(c.Response()).Encode(Data)
// }

// func DeleteDataUser(c echo.Context) error {
// 	id := c.Param("id")
// 	var cekId = false
// 	var index = 0

// 	for i, talent := range Talent {
// 		if id == talent.Id {
// 			cekId = true
// 			index = i
// 		}
// 	}

// 	if !cekId {
// 		c.Response().WriteHeader(http.StatusNotFound)
// 		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
// 	}

// 	Talent = append(Talent[:index], Talent[index+1:]...)

// 	c.Response().Header().Set("Content-type", "application/json")
// 	c.Response().WriteHeader(http.StatusOK)
// 	return json.NewEncoder(c.Response()).Encode(Talent)
// }
