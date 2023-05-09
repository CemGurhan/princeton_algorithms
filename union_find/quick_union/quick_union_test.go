package quickunion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRoot_WhenIndexFarFromRoot_ReturnsRoot(t *testing.T) {
	tree := Tree{0, 0, 2, 5, 4, 4, 6, 9, 8, 1}
	expectedRoot := 0
	actualRoot := tree.FindRoot(9)

	assert.Equal(t, expectedRoot, actualRoot)
}

func TestFindRoot_WhenIndexIsRoot_ReturnsIndexAsRoot(t *testing.T) {
	tree := Tree{0, 0, 2, 5, 4, 4, 6, 9, 8, 1}
	expectedRoot := 8
	actualRoot := tree.FindRoot(8)

	assert.Equal(t, expectedRoot, actualRoot)
}

func TestFindRoot_WhenIndexIsNotInTree_Panics(t *testing.T) {
	tree := Tree{0, 0, 2, 5, 4, 4, 6, 9, 8, 1}

	assert.Panics(t, func() { tree.FindRoot(10) })
}
