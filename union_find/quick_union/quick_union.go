package quickunion

type Tree []int

//  0  1  2  3  4  5  6  7  8  9
// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

// 0  1  2  3  4  5  6  7  8  9
// 0  0  2  5  4  4  6  9  8  1

func (t *Tree) FindRoot(startingIndex int) int {
	pointer := (*t)[startingIndex]
	for true {
		pointer = (*t)[pointer]
		if pointer == (*t)[pointer] {
			break
		}
	}
	return pointer
}

func (t *Tree) WeightedUnion(indexOne int, indexTwo int) {

}
