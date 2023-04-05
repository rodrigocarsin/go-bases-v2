package main

/*Ejercicio 1 - Imprimí tu nombre
Tendrás que:
Crear una aplicación donde tengas como variable tu nombre y dirección.
Imprimir en consola el valor de cada una de las variables.
*/

import (
	"fmt"
)

var nombre string = "Rodrigo"
var direccion string = "San Lucas"

func main() {
	fmt.Printf("Hola %s, tu dirección es %s", nombre, direccion)
}
