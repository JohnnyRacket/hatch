package engine

import (
	"fmt"
	"hatch/hatchery-service/data"
	"hatch/hatchery-service/models"
	"time"
)

//Incubator is the engine that drives egg hatching, its proprty is the ticker that drives it
type Incubator struct {
	Ticker     *time.Ticker
	repository data.EggRepository
}

//NewIncubator creates and Incubaotr object and returns it
func NewIncubator(repo data.EggRepository) Incubator {
	ticker := time.NewTicker(time.Second)
	incubator := Incubator{
		Ticker:     ticker,
		repository: repo,
	}
	defer func() {
		go incubator.tick(*ticker)
	}()
	return incubator
}

func (i Incubator) tick(ticker time.Ticker) {
	for range ticker.C {
		go i.incubate()
	}
}

//Incubate serves to incubate eggs and put them into the work stream when they are due to hatch
func (i Incubator) incubate() {
	//fmt.Println("incubating ")
	now := time.Now().UTC()
	removed := 0

	defer func() {
		i.repository.RemoveEggs(removed)
	}()

	for _, egg := range i.repository.RetrieveEggs() {
		if egg.HatchTime.Before(now) {
			go i.NotifyUser(egg)
			removed++
		} else {
			return
		}
	}

}

//NotifyUser will alert eh user of their egg hatching and is called then it is time.
func (i Incubator) NotifyUser(egg models.Egg) {
	fmt.Println("egg hatching ", egg)
}
