package main

/*Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita
generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.
*/

import (
	"errors"
)

func CalcularPromedio(notas []float64) (float64, error) {

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
