package ast

import (
	"fmt"
	"strconv"
)

type NodeType int

const (
	_ NodeType = iota
	List
	Symbol
	Number
	String
	InternalFunction
	Nil
	True
	False
	Function
)

type Node struct {
	Type     NodeType
	Value    interface{}
	Children []*Node
}

func (n *Node) PrStr(printReadability bool) string {
	switch n.Type {
	case List:
		return n.listString()
	case Number:
		return n.numberString()
	case String:
		return n.stringString(printReadability)
	case Symbol:
		return n.symbolString()
	case Nil:
		return "nil"
	case True:
		return "true"
	case False:
		return "false"
	case Function:
		return "#<function>"
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
		str += child.PrStr(false) + " "
	}
	return "(" + str[:len(str)-1] + ")"
}

func (n *Node) numberString() string {
	t := strconv.Itoa(n.Value.(int))
	return t
}

func (n *Node) stringString(printReadability bool) string {
	var str string
	if printReadability {
		str = fmt.Sprintf("%v", n.Value)
	} else {
		str = fmt.Sprintf("%q", n.Value)
	}
	return str
}

func (n *Node) symbolString() string {
	return n.Value.(string)
}
