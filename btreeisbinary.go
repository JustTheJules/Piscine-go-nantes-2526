package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data >= root.Data {
		return false
	} else if root.Right != nil && root.Right.Data <= root.Data {
		return false
	} else {
		return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
	}
}
