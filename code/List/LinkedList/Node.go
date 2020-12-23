package LinkedList

type Node struct {
	data interface{}
	pNext *Node
	pPrev *Node
}

func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

