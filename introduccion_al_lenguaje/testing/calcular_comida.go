package main

const (
	Dog       = "dog"
	Cat       = "cat"
	Hamster   = "hamster"
	Tarantula = "tarantula"
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

func CalculateFood(animal string, animals int) (float64, string) {

	switch animal {
	case Dog:
		return calculateDogFood(animals), ""

	case Cat:
		return calculateCatFood(animals), ""

	case Hamster:
		return calculateHamsterFood(animals), ""

	case Tarantula:
		return calculateTarantulaFood(animals), ""

	default:
		return 0, "Animal incorrecto"
	}

}
