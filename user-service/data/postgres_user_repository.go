package data

import (
	"database/sql"
	"errors"
	"hatch/user-service/models"
	"log"

	"github.com/google/uuid"
)

//PostgresUserRepository opens up acces to email codes
type PostgresUserRepository struct {
	db *sql.DB
}

//TODO: remove codes past thier time limit

//NewPostgresUserRepository vends a new repo taking in a db
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	//fetch initial data, start timer for further fetching etc
	return &PostgresUserRepository{db: db}
}

//GetUser gets an email code if it exists & is in expiration time, else throws an error
func (r *PostgresUserRepository) GetUser(id uuid.UUID) (models.User, error) {
	user := new(models.User)

	row := r.db.QueryRow("SELECT * FROM user WHERE id=?", id)
	err := row.Scan(&user.Id, &user.Email, &user.Name, &user.NotificationEndpoint)

	switch {
	case err == sql.ErrNoRows:
		return *user, err
	case err != nil:
		log.Fatal(err)
		return *user, err
	}

	if !user.Validate() {
		err = errors.New("User is Malformed")
	}

	return *user, nil
}

//GetUsers returns all users, else throws an error
func (r *PostgresUserRepository) GetUsers() ([]models.User, error) {

	users := make([]models.User, 0)
	rows, err := r.db.Query("SELECT * FROM user")
	//if there are no users
	if err == sql.ErrNoRows {
		return nil, nil
	}

	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.NotificationEndpoint)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, *user)
	}

	return users, nil
}

//AddUser adds and email and code pair to repo with a set time limit, else throws an error
func (r *PostgresUserRepository) AddUser(email string, name string) (uuid.UUID, error) {
	uuid := uuid.New()
	_, err := r.db.Exec("INSERT INTO users (id , email, name, notificationEndpoint) VALUES ($1, $2, $3, $4);",
		uuid,
		email,
		name,
		"")

	return uuid, err
}

//RemoveEmailCode removes an email code entry, else throws an err
func (r *PostgresUserRepository) RemoveUser(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM users where id = ?", id)
	if err == sql.ErrNoRows {
		return nil
	}

	return err
}
