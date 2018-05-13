package data

import (
	"database/sql"
	"errors"
	"hatch/user-service/models"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//PostgresEmailRepository opens up acces to email codes
type PostgresEmailRepository struct {
	db *sqlx.DB
}

//TODO: remove codes past thier time limit

//NewPostgresEmailRepository vends a new repo taking in a db
func NewPostgresEmailRepository(db *sqlx.DB) *PostgresEmailRepository {
	//fetch initial data, start timer for further fetching etc
	return &PostgresEmailRepository{db: db}
}

//GetEmailCode gets an email code if it exists & is in expiration time, else throws an error
func (r *PostgresEmailRepository) GetEmailCode(code uuid.UUID) (models.EmailCode, error) {
	var emailCode models.EmailCode

	err := r.db.Get(emailCode, "SELECT * FROM email_codes WHERE code = ?", code)

	switch {
	case err == sql.ErrNoRows:
		return emailCode, err
	case err != nil:
		log.Fatal(err)
		return emailCode, err
	}

	if time.Now().After(emailCode.Expiration) {
		err = errors.New("Code is Expired")
	}

	return emailCode, err
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
	_, err := r.db.Exec("DELETE FROM email_codes where code = $1", code)
	if err == sql.ErrNoRows {
		return nil
	}

	return err
}
