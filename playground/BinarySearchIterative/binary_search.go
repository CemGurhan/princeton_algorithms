package binarysearchiterative

func BinarySearch(slice []int, target int) bool {
	startIndex := 0
	endIndex := len(slice) - 1

	for startIndex <= endIndex {
		middleIndex := (startIndex + endIndex) / 2

		if slice[middleIndex] == target {
			return true
		}

		if slice[middleIndex] < target {
			startIndex = middleIndex + 1
			continue
		} else if slice[middleIndex] > target {
			endIndex = middleIndex - 1
			continue
		}
	}

	return false
}

// [1,2,3,5,9,55] target = 2
//  ^
