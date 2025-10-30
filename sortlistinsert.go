package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref, Next: nil}
	if l == nil || l.Data >= data_ref {
		newNode.Next = l
		return newNode
	}
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	return l
}
