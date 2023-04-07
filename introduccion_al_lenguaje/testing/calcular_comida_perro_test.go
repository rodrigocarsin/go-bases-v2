package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDogFood(t *testing.T) {

	//Arrange
	var dogs int = 10
	expectedDogFood := 100.0
	expectedDogString := ""

	//Act
	obtainedDogFood, obtainedDogString := CalculateFood(Dog, dogs)

	//Assert
	assert.Equal(t, expectedDogFood, obtainedDogFood)
	assert.Equal(t, expectedDogString, obtainedDogString)

}
