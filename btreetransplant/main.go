package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	node := piscine.BTreeSearchItem(root, "1")
	rplc := &piscine.TreeNode{Data: "3"}
	root = piscine.BTreeTransplant(root, node, rplc)
	piscine.BTreeApplyInorder(root, fmt.Println)
}
