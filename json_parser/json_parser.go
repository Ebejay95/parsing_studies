package json_parser

import "json_validator/printers"

type Parser struct {
	input string
	index int
}

func NewParser(input string) *Parser {
	return &Parser{input: input, index: 0}
}

func (p *Parser) current() byte {
	if p.index >= len(p.input) {
		return 0
	}
	printers.Log(string(p.input[p.index]))
	return p.input[p.index]
}

func (p *Parser) next() {
	p.index++
	printers.Log("next: " + string(p.input[p.index]))
}

// parser  und contrictor ddefinineren
// current, step, skeip whitespace
// parse String
// parse Number
// parse Literal
// parse Array
// parse Object
// parse KeyValue
// parse Element
// "main" function


func ParseJSON(input string) bool {
	printers.Log(input)
	parser := NewParser(input)
	parser.current()
	parser.next()
	parser.next()
	parser.next()
	return true
}
