package main


import "time"
import "fmt"
import "hatchery/data"
import "hatchery/models"
//this file serves to sprout the beans when they are due

var eggChannel chan models.Egg
var ticker *time.Ticker

func StartIncubation() chan models.Egg{
	eggChannel = make(chan models.Egg, 100)
	ticker = time.NewTicker(time.Second)
    go func() {
        for range ticker.C {
            Incubate(eggChannel)
        }
	}()
	return eggChannel
}

func StopIncubation(){
	ticker.Stop()
	close(eggChannel)
}

func Incubate(c chan models.Egg){
	fmt.Println("incubating ")
	now := time.Now().UTC()
	removed := 0

	defer func(){
		data.RemoveEggs(removed)
	}()

	for _,egg := range data.RetrieveEggs(){
		if egg.HatchTime.Before(now) {
			removed++
			c <- egg
		}else{
			return
		}
	}

}