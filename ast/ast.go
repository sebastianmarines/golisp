package ast

type NodeType int

const (
	_ NodeType = iota
	List
	Symbol
	Number
	String
)

type Node struct {
	Type     NodeType
	Value    interface{}
	Children []*Node
}

func (n *Node) String() string {
	switch n.Type {
	case List:
		return n.listString()
	case Number:
		return n.numberString()
	case String:
		return n.stringString()
	case Symbol:
		return n.symbolString()
	default:
		panic("unknown node type")
	}
}

func (n *Node) listString() string {
	var str string
	if len(n.Children) == 0 {
		return "()"
	}
	for _, child := range n.Children {
		str += child.String() + " "
	}
	return "(" + str[:len(str)-1] + ")"
}

func (n *Node) numberString() string {
	return n.Value.(string)
}

func (n *Node) stringString() string {
	return "\"" + n.Value.(string) + "\""
}

func (n *Node) symbolString() string {
	return n.Value.(string)
}