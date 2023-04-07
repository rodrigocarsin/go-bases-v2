package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculaSalrioMarinera(t *testing.T) {

	//Arrange
	var minutes int = 600
	var categoryA string = "A"
	var categoryB string = "B"
	var categoryC string = "C"
	var expectedResultA float64 = 45000
	var expectedResultB float64 = 18000
	var expectedResultC float64 = 10000

	//Act
	obtainedResultA := calculateSalary(minutes, categoryA)
	obtainedResultB := calculateSalary(minutes, categoryB)
	obtainedResultC := calculateSalary(minutes, categoryC)

	//Assert
	assert.Equal(t, expectedResultA, obtainedResultA)
	assert.Equal(t, expectedResultB, obtainedResultB)
	assert.Equal(t, expectedResultC, obtainedResultC)
}
