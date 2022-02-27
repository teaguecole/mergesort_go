package main

import "fmt"

func main() {
	notSorted := []int{10,3,6,5,9,2,8,1}
	sorted := mergeSort(notSorted)
	fmt.Print(sorted)
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	mid := len(items)/2
	first := mergeSort(items[:mid])
	second := mergeSort(items[mid:])
	return merge(first, second)
}

func merge(first []int, second []int) []int {
	var final []int
	f := 0
	s := 0

	for f < len(first) && s < len(second) {
		if first[f] < second[s] {
			final = append(final, first[f])
			f++
		} else {
			final = append(final, second[s])
			s++
		}
	}
	for ; f < len(first); f++ {
		final = append(final, first[f])
	}
	for ; s < len(second); s++ {
		final = append(final, second[s])
	}
	return final
}

