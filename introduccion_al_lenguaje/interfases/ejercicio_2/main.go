package main

import "fmt"

type Product interface {
	Price() float64
}

type SmallProduct struct {
	ProductPrice float64
}

func (p *SmallProduct) Price() float64 {
	return p.ProductPrice
}

type MediumProduct struct {
	ProductPrice float64
}

func (p *MediumProduct) Price() float64 {
	return p.ProductPrice * 1.03
}

type BigProduct struct {
	ProductPrice float64
}

func (p *BigProduct) Price() float64 {
	return p.ProductPrice*1.06 + 2500.00
}

func ProductFactory(productType string, productPrice float64) Product {

	switch productType {
	case "small":
		return &SmallProduct{ProductPrice: productPrice}
	case "medium":
		return &MediumProduct{ProductPrice: productPrice}
	case "big":
		return &BigProduct{ProductPrice: productPrice}
	default:
		return nil

	}
}

func main() {
	product := ProductFactory("medium", 100.00)
	fmt.Println(product.Price())
}
