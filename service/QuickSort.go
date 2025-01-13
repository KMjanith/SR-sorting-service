package service

import (
	"context"
	"sorting/producers"
	"time"
)

// Partition function for QuickSort
func Partition(arr []int64, low, high int) int {
	pivot := arr[high] // Choose the last element as pivot
	i := low - 1       // Pointer for the smaller element

	for j := low; j < high; j++ {
		// If the current element is smaller than the pivot
		if arr[j] < pivot {
			i++
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Swap the pivot element with the element at i+1
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1 // Return the partition index
}

// QuickSort function to sort the array
func QuickSort(arr []int64, low, high int, ctx context.Context, ReplyTo string) []int64 {
	if low < high {
		pi := Partition(arr, low, high) // Partitioning index

		// Recursively sort the two halves
		QuickSort(arr, low, pi-1, ctx, ReplyTo)
		QuickSort(arr, pi+1, high, ctx, ReplyTo)
	}

	// Send the sorted array and duration only after the final recursion returns
	if low == 0 && high == len(arr)-1 {
		startTime := time.Now()
		endTime := time.Now()
		timeTaken := endTime.Sub(startTime)
		duration := timeTaken.String() // Convert the duration to a string

		producers.SendSortedREsponse("quick", arr, duration, ctx, ReplyTo)
	}

	return arr
}
