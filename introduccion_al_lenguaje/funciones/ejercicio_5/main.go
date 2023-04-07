package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func calculateDogFood(numberOfAnimals int) float64 {
	return float64(numberOfAnimals) * 10
}

func calculateCatFood(numberOfAnimals int) float64 {
	return float64(numberOfAnimals) * 5
}

func calculateHamsterFood(numberOfAnimals int) float64 {
	return float64(numberOfAnimals) * 0.25
}

func calculateTarantulaFood(numberOfAnimals int) float64 {
	return float64(numberOfAnimals) * 0.15
}

func calculateFood(animal string, animals int) (float64, string) {

	switch animal {
	case dog:
		return calculateDogFood(animals), ""

	case cat:
		return calculateCatFood(animals), ""

	case hamster:
		return calculateHamsterFood(animals), ""

	case tarantula:
		return calculateTarantulaFood(animals), ""

	default:
		return 0, "Animal incorrecto"
	}

}

func main() {
	fmt.Println(calculateFood(dog, 10))
	fmt.Println(calculateFood(cat, 10))
	fmt.Println(calculateFood(hamster, 10))
	fmt.Println(calculateFood(tarantula, 10))
	fmt.Println(calculateFood("duck", 10))

}
