package service

import (
	"context"
	"sorting/producers"
	"time"
)

func BubbleSort(arr []int64, ctx context.Context, replyTo string) {

	startTime := time.Now()

	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j] + arr[j+1]
				arr[j+1] = arr[j] - arr[j+1]
				arr[j] = arr[j] - arr[j+1]
			}
		}
	}

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	duration := timeTaken.String() // Convert the duration to a string

	producers.SendSortedREsponse("bubble", arr, duration, ctx, replyTo)

}
