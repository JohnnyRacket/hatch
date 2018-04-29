package main

import (
	"fmt"
	"hatchery/data"
	"hatchery/models"
	"time"
)

//MAX_FARMERS is the maximum number of farmers that can exist in the work pool
const MAX_FARMERS = 8192

//MIN_FARMERS is the minimum number of farmers that will be in the work pool
const MIN_FARMERS = 4

//FarmerPool is the pool of workers that will dispatched to handle alerts
var FarmerPool chan Farmer
var eggChannel chan models.Egg
var ticker *time.Ticker
var softCap = 8

//StartFarmers spins up the farmers worker pool and adds the MIN_FARMERS to it.
func StartFarmers(nfarmers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	FarmerPool = make(chan Farmer, nfarmers)
	eggChannel = make(chan models.Egg, nfarmers)
	ticker = time.NewTicker(time.Second)

	// Now, create initial workers.
	for i := 0; i < MIN_FARMERS; i++ {
		fmt.Println("Starting worker", i+1)
		farmer := NewFarmer(i+1, FarmerPool)
		farmer.Start()
	}

	go func() {
		for range ticker.C {
			Incubate(eggChannel)
		}
	}()

	go func() {
		fmt.Println("ready to get work reqs")
		for {
			select {
			case work := <-eggChannel:
				fmt.Println("Received work requeust")
				//do scaling

				go func() {
					farmer := <-FarmerPool
					fmt.Println("Dispatching work request")
					farmer.Work <- work
				}()

			}
		}
	}()
}

//ScaleFarmers serves to dynamically scale farmers up and down. Its an exponential growth, linear shrink.
func ScaleFarmers() {
	fmt.Println("There are ", len(FarmerPool), "farmers")
	if len(FarmerPool) < softCap/4 && softCap < MAX_FARMERS { //max sure the pool has room to expand
		for i := 0; i < softCap; i++ {
			fmt.Println("Starting worker", i+1)
			farmer := NewFarmer(i+1, FarmerPool)
			farmer.Start()
		}
		softCap *= 2 //double softcap after adding workers to represent new number of workers
	} else if len(FarmerPool) >= softCap/4 && len(FarmerPool) <= softCap/2 {
		//do nothing, this is resting point
		//if too many workers are hanging around lower how many we want
		if len(FarmerPool) == softCap/2 && softCap > MIN_FARMERS {
			softCap /= 2
		}
	} else {
		//scale down workers linearly
		farmer := <-FarmerPool
		fmt.Println("Telling a farmer to retire")
		farmer.Retire()
	}
}

//Incubate serves to incubate eggs and put them into the work stream when they are due to hatch
func Incubate(c chan models.Egg) {
	fmt.Println("incubating ")
	now := time.Now().UTC()
	removed := 0

	go ScaleFarmers()

	defer func() {
		data.RemoveEggs(removed)
	}()

	for _, egg := range data.RetrieveEggs() {
		if egg.HatchTime.Before(now) {
			c <- egg
			removed++
		} else {
			return
		}
	}

}
