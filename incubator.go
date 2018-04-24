package main


import "time"
import "fmt"
//this file serves to sprout the beans when they are due

var eggChannel chan Egg
var ticker *time.Ticker

func StartIncubation() chan Egg{
	eggChannel = make(chan Egg, 100)
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

func Incubate(c chan Egg){
	fmt.Println("incubating ")
	now := time.Now().UTC()
	removed := 0

	defer func(){
		RemoveEggs(removed)
	}()

	for _,egg := range RetrieveEggs(){
		if egg.HatchTime.Before(now) {
			removed++
			c <- egg
		}else{
			return
		}
	}

}