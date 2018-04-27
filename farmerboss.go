package main

import (
	"fmt"
  "hatchery/models"
  "time"
  "hatchery/data"
)

const MAX_FARMERS = 8192
const MIN_FARMERS = 4
var FarmerPool chan Farmer
var eggChannel chan models.Egg
var ticker *time.Ticker
var softCap int

func StartFarmers(nfarmers int) {
  // First, initialize the channel we are going to but the workers' work channels into.
  FarmerPool = make(chan Farmer, nfarmers)
  eggChannel = make(chan models.Egg, nfarmers)
  ticker = time.NewTicker(time.Second)
  softCap = 8
  // Now, create 4 of our workers.
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

func ScaleFarmers(){
  fmt.Println("There are ", len(FarmerPool), "farmers")
	if len(FarmerPool) < softCap / 4 {
    softCap *= 2 //double softcap
		for i := 0; i < softCap/2; i++ {
			fmt.Println("Starting worker", i+1)
			farmer := NewFarmer(i+1, FarmerPool)
			farmer.Start()
		  }
	} else if len(FarmerPool) >= softCap/4 && len(FarmerPool) <= softCap/2{
    //do nothing, this is resting point
    if len(FarmerPool) == softCap/2 && len(FarmerPool) > MIN_FARMERS{
      softCap /=2
    }
  }else{
    //check if we have too many workers hanging around
    
    farmer := <-FarmerPool
    fmt.Println("Telling a farmer to retire")
    farmer.Retire()
	}
}

func Incubate(c chan models.Egg){
	fmt.Println("incubating ")
	now := time.Now().UTC()
	removed := 0 

  go ScaleFarmers()

	defer func(){
		data.RemoveEggs(removed)
	}()

	for _,egg := range data.RetrieveEggs(){
		if egg.HatchTime.Before(now) {
      c <- egg
			removed++
		}else{
			return
		}
	}

}