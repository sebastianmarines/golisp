package reader

import (
	"golisp/ast"
	"golisp/lexer"
	"strconv"
)

type Reader struct {
	tokens []lexer.Token

	// The index of the next token to be read.
	index int

	// The current token.
	token string

	// The current token's type.
	tokenType lexer.TokenType
}

func NewReader(tokens []lexer.Token) *Reader {
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
	case lexer.Symbol:
		return r.readSymbol()
	case lexer.LeftParen:
		return r.readList()
	case lexer.String:
		return r.readString()
	case lexer.Integer:
		return r.readNumber()

	default:
		panic("unknown token type")
	}
}

func (r *Reader) readSymbol() *ast.Node {
	switch r.token {
	case "nil":
		return &ast.Node{
			Type:  ast.Nil,
			Value: nil,
		}
	case "true":
		return &ast.Node{
			Type:  ast.True,
			Value: true,
		}
	case "false":
		return &ast.Node{
			Type:  ast.False,
			Value: false,
		}
	default:
		return &ast.Node{
			Type:  ast.Symbol,
			Value: r.token,
		}

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

func (r *Reader) readNumber() *ast.Node {
	i, err := strconv.Atoi(r.token)
	if err != nil {
		panic(err)
	}
	return &ast.Node{
		Type:  ast.Number,
		Value: i,
	}
}

func (r *Reader) Read() *ast.Node {
	r.Next()
	return r.read()
}

func Read(s string) *ast.Node {
	tokens := lexer.Tokenize(s)
	return NewReader(tokens).Read()
}
