package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch_GivenArrayAndPresentNumbers_ReturnsTrue(t *testing.T) {
	sortedArray := []int{-1, 2, 3, 5, 10, 77}

	assert.True(t, BinarySearch(sortedArray, -1))
	assert.True(t, BinarySearch(sortedArray, 2))
	assert.True(t, BinarySearch(sortedArray, 3))
	assert.True(t, BinarySearch(sortedArray, 5))
	assert.True(t, BinarySearch(sortedArray, 10))
	assert.True(t, BinarySearch(sortedArray, 77))
}

func TestBinarySearch_GivenArrayNumberNotPresent_ReturnsFalse(t *testing.T) {
	sortedArray := []int{-1, 2, 3, 5, 10, 77}

	assert.False(t, BinarySearch(sortedArray, 100))
}
