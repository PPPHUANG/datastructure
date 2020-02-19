// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-18
// Time: 22:47

package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int
	Top    int
	Arr    [5]int
}

func (stack *Stack) push(val int) error {
	if stack.Top == stack.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	stack.Top++
	stack.Arr[stack.Top] = val
	return nil
}

func (stack *Stack) list() {
	if stack.Top == -1 {
		fmt.Println("empty stack")
		return
	}
	for i := stack.Top; i >= 0; i-- {
		fmt.Printf("stack[%d]:%d \n", i, stack.Arr[i])
	}
}

func (stack *Stack) pop() (int, error) {
	if stack.Top == -1 {
		fmt.Println("empty stack")
		return 0, errors.New("empty stack")
	}
	val := stack.Arr[stack.Top]
	stack.Top--
	return val, nil
}
func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	err := stack.push(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = stack.push(2)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = stack.push(3)
	if err != nil {
		fmt.Println(err.Error())
	}
	stack.list()
	val, err := stack.pop()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(val)
	stack.list()
}
