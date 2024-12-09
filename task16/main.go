package main

import "fmt"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSortHelp(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSortHelp(arr, low, p-1)
		arr = quickSortHelp(arr, p+1, high)
	}

	return arr
}

func quickSort(arr []int) []int {
	return quickSortHelp(arr, 0, len(arr)-1)
}

func main() {
	fmt.Printf("%v", quickSort([]int{ 9, 8 ,7,6,5,4,3,2}))
}