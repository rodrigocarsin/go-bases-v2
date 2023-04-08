package main

type Alumno struct {
	Nombre       string
	Apellido     string
	DNI          int
	FechaIngreso string
}

func (a *Alumno) detalles() {
	println("Nombre: ", a.Nombre)
	println("Apellido: ", a.Apellido)
	println("DNI: ", a.DNI)
	println("Fecha: ", a.FechaIngreso)
}

func main() {

	a := Alumno{
		Nombre:       "Juan",
		Apellido:     "Perez",
		DNI:          12345678,
		FechaIngreso: "01/01/2020",
	}

	a.detalles()
}
