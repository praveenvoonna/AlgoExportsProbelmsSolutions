package main

import "fmt"

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	sums := []int{}
	calculateBranchSums(root, 0, &sums)
	return sums
}

func calculateBranchSums(root *BinaryTree, sum int, sums *[]int) {
	if root.Left == nil && root.Right == nil {
		*sums = append(*sums, sum)
		return
	}
	sum += root.Value
	if root.Left != nil {
		calculateBranchSums(root.Left, sum, sums)
	} else if root.Right != nil {
		calculateBranchSums(root.Right, sum, sums)
	}
}

func main() {
	tree := NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(BranchSums(tree))
}

func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}
	tree.Insert(values, 0)
	return tree
}

func (tree *BinaryTree) Insert(values []int, i int) *BinaryTree {
	if i >= len(values) {
		return tree
	}
	val := values[i]

	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		var current *BinaryTree
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Right)
	}

	tree.Insert(values, i+1)
	return tree
}
