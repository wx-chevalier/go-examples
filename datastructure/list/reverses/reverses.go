package reverses

import "github.com/Dev-Snippets/algorithm-go-snippets/datastructures/linkedlists/singlylinkedlists"

// Reverse reverse a singly linked list
func Reverse(node *singlylinkedlists.SLLNode) *singlylinkedlists.SLLNode {
	current := node
	var previous *singlylinkedlists.SLLNode
	var next *singlylinkedlists.SLLNode
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
