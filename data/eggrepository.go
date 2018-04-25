package data

import "hatchery/models"

var eggs []models.Egg
var index int = 0
func StoreEgg(egg models.Egg){
	egg.Id = index
	index++

	for i,item := range eggs{
		if egg.HatchTime.Before(item.HatchTime){
			eggs = append(eggs[:i], append([]models.Egg{egg}, eggs[i:]...)...)
			return
		}
	}
	eggs = append(eggs, egg)
}

func RetrieveEgg(id int) models.Egg{
	//do nothing atm
	var egg models.Egg
	return egg
}

func RetrieveEggs() []models.Egg{
	return eggs
}

func RemoveEgg(id int){
	eggs = eggs[1:]
}

func RemoveEggs(number int){
	eggs = eggs[number:]
}