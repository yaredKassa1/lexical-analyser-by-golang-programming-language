package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

// TokenType represents the type of a token.
type TokenType string

// Constants for token types.
const (
	Keyword    TokenType = "Keyword"
	Operator   TokenType = "Operator"
	Delimiter  TokenType = "Delimiter"
	Identifier TokenType = "Identifier"
	Literal    TokenType = "Literal"
	Error      TokenType = "Error"
)

// Token structure to store token information.
type Token struct {
	Type    TokenType
	Value   string
	Line    int
	Column  int
}

// Lexer struct to hold the lexing state.
type Lexer struct {
	input      string
	position   int
	line       int
	column     int
	tokens     []Token
	keywords   map[string]bool
	operators  map[string]bool
	delimiters map[string]bool
}

// NewLexer creates a new Lexer with default keywords, operators, and delimiters.
func NewLexer(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: 0,
		line:     1,
		column:   1,
		tokens:   []Token{},

		keywords: map[string]bool{
                   "break":       true,
                   "case":        true,
                   "chan":        true,
                   "const":       true,
                   "continue":    true,
                   "default":     true,
                   "defer":       true,
                   "else":        true,
                   "fallthrough": true,
                   "for":         true,
                   "func":        true,
                   "go":          true,
                   "if":          true,
                   "import":      true,
                   "interface":   true,
                   "map":         true,
                   "package":     true,
                   "range":       true,
                   "return":      true,
                   "select":      true,
                   "struct":      true,
                   "switch":      true,
                   "type":        true,
                   "var":         true,
                   "true":        true,
                   "false":       true,
                   "iota":        true,
		},

		operators: map[string]bool{
                    "+":   true,
                    "-":   true,
                    "*":   true,
                    "/":   true,
                    "%":   true,
                    "=":   true,
                    "==":  true,
                    "!=":  true,
                    "<":   true,
                    ">":   true,
                    "<=":  true,
                    ">=":  true,
                    "&&":  true,
                    "||":  true,
                    "!":   true,
                    "&":   true,
                    "|":   true,
                    "^":   true,
                    "<<":  true,
                    ">>":  true,
                    "+=":  true,
                    "-=":  true,
                    "*=":  true,
                    "/=":  true,
                    "%=":  true,
                    "&=":  true,
                    "|=":  true,
                    "^=":  true,
                    "&&=": true,
                    "||=": true,
                },

		delimiters: map[string]bool{
			"{": true,
			"}": true,
			"(": true,
			")": true,
			"[": true,
			"]": true,
			";": true,
			",": true,
			".": true,
                        ":": true,
                        "":  true,
		},
	}
}

// lex starts the lexical analysis
func (l *Lexer) lex() []Token {
	for l.position < len(l.input) {
		char := rune(l.input[l.position])
		switch {
		case unicode.IsSpace(char):
			l.skipWhitespace()
		case strings.HasPrefix(l.input[l.position:], "//"):
			l.skipSingleLineComment()
		case strings.HasPrefix(l.input[l.position:], "/*"):
			l.skipMultiLineComment()
		case char == '"':
			l.lexString() // Handle string literals
		case isIdentifierStart(char):
			l.lexIdentifier()
		case unicode.IsDigit(char) || char == '.':
			l.lexNumber()
		case char == '\'':
			l.lexCharacter()
		case l.isOperator(string(char)):
			l.lexOperator()
		case l.isDelimiter(string(char)):
			l.lexDelimiter()
		default:
			l.reportError(fmt.Sprintf("Invalid character: %c", char), Error, l.line, l.column)
			l.next()
		}
	}
	return l.tokens
}

// skipWhitespace skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.position < len(l.input) && unicode.IsSpace(rune(l.input[l.position])) {
		if rune(l.input[l.position]) == '\n' {
			l.line++
			l.column = 1
		} else {
			l.column++
		}
		l.position++
	}
}

// lexIdentifier lexes an identifier
func (l *Lexer) lexIdentifier() {
	start := l.position
	for l.position < len(l.input) && isIdentifierChar(rune(l.input[l.position])) {
		l.next()
	}
	value := l.input[start:l.position]
	if l.isKeyword(value) {
		l.addToken(Keyword, value, l.line, start+1)
	} else {
		l.addToken(Identifier, value, l.line, start+1)
	}
}

// lexNumber lexes a number literal
func (l *Lexer) lexNumber() {
	start := l.position
	isFloat := false
	for l.position < len(l.input) {
		if unicode.IsDigit(rune(l.input[l.position])) {
			l.next()
		} else if rune(l.input[l.position]) == '.' && !isFloat {
			isFloat = true
			l.next()
		} else {
			break
		}
	}
	value := l.input[start:l.position]
	l.addToken(Literal, value, l.line, start+1)
}

