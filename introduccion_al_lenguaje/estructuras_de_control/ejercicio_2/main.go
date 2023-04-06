package main

/*Ejercicio 2 - Préstamo

Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga
préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un
año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los
que posean un sueldo superior a $100.000.
Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de
acuerdo a cada caso.
Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

import (
	"fmt"
)

func main() {

	var edad int
	var empleo string
	var antiguedad int
	var salario float64

	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)
	if edad < 22 {
		fmt.Println("No puede acceder al prestamo porque tienes menos de 22 años")
	} else {
		fmt.Println("Si tiene empleo, ingrese 'si', de lo contrario ingrese 'no':")
		fmt.Scanln(&empleo)
		if empleo == "no" {
			fmt.Println("No puede acceder al prestamo porque no tiene empleo")
		} else {
			fmt.Println("Ingrese su antiguedad en años:")
			fmt.Scanln(&antiguedad)
			if antiguedad < 1 {
				fmt.Println("No puede acceder al prestamo porque tiene menos de 1 año de antiguedad en su empleo")
			} else {
				fmt.Println("Ingrese su salario:")
				fmt.Scanln(&salario)
				if salario < 100000.00 {
					fmt.Println("Accedes al prestamo con intereses")
				} else {
					fmt.Println("Accedes al prestamo sin intereses")
				}
			}

		}
	}
}
