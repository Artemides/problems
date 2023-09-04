package recursion

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func ListBinary() {
	tree := Node{Data: 5, Left: &Node{Data: 3, Left: &Node{Data: 1}}, Right: &Node{Data: 10, Left: &Node{Data: 7}, Right: &Node{Data: 21}}}
	TraverseTree(&tree)
}

func TraverseTree(node *Node) {
	if node == nil {
		return

	}

	if node.Left != nil {
		TraverseTree(node.Left)
		fmt.Println("left")
	}

	fmt.Printf("%v ", node.Data)

	if node.Right != nil {
		TraverseTree(node.Right)
		fmt.Println("right")
	}

}

func Traverse(slc []int) {
	if len(slc) == 0 {
		return
	}
	fmt.Print(slc[0])
	Traverse(slc[1:])
}
