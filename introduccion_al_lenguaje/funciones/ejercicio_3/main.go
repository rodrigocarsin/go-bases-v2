package main

/*Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad
de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
por mes, la categoría y que devuelva su salario.

*/
import "fmt"

func calculateSalary(minutes int, category string) {

	var hours float64
	hours = float64(minutes / 60)

	/*
		switch category {
		case "A":
			return hours*3000 + (hours * 3000 * 0.5)
		case "B":
			return hours*1500 + (hours * 1500 * 0.2)
		case "C":
			return hours * 1000
		} */

	if category == "A" {
		fmt.Println(hours*3000 + (hours * 3000 * 0.5))
	} else {
		if category == "B" {
			fmt.Println(hours*1500 + (hours * 1500 * 0.2))
		} else {
			if category == "C" {
				fmt.Println(hours * 1000)
			}
		}

	}
}

func main() {

	var min int
	var cat string

	fmt.Println("Ingrese la cantidad de minutos trabajados: ")
	fmt.Scanln(&min)
	fmt.Println("Ingrese la categoria (A , B , C): ")
	fmt.Scanln(&cat)

	for {
		if cat != "A" && cat != "B" && cat != "C" {
			fmt.Println("Categoria invalida, reingrese la categoria: ")
			fmt.Scanln(&cat)
		} else {
			calculateSalary(min, cat)
			break
		}

	}

}
