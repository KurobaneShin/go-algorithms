package main

func BubbleSort(list []int) []int {

	for i := 0; i < len(list); i++ {

		for j := 0; j < len(list)-1-i; j++ {

			if list[j] > list[j+1] {
				tmp := list[j]
				list[j] = list[j+1]
				list[j+1] = tmp

			}

		}
	}

	return list

}
