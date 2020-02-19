// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-16
// Time: 22:53

package main

import "fmt"

type Boy struct {
	ID   int
	Next *Boy
}

func addBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("num value error")
		return first
	}
	for i := 1; i <= num; i++ {
		boy := &Boy{
			ID: i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = boy
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

func showBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("empty Link")
	}
	curBoy := first
	for {
		fmt.Printf("boy id = %d -> \n", curBoy.ID)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func playGame(first *Boy, startNum int, countNum int) {
	if first.Next == nil {
		fmt.Println("empty Link")
		return
	}
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	for i := 1; i < startNum; i++ {
		first = first.Next
		tail = tail.Next
	}

	for {
		for i := 1; i < countNum; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("boy %d out \n", first.ID)
		first = first.Next
		tail.Next = first
		if tail == first {
			break
		}
	}
	fmt.Printf("boy %d out \n", first.ID)
}
func main() {
	first := addBoy(5)
	showBoy(first)
	playGame(first, 2, 3)
}
