package service

import (
	"context"
	"sorting/producers"
	"time"
)

func SelctionSort(arr []int64, ctx context.Context, replyTo string) []int64 {
	startTime := time.Now()

	n := len(arr)

	for i := 0; i < n; i++ {
		j := i
		k := i
		current_minimum := arr[j]
		for j < n {
			if arr[j] < current_minimum {
				current_minimum = arr[j]
				k = j
			}
			j++
		}
		arr[i], arr[k] = arr[k], arr[i]
	}

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	duration := timeTaken.String() // Convert the duration to a string

	producers.SendSortedREsponse("selection", arr, duration, ctx, replyTo)

	return arr
}
