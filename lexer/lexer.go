package lexer

type Token struct {
	Type  TokenType
	Value string
}

type TokenType int

const (
	_ TokenType = iota
	LeftParen
	RightParen
	Symbol
	String
	Integer
	EOF
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func Tokenize(input string) []Token {
	l := &Lexer{input: input}
	l.readChar()
	var tokens []Token
	for {
		tok := l.NextToken()
		if tok.Type == EOF {
			break
		}
		tokens = append(tokens, tok)
	}
	return tokens
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = newToken(LeftParen, l.ch)
	case ')':
		tok = newToken(RightParen, l.ch)
	case 0:
		tok.Value = ""
		tok.Type = EOF
	default:
		if isDigit(l.ch) {
			// TODO: handle decimal numbers
			tok.Type = Integer
			tok.Value = l.readNumber()
			return tok
		} else {
			if l.ch == '"' {
				tok.Type = String
				tok.Value = l.readString()
				l.readChar()
				return tok
			} else if l.ch == '-' {
				if isDigit(l.peekChar()) {
					l.readChar()
					tok.Type = Integer
					tok.Value = "-" + l.readNumber()
					return tok
				}
			} else {
				tok.Type = Symbol
				tok.Value = l.readSymbol()
				return tok
			}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	// T
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	str := ""
	for {
		l.readChar()
		if l.ch == '\\' {
			l.readChar()
			str += string(l.ch)
			continue
		}
		if l.ch == '"' || l.ch == 0 {
			break
		}
		str += string(l.ch)
	}
	return str
}

func (l *Lexer) readSymbol() string {
	position := l.position

	for isSymbol(l.ch) && l.ch != 0 {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Value: string(ch)}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSymbol(ch byte) bool {
	return !isDigit(ch) && ch != ' ' && ch != '\t' && ch != '\n' && ch != '\r' && ch != '(' && ch != ')' && ch != '"'
}
