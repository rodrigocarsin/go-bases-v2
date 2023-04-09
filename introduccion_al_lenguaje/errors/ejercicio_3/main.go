package main

/*Ejercicio 3 - Impuestos de salario #3

Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que,
en reemplazo de “Error()”,  se implemente “errors.New()”.
*/

import (
	"errors"
	"fmt"
)

func implementarError(salary int) error {

	if salary <= 10000 {
		var errAux string
		errAux = ErrSalary.Error()
		return errors.New("Salario bajo. " + errAux)
	} else {
		return nil
	}
}

var ErrSalary = errors.New("Error: el salario es menor a 10.000")

func main() {
	salary := 2000 // asignar un valor de ejemplo

	err := implementarError(salary)

	fmt.Println(err)

	//No toma el errors.Is() - VER CLASE PARA VERIFICAR PORQUE NO FUNCIONA
	if err != nil {
		if errors.Is(err, ErrSalary) {
			fmt.Println("Error: el salario es menor a 10.000")
		} else {
			fmt.Println("No agarra viaje")
		}
	} else {
		fmt.Println("Debe pagar impuesto porque su salario es", salary)
	}

}
