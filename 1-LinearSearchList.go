package main

func LinearSearchList[S ~[]E, E comparable](list S, item E) bool {

	for _, val := range list {
		if val == item {
			return true
		}
	}

	return false
}
