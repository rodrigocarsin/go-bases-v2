package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculaSalario3(t *testing.T) {

	//Arrange
	var salary float64 = 1000000

	//Act
	CalculaSalario(salary)

	//Assert
	//Se esperan impuestos de 27% por ser un salario superior a 100000
	assert.Equal(t, 270000.00, CalculaSalario(salary))

}
