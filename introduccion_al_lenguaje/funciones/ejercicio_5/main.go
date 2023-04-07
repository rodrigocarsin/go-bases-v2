package main

/*Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan
darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el
animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del
tipo de animal especificado.
*/

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
