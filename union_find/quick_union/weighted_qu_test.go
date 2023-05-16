package quickunion

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindRoot_WhenIndexFarFromRoot_ReturnsRoot(t *testing.T) {
	tree := Tree{0, 0, 2, 5, 4, 4, 6, 9, 8, 1}
	expectedRoot := 0
	actualRoot, err := tree.FindRoot(9)

	require.NoError(t, err)
	assert.Equal(t, expectedRoot, actualRoot)
}

func TestFindRoot_WhenIndexIsRoot_ReturnsIndexAsRoot(t *testing.T) {
	tree := Tree{0, 0, 2, 5, 4, 4, 6, 9, 8, 1}
	expectedRoot := 8
	actualRoot, err := tree.FindRoot(8)
	require.NoError(t, err)

	treeTwo := Tree{2, 3, 1, 3, 4, 4, 1}
	expectedRootTwo := 4
	actualRootTwo, err := treeTwo.FindRoot(5)
	expectedRootThree := 3
	actualRootThree, err := treeTwo.FindRoot(6)

	require.NoError(t, err)
	assert.Equal(t, expectedRoot, actualRoot)
	assert.Equal(t, expectedRootTwo, actualRootTwo)
	assert.Equal(t, expectedRootThree, actualRootThree)
}

func TestFindRoot_WhenTreeIsEmpty_ReturnsError(t *testing.T) {
	tree := Tree{}
	_, err := tree.FindRoot(8)

	assert.ErrorContains(t, err, "tree is empty")
}

func TestFindRoot_WhenIndexIsNotInTree_Panics(t *testing.T) {
	tree := Tree{0, 0, 2, 5, 4, 4, 6, 9, 8, 1}

	assert.Panics(t, func() { tree.FindRoot(10) })
}

func TestWeightedUnion_BetweenTwoTreesOfDifferentSize_ConnectsRootOfSmallerToLarger(t *testing.T) {
	expectedTree := Tree{0, 0, 2, 5, 0, 4, 6, 9, 8, 1}
	actualTree := Tree{0, 0, 2, 5, 4, 4, 6, 9, 8, 1}
	sizeArray := make(SizeArray, len(expectedTree))

	sizeArray[0] = 4
	sizeArray[4] = 3
	actualTree.WeightedUnion(9, 5, sizeArray)

	assert.Equal(t, expectedTree, actualTree)
}

func TestWeightedUnion_BetweenTwoTreesOfSameSize_ConnectsRootOfOneToAnother(t *testing.T) {
	expectedTree := Tree{4, 0, 2, 5, 4, 4, 6, 7, 8, 1}
	actualTree := Tree{0, 0, 2, 5, 4, 4, 6, 7, 8, 1}
	sizeArray := make(SizeArray, len(expectedTree))

	sizeArray[0] = 3
	sizeArray[4] = 3
	actualTree.WeightedUnion(9, 5, sizeArray)

	assert.Equal(t, expectedTree, actualTree)
}

func TestWeightedUnion_WhenTreeEmpty_ReturnsError(t *testing.T) {
	actualTree := Tree{}
	sizeArray := make(SizeArray, 5)

	sizeArray[0] = 3
	sizeArray[4] = 3
	err := actualTree.WeightedUnion(9, 5, sizeArray)

	assert.ErrorContains(t, err, "error during union: tree is empty")
}

func TestWeightedUnion_WhenSizeArrayEmpty_ReturnsError(t *testing.T) {
	actualTree := Tree{3, 4, 1}
	sizeArray := SizeArray{}

	err := actualTree.WeightedUnion(9, 5, sizeArray)

	assert.ErrorContains(t, err, "size array empty")
}

func TestFind_BetweenTwoConnectedPoints_ReturnsTrue(t *testing.T) {
	actualTree := Tree{2, 3, 1, 3, 3}

	isConnectedOne, err := actualTree.Find(1, 2)
	require.NoError(t, err)
	isConnectedTwo, err := actualTree.Find(4, 3)
	require.NoError(t, err)

	assert.True(t, isConnectedOne)
	assert.True(t, isConnectedTwo)
}

func TestFind_BetweenTwoNonConnectedPoints_ReturnsFalse(t *testing.T) {
	actualTree := Tree{2, 3, 1, 3, 4, 4, 1}

	isConnectedOne, err := actualTree.Find(1, 4)
	require.NoError(t, err)
	isConnectedTwo, err := actualTree.Find(5, 6)
	require.NoError(t, err)

	assert.False(t, isConnectedOne)
	assert.False(t, isConnectedTwo)
}

func TestFind_EmptyTree_ReturnsError(t *testing.T) {
	actualTree := Tree{}

	_, err := actualTree.Find(1, 2)

	assert.ErrorContains(t, err, "error during connection test: tree is empty")
}

func TestFind_IndexNotInTree_Panics(t *testing.T) {
	actualTree := Tree{0, 2, 1}

	assert.Panics(t, func() { actualTree.Find(5, 7) })
}
