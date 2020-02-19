// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-19
// Time: 16:50

package main

import "fmt"

func findWay(myMap *[8][7]int, i int, j int) bool {
	if myMap[6][5] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 {
			myMap[i][j] = 2
			if findWay(myMap, i+1, j) {
				return true
			} else if findWay(myMap, i, j+1) {
				return true
			} else if findWay(myMap, i-1, j) {
				return true
			} else if findWay(myMap, i, j-1) {
				return true
			} else {
				myMap[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func main() {

	/*
		1 1 1 1 1 1 1
		1 0 0 0 0 0 1
		1 0 0 0 0 0 1
		1 1 1 0 0 0 1
		1 0 0 0 0 0 1
		1 0 0 0 0 0 1
		1 0 0 0 0 0 1
		1 1 1 1 1 1 1

		1 1 1 1 1 1 1
		1 2 0 0 0 0 1
		1 2 2 2 0 0 1
		1 1 1 2 0 0 1
		1 0 0 2 0 0 1
		1 0 0 2 0 0 1
		1 0 0 2 2 2 1
		1 1 1 1 1 1 1
	*/
	var myMap [8][7]int
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}

	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	//0为没走过的路 1为墙 2为通路 3为死路
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
	//小猫探路
	findWay(&myMap, 1, 1)
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
