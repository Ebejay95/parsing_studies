package parser

import "unicode"
import "parsing_studies/printers"

type Parser struct {
	input string
	pos int
}

func (p *Parser) Pos() int {
	return p.pos
}

func (p *Parser) Input() string {
	return p.input
}

func NewParser(input string) *Parser {
	return &Parser{input: input, pos: 0}
}

func (p *Parser) Current() byte {
	if p.pos >= len(p.input) {
		return 0
	}
	return p.input[p.pos]
}

func (p *Parser) Advance() {
	if p.pos < len(p.input) {
		p.pos++
	}
	if p.pos < len(p.input) {
		printers.Log("next: " + string(p.input[p.pos]))
	} else {
		printers.Log("next: <EOF>")
	}
}

func (p *Parser) SkipWhitespace() {
	for p.pos < len(p.input) && unicode.IsSpace(rune(p.Current())) {
		p.Advance()
	}
}

func (p *Parser) AtEnd() bool {
	return p.pos >= len(p.input)
}

func (p *Parser) MatchMarker(marker string, caseSensitive bool) bool {
	if p.pos+len(marker) > len(p.input) {
		return false
	}

	candidate := p.input[p.pos:p.pos+len(marker)]

	if caseSensitive {
		return candidate == marker
	}
	return strings.EqualFold(candidate, marker)
}

func (p *Parser) AdvanceBy(n int) {
	p.pos += n
	if p.pos >= len(p.input) {
		p.pos = len(p.input)
	}
}

// func (p *Parser) ParseString() (string, error) {
// 	if p.pos >= len(p.input) {
// 		return "", fmt.Errorf("unexpected end of input")
// 	}

// 	quote := p.Current()
// 	if quote != '"' && quote != '\'' {
// 		return "", fmt.Errorf("expected quote at position %d", p.pos)
// 	}

// 	p.Advance()
// 	var result strings.Builder

// 	for p.pos < len(p.input) {
// 		char := p.Current()

// 		if char == quote {
// 			p.Advance()
// 			return result.String(), nil
// 		}

// 		if char == '\\' {
// 			p.Advance()
// 			if p.pos >= len(p.input) {
// 				return "", fmt.Errorf("unexpected end of input after escape")
// 			}

// 			escaped := p.Current()
// 			switch escaped {
// 			case 'n':
// 				result.WriteByte('\n')
// 			case 't':
// 				result.WriteByte('\t')
// 			case 'r':
// 				result.WriteByte('\r')
// 			case '\\':
// 				result.WriteByte('\\')
// 			case '"':
// 				result.WriteByte('"')
// 			case '\'':
// 				result.WriteByte('\'')
// 			default:
// 				result.WriteByte('\\')
// 				result.WriteByte(escaped)
// 			}
// 		} else {
// 			result.WriteByte(char)
// 		}

// 		p.Advance()
// 	}

// 	return "", fmt.Errorf("unterminated string starting at position %d", p.pos)
// }

// func (p *Parser) ParseNestedQuotes() (string, error) {
// 	if p.pos >= len(p.input) {
// 		return "", fmt.Errorf("unexpected end of input")
// 	}

// 	outerQuote := p.Current()
// 	if outerQuote != '"' && outerQuote != '\'' {
// 		return "", fmt.Errorf("expected quote at position %d", p.pos)
// 	}

// 	p.Advance()
// 	var result strings.Builder

// 	for p.pos < len(p.input) {
// 		char := p.Current()

// 		if char == outerQuote {
// 			p.Advance()
// 			return result.String(), nil
// 		}
// 		if char == '"' || char == '\'' {
// 			if char != outerQuote {
// 				result.WriteByte(char)
// 			} else {
// 				p.Advance()
// 				return result.String(), nil
// 			}
// 		} else if char == '\\' {
// 			p.Advance()
// 			if p.pos >= len(p.input) {
// 				return "", fmt.Errorf("unexpected end of input after escape")
// 			}
// 			escaped := p.Current()
// 			result.WriteByte(escaped)
// 		} else {
// 			result.WriteByte(char)
// 		}

// 		p.Advance()
// 	}

// 	return "", fmt.Errorf("unterminated string")
// }