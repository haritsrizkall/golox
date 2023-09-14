package scanner

import "github.com/haritsrizkall/golox/token"

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

func (s *Scanner) isAtEnd() int {
   return 0
}

func (s *Scanner) scanToken() {
    c := s.source[s.current]
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
    }
}

func (s *Scanner) addToken(tokenType string) {
    s.appendToken(tokenType, nil)
}

func (s *Scanner) appendToken (tokenType string, literal *string) {
    text := s.source[s.start:s.current+1]
    token := token.Token{
        TokenType: tokenType,
        Literal: *literal,
        Lexeme: text,
        Line: s.line,
    }
    s.tokens = append(s.tokens, token) 
}