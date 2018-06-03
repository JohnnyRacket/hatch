package data

import (
	"hatch/hatchery-service/models"
)

//EggRepository is the interface for access eggs from whatever data layer.
type EggRepository interface {
	StoreEgg(egg models.Egg) error
	RetrieveEgg(id int) models.Egg
	RetrieveEggs() []models.Egg
	RemoveEgg(id int)
	RemoveEggs(number int)
}
