// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-04
// Time: 15:02
//双向链表

package main

import "fmt"

type HeroNode struct {
	id       int
	name     string
	nickName string
	pre      *HeroNode
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
	newHeroNode.pre = temp
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
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
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
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("Not Exist HeroNode")
	}
}

//正序打印
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

//逆序打印
func ListHeroNode2(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("Empty Link")
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.id, temp.name, temp.nickName)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
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

func main() {
	head := &HeroNode{}
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
	//DelHeroNode(head, 2)
	//EditHeroNode(head, heroEdit)
	ListHeroNode2(head)

}
