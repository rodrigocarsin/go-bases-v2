package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCatFood(t *testing.T) {

	//Arrange
	var cats int = 10
	expectedCatFood := 50.0
	expectedCatString := ""

	//Act
	obtainedCatFood, obtainedCatString := CalculateFood(Cat, cats)

	//Assert
	assert.Equal(t, expectedCatFood, obtainedCatFood)
	assert.Equal(t, expectedCatString, obtainedCatString)
}
