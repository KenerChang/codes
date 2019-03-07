package BTTraversal

type Node struct {
	Value int
	LeftNode *Node
	RightNode *Node
	Viewed bool
}

// InOrder take root node as input
// and output values in traversal order
func BreadthFirst(root *Node) []int {
	values := []int{}
	queue := []*Node {
		root,
	}

	for len(queue) != 0 {
		node := queue[0] // get first element in the queue
		values = append(values, node.Value)

		// put left node to the queue
		if node.LeftNode != nil && !node.LeftNode.Viewed {
			queue = append(queue, node.LeftNode)
		}

		// put right node to the queue
		if node.RightNode != nil && !node.RightNode.Viewed {
			queue = append(queue, node.RightNode)
		}

		// pop this node from queue
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []*Node{}
		}

		// set this node as viewed to prevent loop
		node.Viewed = true
	}
	return values
}