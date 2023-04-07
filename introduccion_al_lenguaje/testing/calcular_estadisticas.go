package main

/*Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de
calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo,
máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere
realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso
que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y
devuelva el cálculo que se indicó en la función anterior.
*/

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func returnOperation(operador string, notas ...float64) (float64, string) {

	switch operador {

	case Minimum:
		return opOrchestrator(notas, minimumCalc), "its the minimum"

	case Average:
		return opOrchestrator(notas, averageCalc), "its the average"

	case Maximum:
		return opOrchestrator(notas, maximumCalc), "its the maximum"
	}
	return 0, ""

}

func opOrchestrator(notas []float64, operation func(notas ...float64) float64) float64 {

	var result float64
	var slice []float64

	for _, nota := range notas {
		slice = append(slice, nota)
	}

	result = operation(slice...)

	return result
}

func minimumCalc(notas ...float64) float64 {
	var result float64
	result = notas[0]
	for _, nota := range notas {
		if result > nota {
			result = nota
		}
	}
	return result
}

func averageCalc(notas ...float64) float64 {
	var result float64
	var count int
	for _, nota := range notas {
		result += nota
		count++
	}
	result = result / float64(count)
	return result
}

func maximumCalc(notas ...float64) float64 {
	var result float64
	result = notas[0]
	for _, nota := range notas {
		if result < nota {
			result = nota
		}
	}
	return result
}
