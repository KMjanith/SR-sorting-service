package service

import (
	"context"
	"sorting/producers"
	"time"
)

func Merge(arr []int64, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	leftArr := make([]int64, n1)
	rightArr := make([]int64, n2)

	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
	}

	for i := 0; i < n2; i++ {
		rightArr[i] = arr[mid+1+i]
	}

	i, j, k := 0, 0, left

	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = leftArr[i]
		k++
		i++
	}

	for j < n2 {
		arr[k] = rightArr[j]
		k++
		j++
	}
}

func MergeSort(arr []int64, left, right int, ctx context.Context,ReplyTo string) {

	startTime := time.Now()

	if left < right {
		mid := left + (right-left)/2

		MergeSort(arr, left, mid, ctx,ReplyTo)
		MergeSort(arr, mid+1, right, ctx,ReplyTo)

		Merge(arr, left, mid, right)
	}

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	duration := timeTaken.String() // Convert the duration to a string

	producers.SendSortedREsponse("merge", arr, duration, ctx,ReplyTo)

}
