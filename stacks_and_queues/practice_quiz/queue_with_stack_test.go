package practicequiz_test

import (
	"testing"

	pq "github.com/cemgurhan/princetonalgo/stacks_and_queues/practice_quiz"
	sq "github.com/cemgurhan/princetonalgo/stacks_and_queues/stacks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnqueue_WithNonEmptyQueue_SuccessfullyAddsItemToEndOfQueue(t *testing.T) {
	expectedQueue := sq.Stack{"hello", "there", "friend"}
	actualQueue := pq.Enqueue("friend", []string{"hello", "there"})

	assert.Equal(t, expectedQueue, *actualQueue)
}

func TestDequeue_WithNonEmptyQueue_SuccessfullyRemovesAndReturnsItemFromQueue(t *testing.T) {
	expectedItem := "friend"
	testQueue := sq.Stack{"hello", "there", "friend"}

	actualItem, err := pq.Dequeue(testQueue)

	require.NoError(t, err)
	assert.Equal(t, expectedItem, actualItem)
}

func TestIsEmpty_WhenQueueEmpty_ReturnsTrue(t *testing.T) {
	emptyTestQueue := sq.Stack{}

	expectedBool := true
	actualBool := emptyTestQueue.IsEmpty()

	assert.Equal(t, expectedBool, actualBool)
}

func TestDequeue_WhenQueueEmpty_ReturnsEmptyStringAndError(t *testing.T) {
	emptyTestQueue := sq.Stack{}

	actualItem, err := pq.Dequeue(emptyTestQueue)

	assert.ErrorContains(t, err, "queue empty")
	assert.Equal(t, "", actualItem)
}
