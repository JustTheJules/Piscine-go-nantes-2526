package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := BTreeLevelCount(root.Left)
	rightCount := BTreeLevelCount(root.Right)
	if leftCount > rightCount {
		return leftCount + 1
	}
	return rightCount + 1
}
