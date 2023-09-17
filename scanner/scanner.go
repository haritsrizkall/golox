package scanner

import (
	"strconv"

	"github.com/haritsrizkall/golox/errorgolox"
	"github.com/haritsrizkall/golox/token"
)

type Scanner struct {
    source string
    tokens []token.Token 
    start int
    current int
    line int
}

func (s *Scanner) scanTokens() []token.Token {
    return s.tokens
}

func (s *Scanner) isAtEnd() bool {
    return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
    s.current++
    return s.source[s.current]
}

func (s *Scanner) scanToken() {
    c := s.advance() 
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
        }else {
            s.addToken(token.BANG)
        }
    case '=':
        if s.match('=') {
            s.addToken(token.EQUAL_EQUAL)
        }else {
            s.addToken(token.EQUAL)
        }
    case '<':
        if s.match('=') {
            s.addToken(token.LESS_EQUAL)
        }else {
            s.addToken(token.LESS)
        }
    case '>':
        if s.match('=') {
            s.addToken(token.GREATER_EQUAL)
        }else {
            s.addToken(token.GREATER)
        }
    case '/': 
        if s.match('/') {
            if s.peak() != '\n' && s.isAtEnd() {
               s.advance() 
            }
        }else {
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
        }

        errorgolox.LogError(s.line, "Unexpected char dude")
    }
}

func (s *Scanner) number() {
    for s.isDigit(s.peak()) {
        s.advance()
    }

    // Look for a factional part
    if s.peak() == '.' && s.isDigit(s.peakNext()) {
        // Consume "."
        s.advance()

        for s.isDigit(s.peak()) {
            s.peak()
        }
    }

    value := s.source[s.start : s.current]
    valueFloat, err := strconv.ParseFloat(value, 64)
    if err != nil {
        panic("Ivalid number on scanning number")
    }
    s.addTokenWithLiteral(token.NUMBER, valueFloat)
}

func (s *Scanner) peakNext() byte {
    if s.current + 1 >= len(s.source) {
        return 0 
    }
    return s.source[s.current + 1] 
}

func (s *Scanner) isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}

func (s *Scanner) string() {
    if s.peak() != '"' && !s.isAtEnd() {
        if s.peak() == '\n' {
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

    value := s.source[s.start + 1: s.current - 1]  
    s.addTokenWithLiteral(token.STRING, value)
}

func (s *Scanner) peak() byte {
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

func (s *Scanner) addTokenWithLiteral (tokenType string, literal interface{}) {
    text := s.source[s.start:s.current]
    token := token.Token{
        TokenType: tokenType,
        Literal: literal,
        Lexeme: text,
        Line: s.line,
    }
    s.tokens = append(s.tokens, token) 
}
