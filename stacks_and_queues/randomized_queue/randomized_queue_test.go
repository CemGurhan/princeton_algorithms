package randomizedqueue_test

import (
	"testing"

	rq "github.com/cemgurhan/princetonalgo/stacks_and_queues/randomized_queue"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty_WithEmptyRandomQueue_ReturnsTrue(t *testing.T) {
	testRandomQueue := rq.RandomizedQueue[float64]{}
	returnBool := testRandomQueue.IsEmpty()

	assert.True(t, returnBool)
}

func TestIsEmpty_WithNonEmptyRandomQueue_ReturnsFalse(t *testing.T) {
	testRandomQueue := rq.RandomizedQueue[float64]{2.2, 3.1}
	returnBool := testRandomQueue.IsEmpty()

	assert.False(t, returnBool)
}

func TestSize_WithFiveElementRandomQueue_ReturnsFive(t *testing.T) {
	testRandomQueue := rq.RandomizedQueue[string]{"one", "two", "three", "four", "five"}

	actualSize := testRandomQueue.Size()

	assert.Equal(t, 5, actualSize)
}

func TestSize_WithEmptyRandomQueue_ReturnsZero(t *testing.T) {
	testRandomQueue := rq.RandomizedQueue[string]{}

	actualSize := testRandomQueue.Size()

	assert.Equal(t, 0, actualSize)
}

func TestEnqueue_WithFullCapacityQueue_AddsItemToEndOfQueue(t *testing.T) {
	expectedQueue := rq.RandomizedQueue[string]{"Hi", "There", "Friend"}

	actualQueue := rq.RandomizedQueue[string]{"Hi", "There"}
	actualQueue.Enqueue("Friend")

	assert.Equal(t, expectedQueue, actualQueue)
}

func TestEnqueue_WithNonEmptyQueue_AddsItemToEndOfQueue(t *testing.T) {
	expectedQueue := rq.RandomizedQueue[string]{"Hi", "There", "Friend"}

	actualQueue := make(rq.RandomizedQueue[string], 2, 3)
	actualQueue[0] = "Hi"
	actualQueue[1] = "There"
	actualQueue.Enqueue("Friend")

	assert.Equal(t, expectedQueue, actualQueue)
}
