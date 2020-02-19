// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-01-17
// Time: 00:02

package main

import "fmt"

type valNode struct {
	row int
	col int
	val int
}
func main()  {
	//原始数据组
	sourceArr := [11][11]int {}
	sourceArr[2][3] = 1
	sourceArr[3][4] = 2
	for _, val := range sourceArr {
		for _, val1 := range val  {
			fmt.Printf("%d\t", val1)
		}
		fmt.Println()
	}

	//稀疏数组
	var sparseArr []valNode
	valNode := valNode{11,11,0}
	sparseArr = append(sparseArr, valNode)
	for i, val := range sourceArr {
		for j, val1 := range val  {
			if val1 != 0 {
				valNode.row = i
				valNode.col = j
				valNode.val = val1
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	for _, val := range sparseArr {
		fmt.Printf("%d %d %d\n", val.row, val.col, val.val)
	}

	//稀疏数组转成原始数组 正常情况稀疏数组数据需要落盘 从盘中恢复到原始数组
	var sourceArr2 [11][11]int

	for _, val := range sparseArr {
		if val.val != 0  {
			sourceArr2[val.row][val.col] = val.val
		}
	}

	for _, val := range sourceArr2 {
		for _, val1 := range val  {
			fmt.Printf("%d\t", val1)
		}
		fmt.Println()
	}
}
