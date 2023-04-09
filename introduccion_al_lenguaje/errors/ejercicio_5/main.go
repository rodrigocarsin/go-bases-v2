package main

/*Ejercicio 5 -  Impuestos de salario #5
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrollá las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá
descontar el 10 % en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un
número negativo, la función debe devolver un error. El mismo tendrá que
indicar “Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
*/

import (
	"errors"
	"fmt"
)

func calcularSalario(horas int, valor float64) (float64, error) {
	var salario float64
	var err error

	if horas < 80 || horas < 0 {
		err := errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return 0, err
	} else {
		salario = float64(horas) * valor
		if salario >= 150000 {
			salario = salario * 0.9
		}
	}

	return salario, err
}

func main() {
	salario, err := calcularSalario(60, 1600)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salario)
	}
}
