package BInaryTree

import (
	"fmt"
	"sync"
)

type Value struct {
	value  int
	leftn  *Value
	rightn *Value
}

type BTree struct {
	root *Value
	lock sync.RWMutex
}

func BinaryTree() {
	var num int
	var btree *BTree = &BTree{}
cont:
	fmt.Println("Select /2/3/4 \n1: Insert to tree \n2: InorderTraversal \n3: PreorderTraversal \n4:PostOrderTraversal")
	fmt.Scanf("%d\n", &num)

	switch num {
	case 1:
		btree.Insert()
		goto cont
	case 2:
		btree.Inorder()
		goto cont
	case 3:
		btree.Preorder()
		goto cont
	case 4:
		btree.Postorder()
		goto cont
	default:
		fmt.Println("BYE")
	}

}

func (btree *BTree) Insert() {
	fmt.Println("Enter value to be inserted\n")
	var val int
	fmt.Scanf("%d\n", &val)
	btree.lock.Lock()
	defer btree.lock.Unlock()
	var value *Value
	value = &Value{value: val}

	if btree.root == nil {
		btree.root = value
	} else {
		InsertLR(btree.root, value)
	}
}

func InsertLR(root *Value, value *Value) {
	if value.value < root.value {
		if root.leftn == nil {
			root.leftn = value
		} else {
			InsertLR(root.leftn, value)
		}
	} else {
		if root.rightn == nil {
			root.rightn = value
		} else {
			InsertLR(root.rightn, value)
		}
	}
}
func (btree *BTree) Inorder() {
	btree.lock.Lock()
	defer btree.lock.Unlock()
	fmt.Println("Inorder")
	Inorderfunc(btree.root)
}

func Inorderfunc(root *Value) {
	if root != nil {
		Inorderfunc(root.leftn)
		fmt.Println(root.value)
		Inorderfunc(root.rightn)
	}
}

func (btree *BTree) Preorder() {
	btree.lock.Lock()
	defer btree.lock.Unlock()
	fmt.Println("Preorder")
	Preorderfunc(btree.root)
}

func Preorderfunc(root *Value) {
	if root != nil {
		fmt.Println(root.value)
		Preorderfunc(root.leftn)
		Preorderfunc(root.rightn)
	}
}

func (btree *BTree) Postorder() {
	btree.lock.Lock()
	defer btree.lock.Unlock()
	fmt.Println("Postorder")
	Postorderfunc(btree.root)
}

func Postorderfunc(root *Value) {
	if root != nil {
		Postorderfunc(root.leftn)
		Postorderfunc(root.rightn)
		fmt.Println(root.value)
	}
}
