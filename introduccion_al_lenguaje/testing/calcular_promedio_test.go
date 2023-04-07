package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculaPromedio(t *testing.T) {

	//Arrange
	var notas []float64 = []float64{3, 4, 5, 6, 7, 8, 9, 10}

	expectedResult := 6.5

	var expectedError error
	expectedError = nil

	//Act
	obtainedResult, obtainedError := CalcularPromedio(notas)

	//Assert
	//Esperamos recibir el promedio de las notas, 6.5 y nill
	assert.Equal(t, expectedResult, obtainedResult)
	assert.Equal(t, expectedError, obtainedError)

}
