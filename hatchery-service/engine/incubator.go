package engine

import (
	"fmt"
	"hatchery/data"
	"hatchery/models"
	"time"
)

//Incubator is the engine that drives egg hatching, its proprty is the ticker that drives it
type Incubator struct {
	Ticker *time.Ticker
}

//NewIncubator creates and Incubaotr object and returns it
func NewIncubator() Incubator {
	ticker := time.NewTicker(time.Second)
	defer func() {
		go tick(*ticker)
	}()
	incubator := Incubator{Ticker: ticker}
	return incubator
}

func tick(ticker time.Ticker) {
	for range ticker.C {
		go Incubate()
	}
}

//Incubate serves to incubate eggs and put them into the work stream when they are due to hatch
func Incubate() {
	//fmt.Println("incubating ")
	now := time.Now().UTC()
	removed := 0

	defer func() {
		data.RemoveEggs(removed)
	}()

	for _, egg := range data.RetrieveEggs() {
		if egg.HatchTime.Before(now) {
			go NotifyUser(egg)
			removed++
		} else {
			return
		}
	}

}

//NotifyUser will alert eh user of their egg hatching and is called then it is time.
func NotifyUser(egg models.Egg) {
	fmt.Println("egg hatching ", egg)
}
