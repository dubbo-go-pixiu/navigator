package trie

func BuildNode(identify string) Node {
	var node = Node{identify: identify}
	return node
}

func BuildNodeWithDefault(identify string, supplier func() any) Node {
	var node = Node{
		identify: identify,
		data:     supplier(),
	}
	return node
}

type Node struct {
	identify string
	children map[any]*Node
	parent   *Node
	data     any
}

func (parent *Node) BuildChild(identify string) Node {
	var node = Node{
		identify: identify,
		parent:   parent,
	}
	return node
}

func (parent *Node) BuildChildWithDefault(identify string, supplier func() any) Node {
	var node = Node{
		identify: identify,
		data:     supplier(),
		parent:   parent,
	}
	return node
}
