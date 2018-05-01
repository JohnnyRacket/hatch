package data

import (
	"hatchery/models"
)

//EggRepository is the interface for access eggs from whatever data layer.
type EggRepository interface {
	InitializeRepository()
	StoreEgg(egg models.Egg)
	RetrieveEgg(id int) models.Egg
	RetrieveEggs() []models.Egg
	RemoveEgg(id int)
	RemoveEggs(number int)
}
