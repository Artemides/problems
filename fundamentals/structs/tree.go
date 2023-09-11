package structs

type Tree struct {
	value       int
	left, right *Tree
}

func Sort(values []int) {
	var root *Tree
	for _, value := range values {
		root = add(root, value)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, tree *Tree) []int {
	if tree != nil {
		values = appendValues(values, tree.left)
		values = append(values, tree.value)
		values = appendValues(values, tree.right)
	}
	return values
}

func add(tree *Tree, value int) *Tree {
	if tree == nil {
		tree = new(Tree)
		tree.value = value
		return tree
	}

	if value < tree.value {
		tree.left = add(tree.left, value)
	} else {
		tree.right = add(tree.right, value)
	}
	return tree
}
