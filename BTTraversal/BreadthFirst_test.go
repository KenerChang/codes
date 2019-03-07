package BTTraversal

import (
	"testing"
	"fmt"
)

func TestBreadthFirst(t *testing.T) {
	tree := &Node {
		Value:2,
		LeftNode: &Node{
			Value: 1,
			LeftNode: &Node{
				Value:0,
			},
			RightNode: &Node {
				Value:7,
			},
		},
		RightNode: &Node{
			Value: 3,
			LeftNode: &Node{
				Value: 9,
			},
			RightNode: &Node{
				Value: 1,
			},
		},
	}

	target := []int{2,1,3,0,7,9,1}
	result := BreadthFirst(tree)

	fmt.Printf("result: %+v\n", result)

	for i, v := range result  {
		if v != target[i] {
			t.Errorf("expect %d, but got %d", target[i], v)
			break
		}
	}
}