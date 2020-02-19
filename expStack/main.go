// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-19
// Time: 15:11

package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int
	Top    int
	Arr    [20]int
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

func (stack *Stack) isOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

func (stack *Stack) cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("oper err")
	}
	return res
}

func (stack *Stack) priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	}
	if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "30+2*6-2"
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	res := 0
	keepNum := ""
	for {
		ch := exp[index : index+1]
		temp := int([]byte(ch)[0]) //chr -> ASCII
		if operStack.isOper(temp) {
			if operStack.Top == -1 {
				_ = operStack.push(temp)
			} else {
				if operStack.priority(temp) <= operStack.priority(operStack.Arr[operStack.Top]) {
					num1, _ = numStack.pop()
					num2, _ = numStack.pop()
					oper, _ = operStack.pop()
					res = operStack.cal(num1, num2, oper)
					_ = numStack.push(res)
					_ = operStack.push(temp)
				} else {
					_ = operStack.push(temp)
				}
			}
		} else {
			keepNum += ch
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				_ = numStack.push(int(val))
			} else {
				if operStack.isOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					_ = numStack.push(int(val))
					keepNum = ""
				}
			}

		}
		if index == len(exp)-1 {
			break
		}
		index++
	}

	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.pop()
		num2, _ = numStack.pop()
		oper, _ = operStack.pop()
		res = operStack.cal(num1, num2, oper)
		_ = numStack.push(res)
	}
	res, _ = numStack.pop()
	fmt.Println(res)
}
