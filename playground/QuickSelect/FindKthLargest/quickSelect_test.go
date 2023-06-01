package findkthlargest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest_WithKIsTwo_ReturnsValidNumberFromArray(t *testing.T) {
	array := []int{3, 6, 1, 10, 8}

	k := 2
	expectedKthSmallestElement := 8
	actualKthSmallestElement := FindKthLargest(&array, 0, len(array)-1, k)

	assert.Equal(t, expectedKthSmallestElement, actualKthSmallestElement)
}

func TestFindKthLargest_WithAllPossibleK_ReturnsValidNumbersFromArray(t *testing.T) {
	array := []int{3, 6, 1, 10, 8}

	k := 1
	expectedKthSmallestElementOne := 10
	actualKthSmallestElementOne := FindKthLargest(&array, 0, len(array)-1, k)

	k = 2
	expectedKthSmallestElement := 8
	actualKthSmallestElement := FindKthLargest(&array, 0, len(array)-1, k)

	k = 3
	expectedKthSmallestElementThree := 6
	actualKthSmallestElementThree := FindKthLargest(&array, 0, len(array)-1, k)

	k = 4
	expectedKthSmallestElementFour := 3
	actualKthSmallestElementFour := FindKthLargest(&array, 0, len(array)-1, k)

	k = 5
	expectedKthSmallestElementFive := 1
	actualKthSmallestElementFive := FindKthLargest(&array, 0, len(array)-1, k)

	assert.Equal(t, expectedKthSmallestElementOne, actualKthSmallestElementOne)
	assert.Equal(t, expectedKthSmallestElement, actualKthSmallestElement)
	assert.Equal(t, expectedKthSmallestElementThree, actualKthSmallestElementThree)
	assert.Equal(t, expectedKthSmallestElementFour, actualKthSmallestElementFour)
	assert.Equal(t, expectedKthSmallestElementFive, actualKthSmallestElementFive)
}
