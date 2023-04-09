package main

import "fmt"

type salaryError struct {
	Msg string
}

func (e *salaryError) Error() string {
	return fmt.Sprintf("Error: %s", e.Msg)
}

func main() {
	salary := 200000 // asignar un valor de ejemplo

	if salary < 150000 {
		err := &salaryError{"el salario ingresado no alcanza el mínimo imponible"}
		panic(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
