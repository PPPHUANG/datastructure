// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-04
// Time: 12:00
//单向链表

package main

import "fmt"

type HeroNode struct {
	id       int
	name     string
	nickName string
	next     *HeroNode
}

//无序插入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
}

//有序插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.id > newHeroNode.id {
			break
		} else if temp.next.id == newHeroNode.id {
			flag = false
			break
		}
		temp = temp.next
	}
	if flag {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	} else {
		fmt.Println("Exist HeroNode")
		return
	}
}
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.id == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("Not Exist HeroNode")
	}
}
func EditHeroNode(head *HeroNode, newHeroHead *HeroNode) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.id == newHeroHead.id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next.name = newHeroHead.name
		temp.next.nickName = newHeroHead.nickName
	} else {
		fmt.Println("Not Exist HeroNode")
	}
}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("Empty Link")
	}
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.next.id, temp.next.name, temp.next.nickName)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
func main() {
	head := &HeroNode{}
	head2 := &HeroNode{}
	hero1 := &HeroNode{
		id:       1,
		name:     "宋江",
		nickName: "及时雨",
	}
	hero2 := &HeroNode{
		id:       2,
		name:     "卢俊义",
		nickName: "玉麒麟",
	}
	hero3 := &HeroNode{
		id:       3,
		name:     "林冲",
		nickName: "豹子头",
	}
	//heroEdit := &HeroNode{
	//	id:       3,
	//	name:     "林冲1",
	//	nickName: "豹子头1",
	//}
	//无序插入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero3)
	InsertHeroNode(head, hero2)
	ListHeroNode(head)
	//有序插入
	InsertHeroNode2(head2, hero1)
	InsertHeroNode2(head2, hero3)
	InsertHeroNode2(head2, hero2)
	//DelHeroNode(head2, 2)
	//EditHeroNode(head2, heroEdit)
	ListHeroNode(head2)
}
