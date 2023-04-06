package main

/*Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener
cada una de las letras por separado para deletrearla. Para eso tendrán que:
Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad
de letras que contiene la misma.
Luego, imprimí cada una de las letras.
*/

import (
	"fmt"
)

func main() {
	/*
		var palabra string
		fmt.Println("Ingrese una palabra:")
		fmt.Scanln(&palabra)
		fmt.Println("La palabra ingresada tiene", len(palabra), "letras")
		for i := 0; i < len(palabra); i++ {
			fmt.Println(palabra[i])
		}
	*/

	var palabra string
	fmt.Println("Ingrese una palabra:")
	fmt.Scanln(&palabra)
	fmt.Println("La palabra ingresada tiene", len(palabra), "letras")
	for _, letra := range palabra {
		fmt.Println(string(letra))
	}

}
