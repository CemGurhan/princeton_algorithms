package findkthlargest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest_WithKIsTwo_ReturnsValidNumberFromArray(t *testing.T) {
	array := []int{3, 6, 1, 10, 8}

	k := 2
	expectedKthLargestElement := 8
	actualKthSmallestElement := FindKthLargest(&array, 0, len(array)-1, k)

	assert.Equal(t, expectedKthLargestElement, actualKthSmallestElement)
}

func TestFindKthLargest_WithAllPossibleK_ReturnsValidNumbersFromArray(t *testing.T) {
	array := []int{3, 6, 1, 10, 8}

	k := 1
	expectedKthLargestElementOne := 10
	actualKthSmallestElementOne := FindKthLargest(&array, 0, len(array)-1, k)

	k = 2
	expectedKthLargestElement := 8
	actualKthSmallestElement := FindKthLargest(&array, 0, len(array)-1, k)

	k = 3
	expectedKthLargestElementThree := 6
	actualKthSmallestElementThree := FindKthLargest(&array, 0, len(array)-1, k)

	k = 4
	expectedKthLargestElementFour := 3
	actualKthSmallestElementFour := FindKthLargest(&array, 0, len(array)-1, k)

	k = 5
	expectedKthLargestElementFive := 1
	actualKthSmallestElementFive := FindKthLargest(&array, 0, len(array)-1, k)

	assert.Equal(t, expectedKthLargestElementOne, actualKthSmallestElementOne)
	assert.Equal(t, expectedKthLargestElement, actualKthSmallestElement)
	assert.Equal(t, expectedKthLargestElementThree, actualKthSmallestElementThree)
	assert.Equal(t, expectedKthLargestElementFour, actualKthSmallestElementFour)
	assert.Equal(t, expectedKthLargestElementFive, actualKthSmallestElementFive)
}
