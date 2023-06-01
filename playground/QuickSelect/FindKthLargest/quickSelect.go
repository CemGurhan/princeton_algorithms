package findkthlargest

func FindKthLargest(array *[]int, start int, end int, k int) int {
	if start < end {
		pivotIndex := partition(array, start, end)
		if pivotIndex == k-1 {
			return (*array)[pivotIndex]
		}
		if pivotIndex > k-1 {
			return FindKthLargest(array, start, pivotIndex-1, k)
		}
		if pivotIndex < k-1 {
			return FindKthLargest(array, pivotIndex+1, end, k)
		}
	}
	return (*array)[k-1]
}

func partition(array *[]int, start int, end int) int {
	pivot := (*array)[end]
	positionalIndex := start

	for i := start; i <= end-1; i++ {
		if (*array)[i] >= pivot {
			(*array)[i], (*array)[positionalIndex] = (*array)[positionalIndex], (*array)[i]
			positionalIndex++
		}
	}
	(*array)[end], (*array)[positionalIndex] = (*array)[positionalIndex], (*array)[end]

	return positionalIndex
}
