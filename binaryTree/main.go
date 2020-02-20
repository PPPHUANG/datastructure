// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-20
// Time: 15:10

package main

import "fmt"

type Hero struct {
	Id    int
	Name  string
	Left  *Hero
	Right *Hero
}

func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("node=%d name=%s\n", node.Id, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}
func InsideOrder(node *Hero) {
	if node != nil {
		InsideOrder(node.Left)
		fmt.Printf("node=%d name=%s\n", node.Id, node.Name)
		InsideOrder(node.Right)
	}
}
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("node=%d name=%s\n", node.Id, node.Name)
	}
}

func main() {
	root := &Hero{
		Id:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		Id:   2,
		Name: "吴用",
	}

	node10 := &Hero{
		Id:   10,
		Name: "tom",
	}
	node12 := &Hero{
		Id:   12,
		Name: "jace",
	}
	left1.Left = node10
	left1.Right = node12
	right1 := &Hero{
		Id:   3,
		Name: "卢俊义",
	}
	root.Left = left1
	root.Right = right1

	right2 := &Hero{
		Id:   4,
		Name: "林冲",
	}
	right1.Right = right2
	PreOrder(root)
	fmt.Println()
	InsideOrder(root)
	fmt.Println()
	PostOrder(root)
}
