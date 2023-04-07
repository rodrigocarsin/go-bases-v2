package main

/*Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar
el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana
más de $150.000 se le descontará además un 10 % (27% en total).
*/

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
