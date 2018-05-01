package data

import "hatchery/models"

var eggs []models.Egg
var index = 0

//TODO: add function to fetch into memory
func InitializeRepository() {
	//fetch initial data, start timer for further fetching etc
}

//StoreEgg will put an egg into either memory or db storage depending on how far out it should hatch
func StoreEgg(egg models.Egg) {
	egg.Id = index
	index++

	for i, item := range eggs {
		if egg.HatchTime.Before(item.HatchTime) {
			eggs = append(eggs[:i], append([]models.Egg{egg}, eggs[i:]...)...)
			return
		}
	}
	eggs = append(eggs, egg)
}

//RetrieveEgg gets an egg by Id
func RetrieveEgg(id int) models.Egg {
	//do nothing atm
	var egg models.Egg
	return egg
}

//RetrieveEggs gets all eggs
func RetrieveEggs() []models.Egg {
	return eggs
}

//RemoveEgg removes an egg by Id
func RemoveEgg(id int) {
	eggs = eggs[1:]
}

//RemoveEggs removes n eggs from memory
func RemoveEggs(number int) {
	eggs = eggs[number:]
}
