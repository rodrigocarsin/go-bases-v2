package main

/*Ejercicio 2 - Imprimir datos

A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
Ahora que el archivo sí existe, el panic no debe ser lanzado.
Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga.
En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer”
que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.
*/

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(filename string) string {

	file, err := os.Create(filename)

	if err != nil {
		panic("Error creando el archivo: " + err.Error())
	} else {
		file.WriteString("Nombre: Juan Perez, DNI: 12345678, Edad: 20, Salario: 10000")
	}

	text, err := io.ReadAll(file)
	if err != nil {
		panic("Error leyendo el archivo: " + err.Error())
	}

	defer func() {

		if err := recover(); err != nil {
			file.Close()
			fmt.Println("ejecución finalizada", err)
		}
	}()

	// if err != nil {
	// 	panic("el archivo no se pudo crear")
	// }
	// file.WriteString("Nombre: Juan Perez, DNI: 12345678, Edad: 20, Salario: 10000")

	// text, err1 := os.ReadFile(file.Name())

	// if err1 != nil {
	// 	fmt.Println(string(text))
	// } else {
	// 	panic("el archivo no se pudo leer")
	// }

	return string(text)

}

func main() {
	fmt.Println(WriteFile("customers.txt"))
}
