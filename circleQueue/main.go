// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-02
// Time: 21:08
// 环形数组队列

package main

import (
	"errors"
	"fmt"
	"os"
)

//队列结构体
type CircleQueue struct {
	maxSize int    //最大容量 实际可以maxSize - 1
	array   [5]int //队列数字
	head    int    //头 0
	tail    int    //尾 0
}

//存
func (circleQueue *CircleQueue) push(val int) error {
	if circleQueue.isFull() {
		return errors.New("full queue")
	}
	circleQueue.array[circleQueue.tail] = val
	circleQueue.tail = (circleQueue.tail + 1) % circleQueue.maxSize
	return nil
}

//取
func (circleQueue *CircleQueue) pop() (int, error) {
	if circleQueue.isEmpty() {
		return 0, errors.New("empty queue")
	}
	val := circleQueue.array[circleQueue.head]
	circleQueue.head = (circleQueue.head + 1) % circleQueue.maxSize
	return val, nil
}

//判空
func (circleQueue *CircleQueue) isEmpty() bool {
	return circleQueue.tail == circleQueue.head
}

//判满
func (circleQueue *CircleQueue) isFull() bool {
	return (circleQueue.tail+1)%circleQueue.maxSize == circleQueue.head
}

//遍历数据
func (circleQueue *CircleQueue) listQueue() {
	fmt.Println("queue list:")
	length := circleQueue.length()
	if length == 0 {
		fmt.Println("empty queue")
	}
	tmp := circleQueue.head
	for i := 0; i < length; i++ {
		fmt.Printf("arr[%d]:%d\t", tmp, circleQueue.array[tmp])
		tmp = (tmp + 1) % circleQueue.maxSize
	}
	fmt.Println()
}

//获取长度
func (circleQueue *CircleQueue) length() int {
	return (circleQueue.tail + circleQueue.maxSize - circleQueue.head) % circleQueue.maxSize
}

func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 请输入add 表示添加数据到队列")
		fmt.Println("2. 请输入get 表示从队列获取数据")
		fmt.Println("3. 请输入list 表示列举队列")
		fmt.Println("4. 请输入exit 表示退出")
		_, _ = fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入添加数字")
			_, _ = fmt.Scanln(&val)
			err := queue.push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("添加成功")
			}
		case "get":
			val, err := queue.pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(val)
			}
		case "list":
			queue.listQueue()
		case "exit":
			os.Exit(0)
		}

	}
}
