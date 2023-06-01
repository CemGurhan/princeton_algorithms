package quickselect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSelect_WithKIsTwo_ReturnsValidNumberFromArray(t *testing.T) {
	array := []int{3, 6, 1, 10, 8}

	k := 2
	expectedKthSmallestElement := 3
	actualKthSmallestElement := QuickSelect(&array, 0, len(array)-1, k)

	assert.Equal(t, expectedKthSmallestElement, actualKthSmallestElement)
}

func TestQuickSelect_WithAllPossibleK_ReturnsValidNumbersFromArray(t *testing.T) {
	array := []int{3, 6, 1, 10, 8}

	k := 1
	expectedKthSmallestElementOne := 1
	actualKthSmallestElementOne := QuickSelect(&array, 0, len(array)-1, k)

	k = 2
	expectedKthSmallestElement := 3
	actualKthSmallestElement := QuickSelect(&array, 0, len(array)-1, k)

	k = 3
	expectedKthSmallestElementThree := 6
	actualKthSmallestElementThree := QuickSelect(&array, 0, len(array)-1, k)

	k = 4
	expectedKthSmallestElementFour := 8
	actualKthSmallestElementFour := QuickSelect(&array, 0, len(array)-1, k)

	k = 5
	expectedKthSmallestElementFive := 10
	actualKthSmallestElementFive := QuickSelect(&array, 0, len(array)-1, k)

	assert.Equal(t, expectedKthSmallestElementOne, actualKthSmallestElementOne)
	assert.Equal(t, expectedKthSmallestElement, actualKthSmallestElement)
	assert.Equal(t, expectedKthSmallestElementThree, actualKthSmallestElementThree)
	assert.Equal(t, expectedKthSmallestElementFour, actualKthSmallestElementFour)
	assert.Equal(t, expectedKthSmallestElementFive, actualKthSmallestElementFive)
}
