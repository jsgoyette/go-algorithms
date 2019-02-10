package main

/*
 * Heap sort - https://en.wikipedia.org/wiki/Heapsort
 */

import "fmt"

import "github.com/jsgoyette/go-algorithms"

func siftDown(arr []int, i int, arrLen int) []int {
	done := false
	maxChild := 0

	for i*2+1 < arrLen && !done {
		if i*2+1 == arrLen-1 {
			maxChild = i*2 + 1
		} else if arr[i*2+1] > arr[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if arr[i] < arr[maxChild] {
			arr[i], arr[maxChild] = arr[maxChild], arr[i]
			i = maxChild
		} else {
			done = true
		}
	}

	return arr
}

func heapify(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		arr = siftDown(arr, i, len(arr))
	}
}

func heapSort(arr []int) {
	for i := len(arr) - 1; i >= 1; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		arr = siftDown(arr, 0, i)
	}
}

func main() {
	arr := utils.RandArray(12)
	fmt.Println("Initial array is:", arr)

	heapify(arr)
	fmt.Println("Heapify array is:", arr)

	heapSort(arr)
	fmt.Println("Sorted array is: ", arr)
}
