// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-17
// Time: 23:00

package main

import "fmt"

func main() {
	arr := &[6]int{10, 34, 56, 89, 66, 79}
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(arr *[6]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertVal+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
}
