package weightedpathcompression

import (
	"errors"
	"fmt"
)

type Tree []int

type SizeArray []int

// attach each node to it's grandparent
func (t *Tree) FindRoot(index int) (int, error) {
	if len(*t) == 0 {
		return 0, errors.New("tree is empty")
	}

	for index != (*t)[index] {
		(*t)[index] = (*t)[(*t)[index]]
		index = (*t)[index]
	}
	return index, nil
}

func (t *Tree) WeightedUnionPathCompression(indexOne int, indexTwo int, sizeArray []int) error {
	if len(sizeArray) == 0 {
		return errors.New("size array empty")
	}

	rootOfTreeOne, err := t.FindRoot(indexOne)
	if err != nil {
		return fmt.Errorf("error during union: %w", err)
	}
	rootOfTreeTwo, err := t.FindRoot(indexTwo)
	if err != nil {
		return fmt.Errorf("error during union: %w", err)
	}

	if rootOfTreeOne == rootOfTreeTwo {
		return errors.New("trees are already connected")
	}

	if sizeArray[rootOfTreeOne] <= sizeArray[rootOfTreeTwo] {
		(*t)[rootOfTreeOne] = rootOfTreeTwo
	} else if sizeArray[rootOfTreeOne] >= sizeArray[rootOfTreeTwo] {
		(*t)[rootOfTreeTwo] = rootOfTreeOne
	}

	return nil
}

func (t *Tree) Find(indexOne int, indexTwo int) (bool, error) {
	rootIndexOne, err := t.FindRoot(indexOne)
	if err != nil {
		return false, fmt.Errorf("error during connection test: %w", err)
	}
	rootIndexTwo, err := t.FindRoot(indexTwo)
	if err != nil {
		return false, fmt.Errorf("error during connection test: %w", err)
	}
	if rootIndexOne == rootIndexTwo {
		return true, nil
	}

	return false, nil
}
