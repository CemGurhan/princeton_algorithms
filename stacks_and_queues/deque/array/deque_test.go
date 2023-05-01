package deque_test

import (
	"testing"

	dq "github.com/cemgurhan/princetonalgo/stacks_and_queues/deque/array"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSize_DequeWithFiveItems_ReturnsFive(t *testing.T) {
	testDeque := dq.Deque[int]{11, 2, 0, 9, 32}

	actualSize := testDeque.Size()

	assert.Equal(t, 5, actualSize)
}

func TestSize_DequeWithNoItems_ReturnsZero(t *testing.T) {
	testDeque := dq.Deque[int]{}

	actualSize := testDeque.Size()

	assert.Equal(t, 0, actualSize)
}

func TestAddFirst_ToNonEmptyDeque_AddsItemToFrontOfDeque(t *testing.T) {
	expectedDeque := dq.Deque[string]{"Hi", "There", "Friend"}

	actualDeque := make(dq.Deque[string], 2, 3)
	actualDeque[0] = "There"
	actualDeque[1] = "Friend"
	actualDeque.AddFirst("Hi")

	assert.Equal(t, expectedDeque, actualDeque)
}

func TestAddFirst_ToFullCapacityDeque_AddsItemToFrontOfDeque(t *testing.T) {
	expectedDeque := dq.Deque[string]{"Hi", "There", "Friend"}

	actualDeque := dq.Deque[string]{"There", "Friend"}
	actualDeque.AddFirst("Hi")

	assert.Equal(t, expectedDeque, actualDeque)
}

func TestAddFirst_ToEmptyDeque_AddsItemToFrontOfDeque(t *testing.T) {
	expectedDeque := dq.Deque[string]{"Hi"}

	actualDeque := dq.Deque[string]{}
	actualDeque.AddFirst("Hi")

	assert.Equal(t, expectedDeque, actualDeque)
}

func TestAddLast_ToNonEmptyDeque_AddsItemToBackOfDeque(t *testing.T) {
	expectedDeque := dq.Deque[int]{0, 2, 4, 55}

	actualDeque := make(dq.Deque[int], 3, 4)
	actualDeque[0] = 0
	actualDeque[1] = 2
	actualDeque[2] = 4
	actualDeque.AddLast(55)

	assert.Equal(t, expectedDeque, actualDeque)
}

func TestAddLast_ToFullCapicityDeque_AddsItemToBackOfDeque(t *testing.T) {
	expectedDeque := dq.Deque[int]{0, 2, 4, 55}

	actualDeque := dq.Deque[int]{0, 2, 4}
	actualDeque.AddLast(55)

	assert.Equal(t, expectedDeque, actualDeque)
}

func TestAddLast_ToEmptyDeque_AddsItemToBackOfDeque(t *testing.T) {
	expectedDeque := dq.Deque[int]{55}

	actualDeque := dq.Deque[int]{}
	actualDeque.AddLast(55)

	assert.Equal(t, expectedDeque, actualDeque)
}

func TestRemoveFirst_FromNonEmptyDeque_RemovesAndReturnsFirstItem(t *testing.T) {
	expectedDeque := dq.Deque[string]{"There", "Friend"}
	expectedRemovedItem := "Hi"

	actualDeque := dq.Deque[string]{"Hi", "There", "Friend"}
	actualRemovedItem, err := actualDeque.RemoveFirst()

	require.NoError(t, err)
	assert.Equal(t, expectedRemovedItem, actualRemovedItem)
	assert.Equal(t, expectedDeque, actualDeque)
}

func TestRemoveFirst_FromEmptyIntDeque_ReturnsErrorAndZero(t *testing.T) {
	actualDeque := dq.Deque[int]{}
	actualRemovedItem, err := actualDeque.RemoveFirst()

	require.ErrorContains(t, err, "cannot remove from empty deque")
	assert.Equal(t, 0, actualRemovedItem)
}

func TestRemoveFirst_FromEmptyStringDeque_ReturnsErrorAndEmptyString(t *testing.T) {
	actualDeque := dq.Deque[string]{}
	actualRemovedItem, err := actualDeque.RemoveFirst()

	require.ErrorContains(t, err, "cannot remove from empty deque")
	assert.Equal(t, "", actualRemovedItem)
}

func TestRemoveLast_FromNonEmptyDeque_RemovesAndReturnsLastItemFromDeque(t *testing.T) {
	expectedDeque := dq.Deque[int]{22, 3, 4, 66, 7}
	expectedRemovedItem := -237

	actualDeque := dq.Deque[int]{22, 3, 4, 66, 7, -237}
	actualRemovedItem, err := actualDeque.RemoveLast()

	require.NoError(t, err)
	assert.Equal(t, expectedRemovedItem, actualRemovedItem)
	assert.Equal(t, expectedDeque, actualDeque)
}

func TestRemoveLast_FromEmptyIntDeque_ReturnsErrorAndZero(t *testing.T) {
	actualDeque := dq.Deque[int]{}
	actualRemovedItem, err := actualDeque.RemoveLast()

	require.ErrorContains(t, err, "cannot remove from empty deque")
	assert.Equal(t, 0, actualRemovedItem)
}

func TestRemoveLast_FromEmptyStringDeque_ReturnsEmptyString(t *testing.T) {
	actualDeque := dq.Deque[string]{}
	actualRemovedItem, err := actualDeque.RemoveLast()

	assert.Equal(t, "", actualRemovedItem)
	require.ErrorContains(t, err, "cannot remove from empty deque")
}

func TestIsEmpty_WithEmptyDeque_ReturnsTrue(t *testing.T) {
	actualDeque := dq.Deque[string]{}
	actualBool := actualDeque.IsEmpty()

	assert.Empty(t, actualDeque)
	assert.True(t, actualBool)
}

func TestIsEmpty_WithNonEmptyDeque_ReturnsFalse(t *testing.T) {
	actualDeque := dq.Deque[string]{"Hi"}
	actualBool := actualDeque.IsEmpty()

	assert.False(t, actualBool)
}
