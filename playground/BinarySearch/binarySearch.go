package binarysearch

func BinarySearch(array []int, target int) bool {
	if len(array) == 0 {
		return false
	}
	halfLenIndex := len(array) / 2
	if target == array[halfLenIndex] {
		return true
	} else if len(array) == 1 {
		return false
	}

	if target > array[halfLenIndex] {
		return BinarySearch(array[halfLenIndex+1:], target)
	} else if target < array[halfLenIndex] {
		return BinarySearch(array[:halfLenIndex], target)
	}

	return false
}
