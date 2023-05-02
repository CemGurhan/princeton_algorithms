package quickfind_test

import (
	"testing"

	qf "github.com/cemgurhan/princetonalgo/union_find/quick_find"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFind_WhenTwoItemsConnected_ReturnsTrue(t *testing.T) {
	testQuickFind := qf.QuickFind{0, 1, 1, 0, 4}
	returnedBool := testQuickFind.Find(0, 3)

	assert.True(t, returnedBool)
}

func TestFind_WithTwoItemsNotConnected_ReturnsFalse(t *testing.T) {
	testQuickFind := qf.QuickFind{0, 1, 1, 0, 4}
	returnedBool := testQuickFind.Find(0, 1)

	assert.False(t, returnedBool)
}

func TestFind_WithEmptyQuickFind_ReturnsFalse(t *testing.T) {
	testQuickFind := qf.QuickFind{}
	returnedBool := testQuickFind.Find(0, 1)

	assert.False(t, returnedBool)
}

func TestUnion_WithNoConnectedComponents_ConnectsTwoItems(t *testing.T) {
	expectedQuickFind := qf.QuickFind{3, 1, 2, 3, 4}

	actualQuickFind := qf.QuickFind{0, 1, 2, 3, 4}
	actualQuickFind.Union(0, 3)

	assert.Equal(t, expectedQuickFind, actualQuickFind)
}

func TestUnion_WithOneConnectedComponent_ConnectsItemToConnectedComponent(t *testing.T) {
	expectedQuickFind := qf.QuickFind{4, 1, 2, 4, 4}

	actualQuickFind := qf.QuickFind{3, 1, 2, 3, 4}
	err := actualQuickFind.Union(0, 4)

	require.NoError(t, err)
	assert.Equal(t, expectedQuickFind, actualQuickFind)
}

func TestUnion_WithOneConnectedComponentAndNoOthers_DoesNotChangeQuickFindArray(t *testing.T) {
	expectedQuickFind := qf.QuickFind{4, 4, 4, 4, 4}

	actualQuickFind := qf.QuickFind{4, 4, 4, 4, 4}
	err := actualQuickFind.Union(0, 4)

	require.NoError(t, err)
	assert.Equal(t, expectedQuickFind, actualQuickFind)
}

func TestUnion_WithMixedConnectedComponent_ConnectsItemToCorrectConnectedComponent(t *testing.T) {
	expectedQuickFind := qf.QuickFind{4, 4, 4, 2, 3, 4, 2, 2, 3}

	actualQuickFind := qf.QuickFind{4, 4, 1, 2, 3, 1, 2, 2, 3}
	err := actualQuickFind.Union(2, 0)

	require.NoError(t, err)
	assert.Equal(t, expectedQuickFind, actualQuickFind)
}

func TestUnion_WithEmptyQuickFindArray_ReturnsError(t *testing.T) {
	actualQuickFind := qf.QuickFind{}
	err := actualQuickFind.Union(0, 4)

	require.ErrorContains(t, err, "QuickFind array empty")
}
