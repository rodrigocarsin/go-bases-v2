package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculaSalario1(t *testing.T) {

	//Arrange
	var salary float64 = 49999.99

	//Act
	CalculaSalario(salary)

	//Assert
	//No se deberían aplicar impuestos por ser un salario inferior a 50000
	assert.Equal(t, 0.0, CalculaSalario(salary))

}
