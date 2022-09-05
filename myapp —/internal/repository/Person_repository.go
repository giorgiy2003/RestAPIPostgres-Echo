package Repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "4650"
	dbname   = "test"
	sslmode  = "disable"
)

var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

var Connection *sql.DB

func OpenTable() error {
	var err error
	Connection, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	table, err := Connection.Query(`CREATE TABLE IF NOT EXISTS person
	(
		"person_id" integer NOT NULL GENERATED ALWAYS AS IDENTITY,
		"person_email" character varying(32)  NOT NULL,
		"person_phone" character varying(32)  NOT NULL,
		"person_firstName" character varying(32)  NOT NULL,
		"person_lastName" character varying(32)  NOT NULL,
		PRIMARY KEY (person_id)
	)`)
	if err != nil {
		return err
	}
	defer table.Close()
	return nil
}
