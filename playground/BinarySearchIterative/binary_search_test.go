package binarysearchiterative

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch_GivenValPresentInSlice_ReturnsTrue(t *testing.T) {
	slice := []int{1, 2, 3, 5, 9, 55}

	assert.True(t, BinarySearch(slice, 1))
	assert.True(t, BinarySearch(slice, 2))
	assert.True(t, BinarySearch(slice, 3))
	assert.True(t, BinarySearch(slice, 5))
	assert.True(t, BinarySearch(slice, 9))
	assert.True(t, BinarySearch(slice, 55))
}

func TestBinarySearch_GivenValNotPresentInSlice_ReturnsFalse(t *testing.T) {
	slice := []int{1, 2, 3, 5, 9, 55}

	assert.False(t, BinarySearch(slice, -1))
}
