package main

/*Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar
el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana
más de $150.000 se le descontará además un 10 % (27% en total).
*/

import (
	"fmt"
)

var salario float64

func CalculaSalario(salario float64) float64 {

	if salario < 50000 {
		return 0
	} else if salario < 150000 {
		return salario * 0.17
	} else {
		return salario * 0.27
	}

}

func main() {

	fmt.Println("Ingrese el salario: ")
	fmt.Scan(&salario)

	fmt.Println(CalculaSalario(salario))

	/*
		if salario < 50000 {
			fmt.Println("A usted no se le aplican impuestos")
		} else if salario < 150000 {

			fmt.Println("A usted se le aplica un 17'%' de impuestos, por un total de: ", salario*0.17)
		} else {

			fmt.Println("A usted se le aplica un 27'%' de impuestos, por un total de: ", salario*0.27)
		}
	*/
}
