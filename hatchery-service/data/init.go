package data

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//Init creates a db connection and returns its context
func Init() (*sql.DB, error) {
	dat, err := ioutil.ReadFile(os.Getenv("PGPW_LOCATION"))
	if err != nil {
		fmt.Println(err)
		errors.New("Could not read DB Password")
		return nil, err
	}

	pw := string(dat)
	dbname := os.Getenv("PGDBNAME")
	user := os.Getenv("PGUSER")
	host := os.Getenv("PGHOST")

	db, err := connect(user, pw, dbname, host)
	if err != nil {
		return nil, err
	}
	err = prepareEggsTable(db)
	if err != nil {
		return nil, err
	}
	fmt.Println("We made the table!")
	return db, nil
}

func connect(user, pw, dbname, host string) (*sql.DB, error) {
	connStr := "user=" + user + " password=" + pw + " dbname=" + dbname + " host=" + host + " sslmode=disable"
	var err error
	var db *sql.DB
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	//db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func prepareEggsTable(db *sql.DB) error {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS eggs (
		id SERIAL PRIMARY KEY,
		author character varying(255) NOT NULL,
		target character varying(255) NOT NULL,
		message character varying(255),
		picture text,
		layed timestamp with time zone NOT NULL,
		hatchtime timestamp with time zone NOT NULL
		)`)

	_, err := stmt.Exec()

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
