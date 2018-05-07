package data

import (
	"database/sql"
	"hatch/hatchery-service/models"
	"time"

	// fallback for the sql package
	_ "github.com/lib/pq"
)

//PostgresEggRepository is the repo that adapts the app to the db
type PostgresEggRepository struct {
	eggs  []models.Egg
	index int
	db    *sql.DB
}

//TODO: add function to fetch into memory

//NewPostgresRepository returnd a new postgres repository
func NewPostgresRepository(db *sql.DB) *PostgresEggRepository {
	//fetch initial data, start timer for further fetching etc
	return &PostgresEggRepository{db: db}
}

//StoreEgg will put an egg into either memory or db storage depending on how far out it should hatch
func (r *PostgresEggRepository) StoreEgg(egg models.Egg) error {

	egg.Id = r.index
	r.index++
	// db insert
	if egg.HatchTime.After(time.Now().Add(time.Minute * 15)) {
		_, err := r.db.Exec("INSERT INTO eggs (author, target, message, picture, layed, hatchtime) VALUES ($1, $2, $3, $4, $5, $6);",
			egg.Author,
			egg.Target,
			egg.Message,
			egg.Picture,
			egg.Layed,
			egg.HatchTime)

		if err != nil {
			return err
		}
	} else {
		// in memory insert
		for i, item := range r.eggs {
			if egg.HatchTime.Before(item.HatchTime) {
				r.eggs = append(r.eggs[:i], append([]models.Egg{egg}, r.eggs[i:]...)...)
				return nil
			}
		}
		r.eggs = append(r.eggs, egg)
	}

	return nil
}

//RetrieveEgg gets an egg by Id
func (r *PostgresEggRepository) RetrieveEgg(id int) models.Egg {
	//do nothing atm
	var egg models.Egg
	return egg
}

//RetrieveEggs gets all eggs
func (r *PostgresEggRepository) RetrieveEggs() []models.Egg {
	if r.eggs == nil {
		return []models.Egg{}
	}
	return r.eggs
}

//RemoveEgg removes an egg by Id
func (r *PostgresEggRepository) RemoveEgg(id int) {
	r.eggs = r.eggs[1:]
}

//RemoveEggs removes n eggs from memory
func (r *PostgresEggRepository) RemoveEggs(number int) {
	r.eggs = r.eggs[number:]
}
