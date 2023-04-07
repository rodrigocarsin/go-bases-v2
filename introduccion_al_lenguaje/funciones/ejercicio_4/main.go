package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func retunrOperation(operador string, notas ...float64) (float64, string) {

	switch operador {

	case minimum:
		return opOrchestrator(notas, minimumCalc), "its the minimum"

	case average:
		return opOrchestrator(notas, averageCalc), "its the average"

	case maximum:
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

func main() {

	fmt.Println(retunrOperation(minimum, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(retunrOperation(average, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(retunrOperation(maximum, 3, 4, 5, 6, 7, 8, 9, 10))

}
