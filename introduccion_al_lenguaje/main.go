package main

import (
	"encoding/json"
	"fmt"
)

// fmt.Println("Hello, World!")

// frutas := []string{"manzana", "banana", "pera"}

// //Imprimir el indice y el valor
// for i, fruta := range frutas {
// 	fmt.Println(i, fruta)
// }

// //Imprimir solo el valor
// for _, fruta := range frutas {
// 	fmt.Println(fruta)
// }

// //Implementando break
// sum := 0
// for {
// 	sum++
// 	if sum >= 10 {
// 		break
// 	}
// }
// fmt.Println(sum)

// //Implementando continue
// for i := 0; i < 10; i++ {
// 	if i%2 == 0 {
// 		continue
// 	}
// 	fmt.Println(i, "es impar")
// }

// //Ejemplo arrays

// var a [2]string
// a[0] = "Hello"
// a[1] = "World"
// fmt.Println(a[0], a[1])
// fmt.Println(a)

// //Ejemplo slices

// var s = []bool{true, false}
// fmt.Println(s[0])
// fmt.Println(len(s))

// s = append(s, true, false)
// fmt.Println(s)

// //Ejemplo maps

// var myMap = map[string]int{"Maru": 26, "Rodri": 31}
// fmt.Println(myMap["Maru"])
// fmt.Println(myMap)
// myMap["Bauti"] = 1
// fmt.Println(myMap)
// myMap["Bauti"] = 0
// myMap["Fede"] = 29
// fmt.Println(myMap)
// delete(myMap, "Fede")
// fmt.Println(myMap)

type Persona struct {
	PrimerNombre string `json:"primer_nombre"`
	Apellido     string `json:"apellido"`
	Telefono     string `json:"telefono"`
	Direccion    string `json:"direccion"`
}

func main() {
	p := Persona{"Juan", "Perez", "43434343", "Calle falsa 123"}
	miJSON, err := json.Marshal(p)

	fmt.Println(string(miJSON))
	fmt.Println(err)
}
