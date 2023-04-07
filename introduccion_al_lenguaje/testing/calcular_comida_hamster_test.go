package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHamsterFood(t *testing.T) {

	//Arrange
	var hamsters int = 10
	expectedHamsterFood := 2.5
	expectedHamsterString := ""

	//Act
	obtainedHamsterFood, obtainedHamsterString := CalculateFood(Hamster, hamsters)

	//Assert
	assert.Equal(t, expectedHamsterFood, obtainedHamsterFood)
	assert.Equal(t, expectedHamsterString, obtainedHamsterString)

}
