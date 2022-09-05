package main

import (
	Handler "myapp/internal/handlers"

	"github.com/labstack/echo"
)

func main() {
	router := echo.New()
	router.GET("/person", Handler.GetPersons)
	router.GET("/person/:id", Handler.GetById)
	router.POST("/person", Handler.PostPerson)
	router.DELETE("/person/:id", Handler.DeleteById)
	router.PUT("/person/:id", Handler.UpdatePersonById)
	router.Logger.Fatal(router.Start(":8080"))
}
