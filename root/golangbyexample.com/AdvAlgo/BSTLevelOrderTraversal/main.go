package main

import "fmt"

/* 1675000.15.06.22

L1			1
L2		2		3
L3 	4		5		6

OUTPUT - [1], [2, 3], [4, 5, 6]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	levelMapNode := make(map[int][]int)
	visit(root, 0, &levelMapNode)

	output := make([][]int, 0)

	i := 0
	for {
		_, ok := levelMapNode[i]

		if ok {
			output = append(output, levelMapNode[i])
		} else {
			break
		}
		i = i + 1
	}
	return output
}

func visit(root *TreeNode, level int, levelMapNode *map[int][]int) {
	if root == nil {
		return
	}

	_, ok := (*levelMapNode)[level]
	if ok {
		(*levelMapNode)[level] = append((*levelMapNode)[level], root.Val)
	} else {
		(*levelMapNode)[level] = []int{root.Val}
	}

	if root.Left != nil {
		visit(root.Left, level+1, levelMapNode)
	}

	if root.Right != nil {
		visit(root.Right, level+1, levelMapNode)
	}
	return
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	output := levelOrder(&root)
	fmt.Println(output)
}
