package main

import "fmt"

func BinarySearch(a []int, target int, lowIndex int, highIndex int) int {
	// Details: https://www.khanacademy.org/computing/computer-science/algorithms/binary-search/a/implementing-binary-search-of-an-array
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if a[mid] > target {
		return BinarySearch(a, target, lowIndex, mid)
	} else if a[mid] < target {
		return BinarySearch(a, target, mid + 1, highIndex)
	} else {
		return mid
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res1 := BinarySearch(array, 6, 0, len(array))
	fmt.Println(res1)
	res2 := BinarySearch(array, 11, 0, len(array) -1 )
	fmt.Println(res2)
}

