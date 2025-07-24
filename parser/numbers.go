package parser

import "unicode"
import "parsing_studies/printers"

func (jp *Parser) ParseNumber() error {
	if jp.Current() == '-' {
		jp.Advance()
	}

	if !unicode.IsDigit(rune(jp.Current())) {
		return printers.NewError("invalid number")
	}

	for jp.Current() != 0 && unicode.IsDigit(rune(jp.Current())) {
		jp.Advance()
	}

	if jp.Current() == '.' {
		jp.Advance()
		if !unicode.IsDigit(rune(jp.Current())) {
			return printers.NewError("invalid number: missing digits after decimal point")
		}
		for jp.Current() != 0 && unicode.IsDigit(rune(jp.Current())) {
			jp.Advance()
		}
	}

	if jp.Current() == 'e' || jp.Current() == 'E' {
		jp.Advance()
		if jp.Current() == '+' || jp.Current() == '-' {
			jp.Advance()
		}
		if !unicode.IsDigit(rune(jp.Current())) {
			return printers.NewError("invalid number: missing digits in exponent")
		}
		for jp.Current() != 0 && unicode.IsDigit(rune(jp.Current())) {
			jp.Advance()
		}
	}

	return nil
}