package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculaSalario2(t *testing.T) {

	//Arrange
	var salary float64 = 50001.00

	//Act
	CalculaSalario(salary)

	//Assert
	//Se deberían aplicar impuestos de 17% por ser un salario superior a 50000
	assert.Equal(t, 8500.17, CalculaSalario(salary))

}
