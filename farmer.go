package main

import(
	"time"
	"fmt"
)

func Farmer(id int, eggs <-chan Egg) {
    for j := range eggs {
		fmt.Println("farmer ", id, "is now worker to harvest ", j, "at time ", time.Now().UTC())
    }
}