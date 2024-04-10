package main

func BinarySearchList(items []int, target int) bool {
	low := 0
	high := len(items)

	for low < high {
		mid := (low + (high-low)/2)

		value := items[mid]

		if value == target {
			return true
		} else if value > target {
			high = mid
		} else {
			low = mid + 1
		}

	}

	return false
}
