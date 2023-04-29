package stacksandqueues_test

import (
	"testing"

	sq "github.com/cemgurhan/princetonalgo/stacks_and_queues/stacks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPush_WithEmptyStack_SuccessfullyAddsItemToTopOfStack(t *testing.T) {
	testStack := sq.Stack{}

	expectedStack := sq.Stack{"Friend"}
	actualStack := testStack.Push("Friend")

	assert.Equal(t, expectedStack, *actualStack)
}

func TestPop_WithNonEmptyStack_SuccessfullyRemovesAndReturnsItemFromTopOfStack(t *testing.T) {
	testStack := sq.Stack{"Hello", "There", "Friend"}

	expectedItem, err := testStack.Pop()
	actualItem := "Friend"

	require.NoError(t, err)
	assert.Equal(t, expectedItem, actualItem)
}

func TestIsEmpty_WithEmptyStack_ReturnsTrue(t *testing.T) {
	testStack := sq.Stack{}
	expectedBool := true
	actualBool := testStack.IsEmpty()

	assert.Equal(t, expectedBool, actualBool)
}

func TestPop_WithEmptyStack_ReturnsEmptyStringAndError(t *testing.T) {
	testStack := sq.Stack{}

	actualItem, err := testStack.Pop()

	assert.ErrorContains(t, err, "stack empty")
	assert.Equal(t, "", actualItem)
}
