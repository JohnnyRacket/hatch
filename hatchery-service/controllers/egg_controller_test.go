package controllers_test

import (
	"hatch/hatchery-service/controllers"
	"hatch/hatchery-service/models"
	"testing"
	"time"
)

func TestJitterEggTime(t *testing.T) {

	lowerLimit := time.Date(2012, time.February, 01, 4, 26, 30, 0, time.UTC)
	upperLimit := time.Date(2012, time.February, 01, 4, 27, 30, 0, time.UTC)

	for i := 0; i < 100; i++ {
		testTime := time.Date(2012, time.February, 01, 4, 27, 23, 25, time.UTC)
		egg := models.Egg{HatchTime: testTime}
		controllers.JitterEggTime(&egg)
		if egg.HatchTime.Before(lowerLimit) || egg.HatchTime.After(upperLimit) {
			t.Errorf("Jitter is outside acceptable range of 30 seconds +/-")
		}
	}

}
