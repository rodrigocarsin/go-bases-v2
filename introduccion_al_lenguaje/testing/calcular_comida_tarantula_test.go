package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTarantulaFood(t *testing.T) {

	//Arrange
	var tarantulas int = 10
	expectedTarantulaFood := 1.5
	expectedTarantulaString := ""

	//Act
	obtainedTarantulaFood, obtainedTarantulaString := CalculateFood(Tarantula, tarantulas)

	//Assert
	assert.Equal(t, expectedTarantulaFood, obtainedTarantulaFood)
	assert.Equal(t, expectedTarantulaString, obtainedTarantulaString)

}
