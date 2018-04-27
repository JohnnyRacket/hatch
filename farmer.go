package main

import(
	"fmt"
	"hatchery/models"
	"time"
)

type Farmer struct {
	Id          	int
	Work        	chan models.Egg
	FarmerPool 	chan Farmer
	Quit    		chan bool
  }


func NewFarmer(id int, farmerPool chan Farmer) Farmer {
	// Create, and return the worker.
	farmer := Farmer{
	  Id:          	id,
	  Work:        	make(chan models.Egg),
	  FarmerPool: 	farmerPool,
	  Quit:    		make(chan bool)}
	return farmer
}
  

  
  // This function "starts" the worker by starting a goroutine, that is
  // an infinite "for-select" loop.
  func (w *Farmer) Start() {
	  go func() {
		for {
		  // Add ourselves into the worker queue and the stop queue.
		  w.FarmerPool <- *w

		  select {
			case work := <-w.Work:
				// Receive a work request.
				fmt.Printf("worker%d: Received work request %d,\n", w.Id, work.Id)

				time.Sleep(5 * time.Second)
				fmt.Printf("worker%d: Done Counting!\n", w.Id)
				
			case <-w.Quit:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.Id)
				return
			}
		}
	  }()
  }
  
  // Stop tells the worker to stop listening for work requests.
  //
  // Note that the worker will only stop *after* it has finished its work.
  func (w *Farmer) Retire() {
	go func() {
	  w.Quit <- true
	}()
  }