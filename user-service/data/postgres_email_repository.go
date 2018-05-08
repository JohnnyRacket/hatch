package data

import (
	"database/sql"
	"errors"
	"hatch/user-service/models"
	"log"
	"time"

	"github.com/google/uuid"
	// fallback for the sql package
	_ "github.com/lib/pq"
)

//PostgresEmailRepository opens up acces to email codes
type PostgresEmailRepository struct {
	db *sql.DB
}

//TODO: remove codes past thier time limit

//NewPostgresEmailRepository vends a new repo taking in a db
func NewPostgresEmailRepository(db *sql.DB) *PostgresEmailRepository {
	//fetch initial data, start timer for further fetching etc
	return &PostgresEmailRepository{db: db}
}

//GetEmailCode gets an email code if it exists & is in expiration time, else throws an error
func (r *PostgresEmailRepository) GetEmailCode(code uuid.UUID) error {
	var emailCode models.EmailCode

	row := r.db.QueryRow("SELECT * FROM email_codes WHERE code = ?", code)
	err := row.Scan(&emailCode.Code, &emailCode.UserId, &emailCode.Expiration)

	switch {
	case err == sql.ErrNoRows:
		return err
	case err != nil:
		log.Fatal(err)
		return err
	}

	if time.Now().After(emailCode.Expiration) {
		err = errors.New("Code is Expired")
	}

	return err
}

//AddEmailCode adds and email and code pair to repo with a set time limit, else throws an error
func (r *PostgresEmailRepository) AddEmailCode(userId uuid.UUID, code uuid.UUID) error {
	expiration := time.Now().Add(time.Minute * 30)
	_, err := r.db.Exec("INSERT INTO email_codes (code, userId, expiration) VALUES ($1, $2, $3);",
		code,
		userId,
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
