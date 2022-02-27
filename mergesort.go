package main

import "fmt"

func main() {
	notSorted := []int{10,3,6,5,9,2,8,1}
	sorted := mergeSort(notSorted)
	fmt.Print(sorted)
}

/*
	mergeSort is where the divide portion of divide and conquer happens.
	This is breaking down the array in to its smallest pieces.

	items []int: an array of integers that are not sorted.
 */
func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	mid := len(items)/2
	first := mergeSort(items[:mid])
	// When it reaches second for the first time, items on the stack has a length of 2
	second := mergeSort(items[mid:])
	return merge(first, second)
}


// merge is where the actual sorting happens.

func merge(first []int, second []int) []int {
	var final []int
	f := 0
	s := 0

	/*
		When it reaches the for loop for the first time, both first and second
		have one item, and second has a value of 3, and first has a value of 10
		and so itll skip the if clause and append 3 to final.
	 */
	for f < len(first) && s < len(second) {
		if first[f] < second[s] {
			final = append(final, first[f])
			f++
		} else {
			final = append(final, second[s])
			s++
		}
	}
	// f did not get incremented, but s did, so its going to append f to final
	for ; f < len(first); f++ {
		final = append(final, first[f])
	}
	// since s got incremented, 1 = 1, and thus nothing happens here
	for ; s < len(second); s++ {
		final = append(final, second[s])
	}
	// we return [3,10]
	return final
}

