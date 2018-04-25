package main

import(
	"time"
	"fmt"
	"hatchery/models"
)

func Farmer(id int, eggs <-chan models.Egg) {
    for j := range eggs {
		fmt.Println("farmer ", id, "is now worker to harvest ", j, "at time ", time.Now().UTC())
    }
}