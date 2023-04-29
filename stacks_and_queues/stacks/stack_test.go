package stacksandqueues_test

import (
	"testing"

	sq "github.com/cemgurhan/princetonalgo/stacks_and_queues/stacks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPush_WithEmptyStack_SuccessfullyAddsItemToTopOfStack(t *testing.T) {
	testStack := sq.Stack[string]{}

	expectedStack := sq.Stack[string]{"Friend"}
	actualStack := testStack.Push("Friend")

	assert.Equal(t, expectedStack, *actualStack)
}

func TestPop_WithNonEmptyStack_SuccessfullyRemovesAndReturnsItemFromTopOfStack(t *testing.T) {
	testStack := sq.Stack[string]{"Hello", "There", "Friend"}

	expectedItem, err := testStack.Pop()
	actualItem := "Friend"

	require.NoError(t, err)
	assert.Equal(t, expectedItem, actualItem)
}

func TestIsEmpty_WithEmptyStack_ReturnsTrue(t *testing.T) {
	testStack := sq.Stack[string]{}
	expectedBool := true
	actualBool := testStack.IsEmpty()

	assert.Equal(t, expectedBool, actualBool)
}

func TestPop_WithEmptyStack_ReturnsEmptyStringAndError(t *testing.T) {
	testStack := sq.Stack[string]{}

	actualItem, err := testStack.Pop()

	assert.ErrorContains(t, err, "stack empty")
	assert.Equal(t, "", actualItem)
}

func TestClear_WithNonEmptyStack_ClearsStack(t *testing.T) {
	actualStack := sq.Stack[int]{1, 2, 3, 4}

	expectedStack := sq.Stack[int]{}
	actualStack.Clear()

	assert.Equal(t, expectedStack, actualStack)
}

func TestFindMaxItem_WithNonEmptyStack_ReturnsMaxItem(t *testing.T) {
	testStack := sq.Stack[int]{1, 9, 3, 4}

	expectedReturnStack := sq.Stack[int]{9}
	actualReturnStack := testStack.FindMaxItem()

	assert.Equal(t, expectedReturnStack, *actualReturnStack)
}

func TestFindMaxItem_WithTwoMaximums_ReturnsBothMaxItems(t *testing.T) {
	testStack := sq.Stack[int]{1, 9, 9, 4}

	expectedReturnStack := sq.Stack[int]{9, 9}
	actualReturnStack := testStack.FindMaxItem()

	assert.Equal(t, expectedReturnStack, *actualReturnStack)
}

func TestFindMaxItem_WithOneMaximumAndTwoPotentialMaximums_ReturnsOneMaxItem(t *testing.T) {
	testStack := sq.Stack[int]{1, 9, 9, 40, 0}

	expectedReturnStack := sq.Stack[int]{40}
	actualReturnStack := testStack.FindMaxItem()

	assert.Equal(t, expectedReturnStack, *actualReturnStack)
}

func TestFindMaxItem_WithEmptyStack_ReturnsEmptyStack(t *testing.T) {
	testStack := sq.Stack[int]{}

	expectedReturnStack := sq.Stack[int]{}
	actualReturnStack := testStack.FindMaxItem()

	assert.Equal(t, expectedReturnStack, *actualReturnStack)
}
