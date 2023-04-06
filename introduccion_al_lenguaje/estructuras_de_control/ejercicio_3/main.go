package main

/*Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que reciba  una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
Nota: Validar que el número del mes, sea correcto.
*/

import (
	"fmt"
)

var meses = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
var mes int

func main() {

	sum := 0
	for {
		sum++
		fmt.Println("Ingrese un numero del 1 al 12:")
		fmt.Scanln(&mes)
		if mes < 1 || mes > 12 {
			fmt.Println("El numero ingresado no es valido")
			continue
		} else {
			fmt.Println("El mes seleccionado es: ", meses[mes])
			break
		}

	}
}
