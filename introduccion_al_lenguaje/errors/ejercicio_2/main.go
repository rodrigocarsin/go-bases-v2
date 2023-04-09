package main

import (
	"errors"
	"fmt"
)

/*En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el
salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000.
La validación debe ser hecha con la función “Is()” dentro del “main”.
*/

// type salaryError2I interface {
// 	Error(salary int) error
// }

type salaryError2 struct {
	//message string
}

var ErrSalary = errors.New("Error: el salario es menor a 10.000")

func (e *salaryError2) Error(salary int) error {

	if salary <= 10000 {
		return fmt.Errorf("Su salario es %d. %w", salary, ErrSalary)
	} else {
		return nil
	}
}

func main() {
	salary := 10000 // asignar un valor de ejemplo

	var ss = salaryError2{}

	var salaryErr = &ss

	err := salaryErr.Error(salary)

	if err != nil {
		if errors.Is(err, ErrSalary) {
			fmt.Println("Error: el salario es menor a 10.000")
			return
		}
	} else {
		fmt.Println("Debe pagar impuesto porque su salario es", salary)
	}

}
