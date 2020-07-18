package main

import "fmt"

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func MergeSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}
	done := make(chan bool)
	mid := len(items) / 2
	var left []int

	//In mergesort we divide our array in left side and right side and call the mergesort
	//function on both sides
	//these two call can be made independently and one can be executed in a goroutine
	go func() {
		left = MergeSort(items[:mid])
		done <- true
	}()
	right := MergeSort(items[mid:])
	<-done
	return Merge(left, right)
}

func main() {
	items := []int{12, 12, 56, 54, 8, 7654, 89, 765, 46, 4, 84, 465, 465, 46, 46, 484, 64, 64, 564, 654, 684, 648, 46, 49, 4}
	fmt.Printf("%v\n%v\n", items, MergeSort(items))
}
