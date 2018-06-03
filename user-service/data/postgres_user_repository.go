package data

import (
	"database/sql"
	"errors"
	"hatch/user-service/models"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	// fallback for the sql package
	_ "github.com/lib/pq"
)

//PostgresUserRepository opens up acces to email codes
type PostgresUserRepository struct {
	db *sqlx.DB
}

//TODO: remove codes past thier time limit

//NewPostgresUserRepository vends a new repo taking in a db
func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	//fetch initial data, start timer for further fetching etc
	return &PostgresUserRepository{db: db}
}

//GetUser gets an email code if it exists & is in expiration time, else throws an error
func (r *PostgresUserRepository) GetUser(id uuid.UUID) (models.User, error) {
	user := new(models.User)

	err := r.db.Get(&user, "SELECT * FROM user WHERE id=$1", id)
	//err := row.Scan(&user.Id, &user.Email, &user.Name, &user.NotificationDetails)

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
	err := r.db.Get(users, "SELECT * FROM user")
	//if there are no users
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return users, nil
}

//CheckUserExists checks if an email is already registered in the db
func (r *PostgresUserRepository) CheckUserExists(email string) (bool, error) {
	user := new(models.User)

	err := r.db.Get(user, "SELECT * FROM user WHERE email = $1", email)

	switch {
	case err == sql.ErrNoRows:
		return false, err
	case err != nil:
		log.Fatal(err)
		return true, err
	}
	return true, nil
}

//AddUser creates the user in the db, else throws an error
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
	_, err := r.db.Exec("DELETE FROM users where id = $1", id)
	if err == sql.ErrNoRows {
		return nil
	}

	return err
}
