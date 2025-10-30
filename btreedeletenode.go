package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	if node.Left == nil && node.Right == nil {
		if node.Parent == nil {
			return nil
		}
		if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
		return root
	}
	if node.Left != nil && node.Right != nil {
		successor := node.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		node.Data = successor.Data
		return BTreeDeleteNode(root, successor)
	}
	child := node.Left
	if child == nil {
		child = node.Right
	}
	if node.Parent == nil {
		child.Parent = nil
		return child
	}
	if node.Parent.Left == node {
		node.Parent.Left = child
	} else {
		node.Parent.Right = child
	}
	child.Parent = node.Parent
	return root
}
