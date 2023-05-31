package quickselect

func QuickSelect(array *[]int, start int, end int) {
	if start < end {
		pivotIndex := partition(array, start, end)
		QuickSelect(array, start, pivotIndex-1)
		QuickSelect(array, pivotIndex+1, end)
	}
}

func partition(array *[]int, start int, end int) int {
	pivot := (*array)[end]
	positionalIndex := start

	for i := start; i <= end-1; i++ {
		if (*array)[i] < pivot {
			(*array)[i], (*array)[positionalIndex] = (*array)[positionalIndex], (*array)[i]
			positionalIndex++
		}
	}
	(*array)[end], (*array)[positionalIndex] = (*array)[positionalIndex], (*array)[end]

	return positionalIndex
}
