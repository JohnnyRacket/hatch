package data

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
	// fallback for the sql package
	_ "github.com/lib/pq"
)

//Init creates the db context and returns it, also ensures tables are set up
func Init() (*sql.DB, error) {
	dat, err := ioutil.ReadFile(os.Getenv("PGPW_LOCATION"))
	if err != nil {
		fmt.Println(err)
	}

	pw := string(dat)
	dbname := os.Getenv("PGDBNAME")
	user := os.Getenv("PGUSER")
	host := os.Getenv("PGHOST")

	db := connect(user, pw, dbname, host)

	prepareUserDB(db)
	prepareEmailCodeDB(db)

	return db, nil

}

func connect(user, pw, dbname, host string) *sql.DB {
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
	}
	return db
}

func prepareUserDB(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		email character varying(255) NOT NULL,
		name character varying(255) NOT NULL,
		notifcationEndpoint character varying(255) NOT NULL,
		expiration timestamp with time zone NOT NULL
		)`)

	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = stmt.Exec()
	if err == nil {
		fmt.Println("User DB Success")
	}
	return err
}

func prepareEmailCodeDB(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS email_codes (
		code UUID PRIMARY KEY,
		userId UUID REFERENCES users NOT NULL,
		expiration timestamp with time zone NOT NULL
		)`)

	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = stmt.Exec()
	if err == nil {
		fmt.Println("EmailCode DB Success")
	}
	return err
}
