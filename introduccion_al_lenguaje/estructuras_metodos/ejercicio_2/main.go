package main

import "fmt"

// 1. Crear una estructura Person con los campos ID, Name, DateOfBirth.

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

// 2. Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.

type Employee struct {
	ID       int
	Position string
	Datos    Person
}

// 3. Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es
// realizar la impresión de los campos de un empleado.

func (e *Employee) PrintEmployee() {
	fmt.Printf("ID: %d, Name: %s, DateOfBirth: %s, ID: %d, Position: %s", e.Datos.ID, e.Datos.Name,
		e.Datos.DateOfBirth, e.ID, e.Position)
}

// 4. Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos
// y por último ejecutar el método PrintEmployee().

func main() {

	person1 := Person{ID: 48781450, Name: "Rodrigo", DateOfBirth: "05/03/1992"}
	employee1 := Employee{ID: 1, Position: "Developer", Datos: person1}

	employee1.PrintEmployee()
}
