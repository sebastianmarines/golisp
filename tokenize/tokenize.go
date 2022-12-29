package tokenize

type Token struct {
	Type  string
	Value string
}

func Tokenize(str string) []Token {
	var tokens []Token
	var token string
	var inString bool
	for _, c := range str {
		if c == '"' {
			inString = !inString
			if inString {
				token = ""
			} else {
				tokens = append(tokens, Token{"string", token})
				token = ""
			}
		} else if inString {
			token += string(c)
		} else if c == '(' {
			tokens = append(tokens, Token{"(", "("})
		} else if c == ')' {
			if token != "" {
				tokens = append(tokens, Token{"symbol", token})
			}
			tokens = append(tokens, Token{")", ")"})
			token = ""
		} else if c == ' ' {
			if token != "" {
				tokens = append(tokens, Token{"symbol", token})
				token = ""
			}
		} else {
			token += string(c)
		}
	}

	if token != "" {
		tokens = append(tokens, Token{"symbol", token})
	}
	return tokens
}
