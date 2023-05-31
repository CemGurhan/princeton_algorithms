package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort_WithUnsortedArray_SortsArray(t *testing.T) {
	actualArray := []int{9, 2, 1, 3, 10, 7, 0}
	expectedArray := []int{0, 1, 2, 3, 7, 9, 10}

	QuickSort(&actualArray, 0, len(actualArray)-1)

	assert.Equal(t, expectedArray, actualArray)
}
