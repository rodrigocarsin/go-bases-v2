package main

/*Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita
generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.
*/

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas []float64) (float64, error) {

	var average float64
	var count int

	for _, nota := range notas {
		if nota < 0 || nota > 12 {
			return nota, errors.New("Nota invalida ")
		}
		average += nota
		count++
	}

	return average / float64(count), nil
}

func main() {

	fmt.Println("Ingrese una por una las notas del alumno, para finalizar ingrese '0'")

	var nota float64
	var notas []float64

	for {
		fmt.Scanln(&nota)
		if nota == 0 {
			break
		} else {
			notas = append(notas, nota)

		}
	}

	fmt.Println(calcularPromedio(notas))

}
