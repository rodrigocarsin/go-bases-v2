package main

/*Ejercicio 2 - Clima

Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura,
humedad y presión atmosférica de distintos lugares.
Declará 3 variables especificando el tipo de dato, como valor deben tener la
temperatura, humedad y presión de donde te encuentres.
Imprimí los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?
*/

import (
	"fmt"
)

var temperatura float64 = 25.5
var humedad float64 = 50.5
var presion float64 = 1010.5

func main() {

	/*
		fmt.Println("Temperatura:", temperatura)
		fmt.Println("Humedad:", humedad)
		fmt.Println("Presion:", presion)
	*/

	fmt.Printf("Temperatura: %v \n", temperatura)
	fmt.Println("Humedad:", humedad)
	fmt.Print("Presion:", presion)
}
