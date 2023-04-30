package practicequiz_test

import (
	"testing"

	pq "github.com/cemgurhan/princetonalgo/stacks_and_queues/practice_quiz/question_one"
	"github.com/stretchr/testify/assert"
)

func TestEnqueue_WithNonEmptyQueue_SuccessfullyAddsItemToEndOfQueue(t *testing.T) {
	testQueue := pq.Queue[string]{"hello", "there"}
	expectedQueue := pq.Queue[string]{"hello", "there", "friend"}
	actualQueue := testQueue.Enqueue("friend")

	assert.Equal(t, expectedQueue, *actualQueue)
}

func TestDequeue_WithNonEmptyQueue_SuccessfullyRemovesAndReturnsItemFromQueue(t *testing.T) {
	expectedItem := "hello"
	expectedQueue := pq.Queue[string]{"there", "friend"}
	actualQueue := pq.Queue[string]{"hello", "there", "friend"}

	actualItem := actualQueue.Dequeue()

	assert.Equal(t, expectedItem, actualItem)
	assert.Equal(t, expectedQueue, actualQueue)
}

func TestDequeue_WhenQueueEmpty_ReturnsEmptyString(t *testing.T) {
	emptyTestQueue := pq.Queue[string]{}

	actualItem := emptyTestQueue.Dequeue()

	assert.Equal(t, "", actualItem)
}
