// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-20
// Time: 11:55

package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (emp *Emp) Show() {
	fmt.Printf("find emp %d in Link %d", emp.Id%7, emp.Id)
	fmt.Println()
}

type EmpLink struct {
	Head *Emp
}

func (empLink *EmpLink) Insert(emp *Emp) {
	cur := empLink.Head
	var pre *Emp = nil
	if cur == nil {
		empLink.Head = emp
		return
	}
	for {
		if cur != nil {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	if pre != nil {
		pre.Next = emp
	} else {
		empLink.Head = emp
	}
	emp.Next = cur
}

func (empLink *EmpLink) ShowLink(no int) {
	if empLink.Head == nil {
		fmt.Printf("Link %d empty", no)
	}
	cur := empLink.Head
	for {
		if cur != nil {
			fmt.Printf("Link %d id %d name %s ->", no, cur.Id, cur.Name)
		} else {
			break
		}
		cur = cur.Next
	}
	fmt.Println()
}

func (empLink *EmpLink) FindById(id int) *Emp {
	cur := empLink.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			return &Emp{}
		}
		cur = cur.Next
	}
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (hashTable *HashTable) Insert(emp *Emp) {
	linkNo := hashTable.HashFun(emp.Id)
	hashTable.LinkArr[linkNo].Insert(emp)
}

func (hashTable *HashTable) HashFun(id int) int {
	return id % 7
}

func (hashTable *HashTable) ShowAll() {
	for i := 0; i < len(hashTable.LinkArr); i++ {
		hashTable.LinkArr[i].ShowLink(i)
	}
}

func (hashTable *HashTable) FindById(id int) *Emp {
	linkNo := hashTable.HashFun(id)
	return hashTable.LinkArr[linkNo].FindById(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("==================emp sys===================")
		fmt.Println("input add emp")
		fmt.Println("show list emp")
		fmt.Println("find find emp")
		fmt.Println("exit exit sys")
		fmt.Println("请输入你的选择")
		_, _ = fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员ID")
			_, _ = fmt.Scanln(&id)
			fmt.Println("请输入雇员Name")
			_, _ = fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "find":
			fmt.Println("请输入雇员ID")
			_, _ = fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp != nil {
				emp.Show()
			}
		case "show":
			hashTable.ShowAll()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
