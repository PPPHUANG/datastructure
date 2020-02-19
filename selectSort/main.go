// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-17
// Time: 22:29

package main

import "fmt"

func main() {
	arr := &[6]int{10, 34, 56, 89, 66, 79}
	selectSort(arr)
	fmt.Println(arr)
}

func selectSort(arr *[6]int) {
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[j] {
				max = arr[j]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
}
