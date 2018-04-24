package main

var eggs []Egg
var index int = 0
func StoreEgg(egg Egg){
	egg.Id = index
	index++

	for i,item := range eggs{
		if egg.HatchTime.Before(item.HatchTime){
			eggs = append(eggs[:i], append([]Egg{egg}, eggs[i:]...)...)
			return
		}
	}
	eggs = append(eggs, egg)
}

func RetrieveEgg(id int) Egg{
	//do nothing atm
	var egg Egg
	return egg
}

func RetrieveEggs() []Egg{
	return eggs
}

func RemoveEgg(id int){
	eggs = eggs[1:]
}

func RemoveEggs(number int){
	eggs = eggs[number:]
}