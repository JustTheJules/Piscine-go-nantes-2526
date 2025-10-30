package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil || rplc == nil {
		return root
	}
	if root == node {
		rplc.Left = root.Left
		rplc.Right = root.Right
		return rplc
	}
	parent := findParent(root, node)
	if parent == nil {
		return root
	}
	if parent.Left == node {
		parent.Left = rplc
	}
	if parent.Right == node {
		parent.Right = rplc
	}
	rplc.Left = node.Left
	rplc.Right = node.Right
	return root
}

func findParent(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return nil
	}
	if root.Left == node || root.Right == node {
		return root
	}
	if node.Data < root.Data {
		return findParent(root.Left, node)
	}
	return findParent(root.Right, node)
}
