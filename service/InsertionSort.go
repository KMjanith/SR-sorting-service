package service

import (
	"context"
	"sorting/producers"
	"time"
)

func InsertionSort(arr []int64, ctx context.Context,ReplyTo string) {

	startTime := time.Now()

	n := len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	duration := timeTaken.String() // Convert the duration to a string

	producers.SendSortedREsponse("insertion", arr, duration, ctx, ReplyTo)
}
