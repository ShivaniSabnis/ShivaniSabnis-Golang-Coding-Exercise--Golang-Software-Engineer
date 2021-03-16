package main

import (
	"fmt"

	"github.com/ShivaniSabnis-Golang-Coding-Exercise--Golang-Software-Engineer/BInaryTree"
	"github.com/ShivaniSabnis-Golang-Coding-Exercise--Golang-Software-Engineer/RobberAssignment"
)

func main() {
	var num int
	fmt.Println("Select 1 or 2 \n1: Robber Problem \n2: Binary Tree")
	fmt.Scanf("%d ", &num)
	switch num {
	case 1:
		RobberAssignment.RobberAssignment()
		break
	case 2:
		BInaryTree.BinaryTree()
		break
	default:
		fmt.Println("Select only between 1 or 2")
	}
}
