package Handler

import (
	"fmt"
	"log"
	Logic "myapp/internal/logic"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func PostPerson(c echo.Context) error {
	if err := Repository.OpenTable(); err != nil {
		log.Fatal(err)
	}
	var newPerson Model.Person
	newPerson.Email = c.FormValue("email")
	newPerson.Phone = c.FormValue("phone")
	newPerson.FirstName = c.FormValue("firstName")
	newPerson.LastName = c.FormValue("lastName")
	err := Logic.Create(newPerson)
	if err != nil {
		log.Println(`"Error": "Bad request"`)
		return c.JSON(http.StatusBadRequest, `"Error": "Bad request"`)
	}
	log.Println("Создана запись", newPerson)
	return c.JSON(http.StatusCreated, newPerson)
}

func GetPersons(c echo.Context) error {
	if err := Repository.OpenTable(); err != nil {
		log.Fatal(err)
	}
	persons, err := Logic.Read()
	if len(persons) == 0 {
		log.Println(`"Error": "Записей ещё нет!"`)
		return c.JSON(http.StatusNotFound, "Записей ещё нет!")
	}
	if err != nil {
		log.Println(`"Error": "Bad request"`)
		return c.JSON(http.StatusBadRequest, `"Error": "Bad request"`)
	}
	return c.JSON(http.StatusOK, persons)
}

func GetById(c echo.Context) error {
	if err := Repository.OpenTable(); err != nil {
		log.Fatal(err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(`Неверно введён параметр id`)
		return c.JSON(http.StatusNotFound, `Неверно введён параметр id`)
	}
	persons, err := Logic.ReadOne(id)
	if len(persons) == 0 {
		log.Println("Записи с введённым id не существует!")
		return c.JSON(http.StatusNotFound, "Записи с введённым id не существует!")
	}
	if err != nil {
		log.Println(`Error: Bad request`)
		return c.JSON(http.StatusBadRequest, `"Error": "Bad request"`)
	}
	log.Println(persons)
	return c.JSON(http.StatusOK, persons)
}

func DeleteById(c echo.Context) error {
	if err := Repository.OpenTable(); err != nil {
		log.Fatal(err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(`Неверно введён параметр id`)
		return c.JSON(http.StatusNotFound, `Неверно введён параметр id`)
	}
	persons, err := Logic.ReadOne(id)
	if len(persons) == 0 {
		log.Println("Записи с введённым id не существует!")
		return c.JSON(http.StatusNotFound, "Записи с введённым id не существует!")
	}
	if err != nil {
		log.Println(`Error: Bad request`)
		return c.JSON(http.StatusNotFound, `"Error": "Bad request"`)
	}
	err = Logic.Delete(id)
	if err != nil {
		log.Println(`Error: Bad request`)
		return c.JSON(http.StatusBadRequest, `"Error": "Bad request"`)
	}
	log.Printf("Запись с id = %d  успешно удалена", id)
	return c.JSON(http.StatusOK, fmt.Sprintf("Запись с id = %d  успешно удалена", id))
}

func UpdatePersonById(c echo.Context) error {
	if err := Repository.OpenTable(); err != nil {
		log.Fatal(err)
	}
	var newPerson Model.Person
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(`Неверно введён параметр id`)
		return c.JSON(http.StatusNotFound, `Неверно введён параметр id`)
	}
	persons, err := Logic.ReadOne(id)
	if len(persons) == 0 {
		log.Println("Записи с введённым id не существует!")
		return c.JSON(http.StatusNotFound, "Записи с введённым id не существует!")
	}
	if err != nil {
		log.Println(`Error: Bad request`)
		return c.JSON(http.StatusNotFound, `"Error": "Bad request"`)
	}
	newPerson.Email = c.FormValue("email")
	newPerson.Phone = c.FormValue("phone")
	newPerson.FirstName = c.FormValue("firstName")
	newPerson.LastName = c.FormValue("lastName")
	err = Logic.Update(newPerson, id)
	if err != nil {
		log.Println(`Error: Bad request`)
		return c.JSON(http.StatusBadRequest, `"Error": "Bad request"`)
	}
	log.Printf("Запись с id = %d  успешно обновлена", id)
	return c.JSON(http.StatusOK, fmt.Sprintf("Запись с id = %d  успешно обновлена", id))
}
