package httpapi

import files "github.com/ipfs/go-ipfs-files"

// Node implements a files.Node associated with name.
type Node struct {
	files.Node
	name string
}

// NewNode returns a new Node from given parameters.
func NewNode(name string, node files.Node) *Node {
	return &Node{
		Node: node,
		name: name,
	}
}


