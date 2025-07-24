package json_parser

import "parsing_studies/parser"
import "parsing_studies/printers"

type JSONParser struct {
	*parser.Parser
}

func NewJSONParser(input string) *JSONParser {
	return &JSONParser{
		Parser: parser.NewParser(input),
	}
}

func (jp *JSONParser) parseJSONString() error {
	if jp.Current() != '"' {
		return printers.NewErrorf("expected \" at position %d", jp.Pos())
	}

	jp.Advance()

	for jp.Pos() < len(jp.Input()) {
		if jp.Current() == '"' {
			jp.Advance()
			return nil
		}
		if jp.Current() == '\\' {
			jp.Advance()
			if jp.Pos() >= len(jp.Input()) {
				return printers.NewErrorf("unexpected end of input after escape character")
			}
		}
		jp.Advance()
	}
	return printers.NewErrorf("unclosed string starting at position %d", jp.Pos())
}

func (jp *JSONParser) parseKeyValue() error {
	jp.SkipWhitespace()
	err := jp.parseJSONString()
	if err != nil {
		return printers.NewErrorf("key must be a string")
	}
	jp.SkipWhitespace()
	if jp.Current() != ':' {
		return printers.NewError("expected an : but got " + string(jp.Current()))
	}
	jp.Advance()
	jp.SkipWhitespace()
	return jp.parseJSONElement()
}

func (jp JSONParser) parseLiteral(literal string) error {
	for i := 0 ; i < len(literal) ; i++ {
		if jp.Pos() >= len(jp.Input()) && literal[i] != jp.Current() {
			return printers.NewErrorf("unexpected end of input after escape character")
		}
		jp.Advance()
	}
	return nil
}

func (jp *JSONParser) parseJSONElement() error {
	printers.Log("parseJSONElement")
	jp.SkipWhitespace()

	switch jp.Current() {
		case '"':
			printers.Log("parseJSONString")
			return jp.parseJSONString()
		case '{':
			printers.Log("ParseEnclosedStructure OBJECT")
			return jp.Parser.ParseEnclosedStructure(parser.EnclosedStructureConfig{
				OpenChar:    '{',
				CloseChar:   '}',
				Separator:   ',',
				ElementFunc: func() error { return jp.parseKeyValue() },
				AllowEmpty:  true,
			})
		case '[':
			printers.Log("ParseEnclosedStructure ARRAY")
			return jp.Parser.ParseEnclosedStructure(parser.EnclosedStructureConfig{
				OpenChar:    '[',
				CloseChar:   ']',
				Separator:   ',',
				ElementFunc: func() error { return jp.parseJSONElement() },
				AllowEmpty:  true,
			})
		case 't':
			printers.Log("parseLiteral true")
			return jp.parseLiteral("true")
		case 'f':
			printers.Log("parseLiteral false")
			return jp.parseLiteral("false")
		case 'n':
			printers.Log("parseLiteral null")
			return jp.parseLiteral("null")
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			printers.Log("ParseNumber")
			return jp.ParseNumber()
		default:
			return printers.NewErrorf("unexpected character '%c'", jp.Current())
	}
}

func IsValidJSON(input string) bool {
	jp := NewJSONParser(input)

	err := jp.parseJSONElement()
	if err != nil {
		printers.Errorf("%s", err)
		return false
	}
	jp.SkipWhitespace()
	if !jp.AtEnd() {
		printers.Error("Unexpected characters after valid JSON")
		return false
	}
	return true
}
