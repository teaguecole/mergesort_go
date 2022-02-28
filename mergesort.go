package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var unsorted = rand.Perm(100)
	fmt.Println("Unsorted array:", unsorted)
	sorted := mergeSort(unsorted)
	fmt.Println("Sorted array:", sorted)
}

/*
	mergeSort is where the divide portion of divide and conquer happens.
	This is breaking down the array in to its smallest pieces.

	items []int: an array of integers that are not sorted.
 */
func mergeSort(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}
	mid := len(unsorted)/2
	leftSide := mergeSort(unsorted[:mid])
	// When it reaches rightSide for the leftSide time, unsorted on the stack has a length of 2
	rightSide := mergeSort(unsorted[mid:])
	return merge(leftSide, rightSide)
}

// merge is where the actual sorting happens.

func merge(left []int, right []int) []int {
	var sorted []int
	leftIndex := 0
	rightIndex := 0

	/*
		When it reaches the for loop for the left time, both left and right
		have one item, and right has a value of 3, and left has a value of 10
		and so itll skip the if clause and append 3 to sorted.
	 */
	for leftIndex < len(left) && rightIndex < len(right) {
		leftElement := left[leftIndex]
		rightElement := right[rightIndex]

		if leftElement < rightElement {
			sorted = append(sorted, leftElement)
			leftIndex++
		} else if leftElement > rightElement {
			sorted = append(sorted, rightElement)
			rightIndex++
		} else {
			sorted = append(sorted, leftElement)
			leftIndex++
			sorted = append(sorted, rightElement)
			rightIndex++
		}
	}
	// leftIndex did not get incremented, but rightIndex did, so its going to append leftIndex to sorted
	for leftIndex < len(left) {
		sorted = append(sorted, left[leftIndex])
		leftIndex++
	}
	// since rightIndex got incremented, 1 = 1, and thus nothing happens here
	for rightIndex < len(right) {
		sorted = append(sorted, right[rightIndex])
		rightIndex++
	}
	// we return [3,10]
	return sorted
}

