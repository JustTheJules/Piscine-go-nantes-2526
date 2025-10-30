package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	} else {
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			f(current.Data)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}
}
