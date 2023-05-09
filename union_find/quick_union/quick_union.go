package quickunion

type Tree []int

type SizeArray []int

func (t *Tree) FindRoot(startingIndex int) int {
	pointer := (*t)[startingIndex]
	for pointer != (*t)[pointer] {
		pointer = (*t)[pointer]
	}
	return pointer
}

func (t *Tree) WeightedUnion(indexOne int, indexTwo int, sizeArray []int) {
	rootOfTreeOne := t.FindRoot(indexOne)
	rootOfTreeTwo := t.FindRoot(indexTwo)

	if sizeArray[rootOfTreeOne] <= sizeArray[rootOfTreeTwo] {
		(*t)[rootOfTreeOne] = rootOfTreeTwo
	} else if sizeArray[rootOfTreeOne] >= sizeArray[rootOfTreeTwo] {
		(*t)[rootOfTreeTwo] = rootOfTreeOne
	}
}
