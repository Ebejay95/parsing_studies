package json_parser

import (
	"fmt"
	"unicode"
)

type Parser struct {
	input string
	pos   int
}

func NewParser(input string) *Parser {
	return &Parser{input: input, pos: 0}
}

// Hilfsfunktionen für Character-Navigation
func (p *Parser) current() byte {
	if p.pos >= len(p.input) {
		return 0
	}
	return p.input[p.pos]
}

func (p *Parser) advance() {
	if p.pos < len(p.input) {
		p.pos++
	}
}

func (p *Parser) skipWhitespace() {
	for p.pos < len(p.input) && unicode.IsSpace(rune(p.current())) {
		p.advance()
	}
}

// String-Parsing
func (p *Parser) parseString() error {
	if p.current() != '"' {
		return fmt.Errorf("expected '\"' at position %d", p.pos)
	}

	p.advance() // Skip opening quote

	for p.pos < len(p.input) {
		ch := p.current()

		if ch == '"' {
			p.advance() // Skip closing quote
			return nil
		}

		if ch == '\\' {
			p.advance() // Skip escape character
			if p.pos >= len(p.input) {
				return fmt.Errorf("unexpected end of input after escape character")
			}
			// Hier könntest du spezifische Escape-Sequenzen validieren
		}

		p.advance()
	}

	return fmt.Errorf("unclosed string starting at position %d", p.pos)
}

// Number-Parsing (vereinfacht)
func (p *Parser) parseNumber() error {
	start := p.pos

	// Optionales Minus
	if p.current() == '-' {
		p.advance()
	}

	// Mindestens eine Ziffer
	if !unicode.IsDigit(rune(p.current())) {
		return fmt.Errorf("invalid number at position %d", start)
	}

	// Ziffern parsen
	for p.pos < len(p.input) && unicode.IsDigit(rune(p.current())) {
		p.advance()
	}

	// Optionaler Dezimalteil
	if p.current() == '.' {
		p.advance()
		if !unicode.IsDigit(rune(p.current())) {
			return fmt.Errorf("invalid number: missing digits after decimal point at position %d", p.pos)
		}
		for p.pos < len(p.input) && unicode.IsDigit(rune(p.current())) {
			p.advance()
		}
	}

	// Optionaler Exponent
	if p.current() == 'e' || p.current() == 'E' {
		p.advance()
		if p.current() == '+' || p.current() == '-' {
			p.advance()
		}
		if !unicode.IsDigit(rune(p.current())) {
			return fmt.Errorf("invalid number: missing digits in exponent at position %d", p.pos)
		}
		for p.pos < len(p.input) && unicode.IsDigit(rune(p.current())) {
			p.advance()
		}
	}

	return nil
}

// Literale (true, false, null)
func (p *Parser) parseLiteral(literal string) error {
	for i := 0; i < len(literal); i++ {
		if p.pos >= len(p.input) || p.current() != literal[i] {
			return fmt.Errorf("invalid literal at position %d", p.pos-i)
		}
		p.advance()
	}
	return nil
}

// Array-Parsing (rekursiv)
func (p *Parser) parseArray() error {
	if p.current() != '[' {
		return fmt.Errorf("expected '[' at position %d", p.pos)
	}

	p.advance() // Skip '['
	p.skipWhitespace()

	// Leeres Array
	if p.current() == ']' {
		p.advance()
		return nil
	}

	// Erstes Element
	if err := p.parseJSONElement(); err != nil {
		return err
	}

	// Weitere Elemente
	for {
		p.skipWhitespace()

		if p.current() == ']' {
			p.advance()
			return nil
		}

		if p.current() != ',' {
			return fmt.Errorf("expected ',' or ']' at position %d", p.pos)
		}

		p.advance() // Skip ','
		p.skipWhitespace()

		if err := p.parseJSONElement(); err != nil {
			return err
		}
	}
}

// Object-Parsing (rekursiv)
func (p *Parser) parseObject() error {
	if p.current() != '{' {
		return fmt.Errorf("expected '{' at position %d", p.pos)
	}

	p.advance() // Skip '{'
	p.skipWhitespace()

	// Leeres Object
	if p.current() == '}' {
		p.advance()
		return nil
	}

	// Erstes Key-Value Paar
	if err := p.parseKeyValue(); err != nil {
		return err
	}

	// Weitere Key-Value Paare
	for {
		p.skipWhitespace()

		if p.current() == '}' {
			p.advance()
			return nil
		}

		if p.current() != ',' {
			return fmt.Errorf("expected ',' or '}' at position %d", p.pos)
		}

		p.advance() // Skip ','
		p.skipWhitespace()

		if err := p.parseKeyValue(); err != nil {
			return err
		}
	}
}

// Key-Value Paar parsen
func (p *Parser) parseKeyValue() error {
	// Key muss ein String sein
	if err := p.parseString(); err != nil {
		return fmt.Errorf("invalid object key: %v", err)
	}

	p.skipWhitespace()

	// Doppelpunkt erwarten
	if p.current() != ':' {
		return fmt.Errorf("expected ':' after object key at position %d", p.pos)
	}

	p.advance() // Skip ':'
	p.skipWhitespace()

	// Value parsen
	return p.parseJSONElement()
}

// Haupt-Parsing-Funktion (rekursiv)
func (p *Parser) parseJSONElement() error {
	fmt.Print("parseJSONElement ")
	p.skipWhitespace()

	switch p.current() {
	case '"':
		fmt.PrintLn(" is \" ")
		return p.parseString()
	case '{':
		fmt.PrintLn(" is { ")
		return p.parseObject()
	case '[':
		fmt.PrintLn(" is [ ")
		return p.parseArray()
	case 't':
		fmt.PrintLn(" is t ")
		return p.parseLiteral("true")
	case 'f':
		fmt.PrintLn(" is f ")
		return p.parseLiteral("false")
	case 'n':
		fmt.PrintLn(" is n ")
		return p.parseLiteral("null")
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		fmt.PrintLn(" is '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' ")
		return p.parseNumber()
	default:
		return fmt.Errorf("unexpected character '%c' at position %d", p.current(), p.pos)
	}
}

// Hauptfunktion für JSON-Validierung
func IsValidJSON(jsonStr string) bool {
	parser := NewParser(jsonStr)

	// Haupt-JSON-Wert parsen
	if err := parser.parseJSONElement(); err != nil {
		fmt.Printf("JSON parsing error: %v\n", err)
		return false
	}

	// Überprüfen, ob noch unverarbeitete Zeichen vorhanden sind
	parser.skipWhitespace()
	if parser.pos < len(parser.input) {
		fmt.Printf("Unexpected characters after valid JSON at position %d\n", parser.pos)
		return false
	}

	return true
}