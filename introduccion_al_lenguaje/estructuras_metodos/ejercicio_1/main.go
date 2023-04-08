package main

import "fmt"

// 1. Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// 2. Tener un slice global de Product llamado Products instanciado con valores.

var Products = []Product{
	{ID: 1, Name: "Macbook Pro", Price: 2000, Description: "Laptop", Category: "Computers"},
	{ID: 2, Name: "iPhone 12", Price: 1000, Description: "Smartphone", Category: "Phones"},
	{ID: 3, Name: "iPad Pro", Price: 1500, Description: "Tablet", Category: "Tablets"},
}

// 3. 2 métodos asociados a la estructura Product: Save(), GetAll().
//El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
//El método GetAll() deberá imprimir todos los productos guardados en el slice Products.

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p *Product) GetAll() {
	fmt.Println(Products)
}

// 4. Una función getById() al cual se le deberá pasar un INT como parámetro y retorna
//el producto correspondiente al parámetro pasado.

func getById(id int) {

	for _, product := range Products {
		if product.ID == id {
			fmt.Println(product)
		}
	}
}

// 5. Ejecutar al menos una vez cada método y función definido desde main().

func main() {
	p1 := Product{ID: 4, Name: "iMac", Price: 3000, Description: "Desktop", Category: "Computers"}
	p1.Save()

	var p Product
	p.GetAll()

	getById(1)
	getById(2)
	getById(3)
	getById(4)
}
