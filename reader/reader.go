package reader

import (
	"cheemscript/ast"
	"cheemscript/tokenize"
)

type Reader struct {
	tokens []tokenize.Token

	// The index of the next token to be read.
	index int

	// The current token.
	token string

	// The current token's type.
	tokenType string
}

func NewReader(tokens []tokenize.Token) *Reader {
	return &Reader{
		tokens: tokens,
		index:  0,
	}
}

func (r *Reader) Next() {
	r.token = r.tokens[r.index].Value
	r.tokenType = r.tokens[r.index].Type
	r.index++
}

func (r *Reader) Peek() string {
	return r.tokens[r.index].Value
}

func (r *Reader) read() *ast.Node {
	switch r.tokenType {
	case "symbol":
		return r.readSymbol()
	case "(":
		return r.readList()
	case "string":
		return r.readString()
	default:
		panic("unknown token type")
	}
}

func (r *Reader) readSymbol() *ast.Node {
	return &ast.Node{
		Type:  ast.Symbol,
		Value: r.token,
	}
}

func (r *Reader) readList() *ast.Node {
	r.Next()
	var children []*ast.Node
	for r.token != ")" {
		children = append(children, r.read())
		r.Next()
	}
	return &ast.Node{
		Type:     ast.List,
		Children: children,
	}
}

func (r *Reader) readString() *ast.Node {
	return &ast.Node{
		Type:  ast.String,
		Value: r.token,
	}
}

func (r *Reader) Read() *ast.Node {
	r.Next()
	return r.read()
}

func Read(s string) *ast.Node {
	tokens := tokenize.Tokenize(s)
	return NewReader(tokens).Read()
}
