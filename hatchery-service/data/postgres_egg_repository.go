package data

import (
	"database/sql"
	"fmt"
	"hatch/hatchery-service/models"
	"io/ioutil"
	"log"
	"os"
	"time"

	// fallback for the sql package
	_ "github.com/lib/pq"
)

//PostgresEggRepository is the repo that adapts the app to the db
type PostgresEggRepository struct {
	eggs                   []models.Egg
	index                  int
	db                     *sql.DB
	pw, dbname, host, user string
}

//TODO: add function to fetch into memory
//NewPostgresRepository returnd a new postgres repository
func NewPostgresRepository() PostgresEggRepository {
	//fetch initial data, start timer for further fetching etc
	r := new(PostgresEggRepository)
	dat, err := ioutil.ReadFile(os.Getenv("PGPW_LOCATION"))
	if err != nil {
		fmt.Println(err)
	}

	r.pw = string(dat)
	r.dbname = os.Getenv("PGDBNAME")
	r.user = os.Getenv("PGUSER")
	r.host = os.Getenv("PGHOST")

	r.connect()

	stmt, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS eggs (
		id SERIAL PRIMARY KEY,
		author character varying(255) NOT NULL,
		target character varying(255) NOT NULL,
		message character varying(255),
		picture text,
		layed timestamp with time zone NOT NULL,
		hatchtime timestamp with time zone NOT NULL
		)`)

	if err != nil {
		log.Fatal(err)
		fmt.Print("I broke")
	}

	_, execErr := stmt.Exec()

	if execErr != nil {
		log.Fatal(execErr)
		fmt.Print("I broke")
	}

	fmt.Print("We made the table!")

	r.db.Close() //should be deferred and up by the open

	return *r
}

func (r *PostgresEggRepository) connect() {
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

//StoreEgg will put an egg into either memory or db storage depending on how far out it should hatch
func (r *PostgresEggRepository) StoreEgg(egg models.Egg) {

	egg.Id = r.index
	r.index++
	// db insert
	if egg.HatchTime.After(time.Now().Add(time.Minute * 15)) {
		r.connect()
		r.db.Exec("INSERT INTO eggs (author, target, message, picture, layed, hatchtime) VALUES ($1, $2, $3, $4, $5, $6);",
			egg.Author,
			egg.Target,
			egg.Message,
			egg.Picture,
			egg.Layed,
			egg.HatchTime)

		r.db.Close()
	} else {
		// in memory insert
		for i, item := range r.eggs {
			if egg.HatchTime.Before(item.HatchTime) {
				r.eggs = append(r.eggs[:i], append([]models.Egg{egg}, r.eggs[i:]...)...)
				return
			}
		}
		r.eggs = append(r.eggs, egg)
	}
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
