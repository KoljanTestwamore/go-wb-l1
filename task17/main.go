package main

import "fmt"

func binarySearch(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
		case len(a) == 0:
			result = -1 // not found
		case a[mid] > search:
			result, searchCount = binarySearch(a[:mid], search)
		case a[mid] < search:
			result, searchCount = binarySearch(a[mid+1:], search)
			if result >= 0 { // if anything but the -1 "not found" result
				result += mid + 1
			}
		default: // a[mid] == search
			result = mid // found
	}
	searchCount++
	return
}

func main() {
	res, _ := binarySearch([]int{1, 2, 3, 4, 5, 6}, 6)
	fmt.Printf("%v", res)
}