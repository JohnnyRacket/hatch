package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

//PostgresEmailRepository opens up acces to email codes
type PostgresEmailRepository struct {
	db                     *sql.DB
	pw, dbname, host, user string
}

//TODO: remove codes past thier time limit

func NewPostgresEmailRepository() PostgresEmailRepository {
	//fetch initial data, start timer for further fetching etc
	r := new(PostgresEmailRepository)
	dat, err := ioutil.ReadFile(os.Getenv("PGPW_LOCATION"))
	if err != nil {
		fmt.Println(err)
	}

	r.pw = string(dat)
	r.dbname = os.Getenv("PGDBNAME")
	r.user = os.Getenv("PGUSER")
	r.host = os.Getenv("PGHOST")

	r.connect()
	defer r.db.Close()

	stmt, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS email_codes (
		code UUID PRIMARY KEY,
		email character varying(255) NOT NULL,
		expiration timestamp with time zone NOT NULL
		)`)

	if err != nil {
		log.Fatal(err)
	}

	_, execErr := stmt.Exec()

	if execErr != nil {
		log.Fatal(execErr)
		fmt.Print("I broke")
	}

	return *r
}

func (r *PostgresEmailRepository) connect() {
	connStr := "user=" + r.user + " password=" + r.pw + " dbname=" + r.dbname + " host=" + r.host + " sslmode=disable"
	var err error
	for i := 0; i < 10; i++ {
		r.db, err = sql.Open("postgres", connStr)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	//db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

//GetEmailCode gets an email code if it exists & is in expiration time, else throws an error
func (r *PostgresEmailRepository) GetEmailCode(code uuid.UUID) error {
	var uuid uuid.UUID
	var email string
	var expiration time.Time

	row := r.db.QueryRow("SELECT * FROM email_codes WHERE code=?", code)
	err := row.Scan(&uuid, &email, &expiration)

	switch {
	case err == sql.ErrNoRows:
		return err
	case err != nil:
		log.Fatal(err)
		return err
	}

	if time.Now().After(expiration) {
		err = errors.New("Code is Expired")
	}

	return err
}

//AddEmailCode adds and email and code pair to repo with a set time limit, else throws an error
func (r *PostgresEmailRepository) AddEmailCode(email string, code uuid.UUID) error {
	r.connect()
	defer r.db.Close()
	expiration := time.Now().Add(time.Minute * 30)
	_, err := r.db.Exec("INSERT INTO email_codes (code, email, expiration) VALUES ($1, $2, $3);",
		email,
		code,
		expiration)

	return err
}

//RemoveEmailCode removes an email code entry, else throws an err
func (r *PostgresEmailRepository) RemoveEmailCode(code uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM email_codes where code = ?", code)
	if err == sql.ErrNoRows {
		return nil
	}

	return err
}
