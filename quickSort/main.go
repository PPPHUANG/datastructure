// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-17
// Time: 23:38

package main

import (
	"fmt"
)

func main() {
	arr := &[6]int{10, 34, 56, 89, 66, 79}
	quickSort(0, 5, arr)
	fmt.Println(arr)
}

func quickSort(left int, right int, arr *[6]int) {
	l := left
	r := right
	pivot := arr[(left+right)/2]
	temp := 0
	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		temp = arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		quickSort(left, r, arr)
	}
	if right > l {
		quickSort(l, right, arr)
	}
}