// lexCharacter lexes a character literal
func (l *Lexer) lexCharacter() {
	start := l.position
	l.next()
	if l.position < len(l.input) && rune(l.input[l.position]) != '\'' {
		l.next()
		if l.position < len(l.input) && rune(l.input[l.position]) == '\'' {
			l.next()
			value := l.input[start:l.position]
			l.addToken(Literal, value, l.line, start+1)
		} else {
			l.reportError("Invalid character literal", Error, l.line, l.column)
		}
	} else {
		l.reportError("Invalid character literal", Error, l.line, l.column)
	}
}

// lexString lexes a string literal (e.g., "Hello")
func (l *Lexer) lexString() {
	start := l.position
	l.next() // Move past the opening quote (")
	
	// Consume characters until the closing quote or EOF
	for l.position < len(l.input) && rune(l.input[l.position]) != '"' {
		l.next()
	}
	if l.position < len(l.input) && rune(l.input[l.position]) == '"' {
		l.next() // Move past the closing quote (")
		value := l.input[start:l.position] // Extract the entire string literal (including quotes)
		l.addToken(Literal, value, l.line, start+1) // Add the literal token
	} else {
		l.reportError("Unterminated string literal", Error, l.line, l.column) // Report error if closing quote is missing
	}
}

// lexOperator lexes an operator
func (l *Lexer) lexOperator() {
	start := l.position
	if l.position+1 < len(l.input) {
		doubleCharOp := string(rune(l.input[l.position])) + string(rune(l.input[l.position+1]))
		if l.isOperator(doubleCharOp) {
			l.position += 2
			l.addToken(Operator, doubleCharOp, l.line, start+1)
			return
		}
	}
	value := string(rune(l.input[start]))
	l.addToken(Operator, value, l.line, start+1)
	l.next()
}

// lexDelimiter lexes a delimiter
func (l *Lexer) lexDelimiter() {
	value := string(rune(l.input[l.position]))
	l.addToken(Delimiter, value, l.line, l.column)
	l.next()
}

// skipSingleLineComment skips a single line comment
func (l *Lexer) skipSingleLineComment() {
	for l.position < len(l.input) && rune(l.input[l.position]) != '\n' {
		l.next()
	}
	l.skipWhitespace()
}

// skipMultiLineComment skips a multi-line comment
func (l *Lexer) skipMultiLineComment() {
	l.position += 2 // Skip the opening /*
	for l.position < len(l.input)-1 {
		if rune(l.input[l.position]) == '*' && rune(l.input[l.position+1]) == '/' {
			l.position += 2 // Skip the closing */
			l.skipWhitespace()
			return
		}
		l.next()
	}
	l.reportError("Unclosed multiline comment", Error, l.line, l.column)
}

// isKeyword checks if a string is a keyword
func (l *Lexer) isKeyword(value string) bool {
	return l.keywords[value]
}

// isOperator checks if a string is an operator
func (l *Lexer) isOperator(value string) bool {
	return l.operators[value]
}

// isDelimiter checks if a string is a delimiter
func (l *Lexer) isDelimiter(value string) bool {
	return l.delimiters[value]
}

// addToken adds a token to the tokens list
func (l *Lexer) addToken(tokenType TokenType, value string, line int, column int) {
	l.tokens = append(l.tokens, Token{Type: tokenType, Value: value, Line: line, Column: column})
}

// next moves to the next character in the input
func (l *Lexer) next() {
	l.position++
	l.column++
}

// reportError reports an error
func (l *Lexer) reportError(message string, tokenType TokenType, line int, column int) {
	l.addToken(tokenType, message, line, column)
	fmt.Printf("Error at line %d, column %d : %s \n", line, column, message)
}

// isIdentifierStart checks if a character is valid as the first char of an identifier.
func isIdentifierStart(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

// isIdentifierChar checks if a character is valid for the rest of an identifier.
func isIdentifierChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// main function for testing
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run lexer.go <filename>")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	var sb strings.Builder
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		sb.WriteString(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}

	input := sb.String()

	lexer := NewLexer(input)
	tokens := lexer.lex()

	fmt.Printf("%-15s %-20s %-12s %-12s\n", "Token Type", "Token Value", "Line Number", "Column Number")
       for _, token := range tokens {
         fmt.Printf("%-15s %-20s %-12d %-12d\n", token.Type, token.Value, token.Line, token.Column)
     }
}