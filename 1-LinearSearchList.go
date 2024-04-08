package main

func LinearSearchList[List ~[]Item, Item comparable](list List, item Item) bool {
	for _, val := range list {
		if val == item {
			return true
		}
	}

	return false
}
