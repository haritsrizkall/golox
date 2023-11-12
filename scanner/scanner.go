package scanner

import (
	"fmt"
	"strconv"

	"github.com/haritsrizkall/golox/errorgolox"
	"github.com/haritsrizkall/golox/token"
)

var (
	keywords = map[string]string{
		"and":    token.AND,
		"class":  token.CLASS,
		"else":   token.ELSE,
		"false":  token.FALSE,
		"for":    token.FOR,
		"fun":    token.FUN,
		"if":     token.IF,
		"nil":    token.NIL,
		"or":     token.OR,
		"print":  token.PRINT,
		"return": token.RETURN,
		"super":  token.SUPER,
		"this":   token.THIS,
		"true":   token.TRUE,
		"var":    token.VAR,
		"while":  token.WHILE,
	}
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, token.Token{
		TokenType: token.EOF,
		Lexeme:    "",
		Literal:   nil,
		Line:      s.line,
	})
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	curr := s.source[s.current]
	s.current++
	return curr
}

func (s *Scanner) scanToken() {
	c := s.advance()
	fmt.Println(string(c))
	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	case '!':
		if s.match('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}
	case '/':
		if s.match('/') {
			if s.peek() != '\n' && s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			errorgolox.LogError(s.line, "unexpected char")
		}
	}
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	// Look for a factional part
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		// Consume "."
		s.advance()

		for s.isDigit(s.peek()) {
			s.peek()
		}
	}

	value := s.source[s.start:s.current]
	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic("Ivalid number on scanning number")
	}
	s.addTokenWithLiteral(token.NUMBER, valueFloat)
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) string() {
	if s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		errorgolox.LogError(s.line, "Unterminated string dude.")
		return
	}

	// the closing "
	s.advance()

	value := s.source[s.start+1 : s.current]
	s.addTokenWithLiteral(token.STRING, value)
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) match(char byte) bool {
	if !s.isAtEnd() {
		return false
	}
	next := s.current + 1
	if s.source[next] != char {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) addToken(tokenType string) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType string, literal interface{}) {
	text := s.source[s.start : s.current+1]
	token := token.Token{
		TokenType: tokenType,
		Literal:   literal,
		Lexeme:    text,
		Line:      s.line,
	}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) identifier() {
	for s.isAlphanumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start : s.current+1]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = token.IDENTIFIER
	}
	s.addToken(tokenType)
}

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c == '_')
}

func (s *Scanner) isAlphanumeric(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
}
