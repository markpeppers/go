package leveltraverse

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	o := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	level := 0
	queue = append(queue, root) // add to Q the root node
	for len(queue) > 0 {
		n := len(queue) // Number of nodes in current level
		o = append(o, []int{})
		o[level] = make([]int, 0, n)
		for i := 0; i < n; i++ {
			queue = append(queue, helper(queue[i])...)
			o[level] = append(o[level], queue[i].Val)
		}
		queue = queue[n:]
		level++
		// add all child nodes for level in the queue
	}
	return o

}

func helper(node *TreeNode) []*TreeNode {
	// Takes a node, returns list of all child nodes
	if node == nil {
		return nil
	}
	o := []*TreeNode{}
	if node.Left != nil {
		o = append(o, node.Left)
	}
	if node.Right != nil {
		o = append(o, node.Right)
	}
	return o
}
