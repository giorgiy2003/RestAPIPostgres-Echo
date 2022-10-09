package Logic

import (
	"errors"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"
	"strconv"
)

func Create(p Model.Person) error {
	if _, err := Repository.Connection.Exec(`INSERT INTO "person" ("person_email", "person_phone", "person_firstName", "person_lastName") VALUES ($1, $2,$3,$4)`, p.Email, p.Phone, p.FirstName, p.LastName); err != nil {
		return err
	}
	return nil
}

func ReadOne(id string) ([]Model.Person, error) {
	person_id, err := strconv.Atoi(id)
	if err != nil {
		return nil,err
	}
	row, err := Repository.Connection.Query(`SELECT * FROM "person" WHERE "person_id" = $1`, person_id)
	if err != nil {
		return nil, err
	}
	var personInfo = []Model.Person{}
	for row.Next() {
		var p Model.Person
		err := row.Scan(&p.Id, &p.Email, &p.Phone, &p.FirstName, &p.LastName)
		if err != nil {
			return nil, err
		}
		personInfo = append(personInfo, p)
	}
	return personInfo, nil
}

func Read() ([]Model.Person, error) {
	row, err := Repository.Connection.Query(`SELECT * FROM "person" ORDER BY "person_id"`)
	if err != nil {
		return nil, errors.New("not found")
	}
	var personInfo = []Model.Person{}
	for row.Next() {
		var p Model.Person
		err := row.Scan(&p.Id, &p.Email, &p.Phone, &p.FirstName, &p.LastName)
		if err != nil {
			return nil, err
		}
		personInfo = append(personInfo, p)
	}
	return personInfo, nil
}

func Update(p Model.Person, id string) error {
	person_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if _, err := Repository.Connection.Exec(`UPDATE "person" SET "person_email" = $1,"person_phone" = $2,"person_firstName" = $3,"person_lastName" = $4  WHERE "person_id" = $5`, p.Email, p.Phone, p.FirstName, p.LastName, person_id); err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	person_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if _, err := Repository.Connection.Exec(`DELETE FROM "person" WHERE "person_id" = $1`, person_id); err != nil {
		return err
	}
	return nil
}


