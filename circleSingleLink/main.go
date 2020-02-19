// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-06
// Time: 20:45

package main

import "fmt"

type CatNode struct {
	id   int
	name string
	next *CatNode
}

func insertCatNode(head *CatNode, newCatNode *CatNode) {
	temp := head
	if temp.next == nil {
		head.id = newCatNode.id
		head.name = newCatNode.name
		head.next = head
		return
	}

	for {
		if temp.next == head {
			temp.next = newCatNode
			newCatNode.next = head
			break
		}
		temp = temp.next
	}
}

func delCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("Empty Link")
		return head
	}
	if temp.next == head {
		if temp.id == id {
			temp.id = 0
			temp.name = ""
			temp.next = nil
			return head
		} else {
			fmt.Printf("Only Head,%d Not Exist\n", id)
			return head
		}
	}
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	flag := true
	for {
		if temp.next == head {
			break
		}
		if temp.id == id {
			if temp == head {
				head = temp.next
			}
			helper.next = temp.next
			fmt.Printf("Deleted CatNode:%d\n", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}
	if flag {
		if temp.id == id {
			helper.next = temp.next
			fmt.Printf("Deleted CatNode:%d\n", id)
		} else {
			fmt.Printf("%d Not Exist", id)
		}
	}
	return head
}

func listCatNode(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("Empty Link")
		return
	}
	for {
		fmt.Println("猫的信息为=", temp.id, temp.name, "->")
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	head := &CatNode{}
	cat1 := &CatNode{
		id:   1,
		name: "cat1",
	}
	cat2 := &CatNode{
		id:   2,
		name: "cat2",
	}
	cat3 := &CatNode{
		id:   3,
		name: "cat3",
	}
	insertCatNode(head, cat1)
	insertCatNode(head, cat2)
	insertCatNode(head, cat3)
	listCatNode(head)
	head = delCatNode(head, 2)
	listCatNode(head)
}
