package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateEstadistics(t *testing.T) {

	//Arrange
	expectedMinFloat := 1.0
	expectedAvgFloat := 5.5
	expectedMaxFloat := 10.0
	expectedMinString := "its the minimum"
	expectedAvgString := "its the average"
	expectedMaxString := "its the maximum"

	//Act
	obtainedMinFloat, obtainedMinString := returnOperation(Minimum, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	obtainedAvgFloat, obtainedAvgString := returnOperation(Average, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	obtainedMaxFloat, obtainedMaxString := returnOperation(Maximum, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	//Assert
	assert.Equal(t, expectedMinFloat, obtainedMinFloat)
	assert.Equal(t, expectedMinString, obtainedMinString)
	assert.Equal(t, expectedAvgFloat, obtainedAvgFloat)
	assert.Equal(t, expectedAvgString, obtainedAvgString)
	assert.Equal(t, expectedMaxFloat, obtainedMaxFloat)
	assert.Equal(t, expectedMaxString, obtainedMaxString)

}
