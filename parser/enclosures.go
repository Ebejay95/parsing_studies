package parser

import "parsing_studies/printers"

type ElementParser func() error

type EnclosedStructureConfig struct {
	OpenChar	byte
	CloseChar	byte
	Separator	byte
	ElementFunc	ElementParser
	AllowEmpty	bool
}

func (p *Parser) ParseEnclosedStructure(config EnclosedStructureConfig) error {
	if p.Current() != config.OpenChar {
		return printers.NewErrorf("expected '%c' at position %d", config.OpenChar, p.pos)
	}
	p.Advance()
	p.SkipWhitespace()

	if p.Current() == config.CloseChar {
		if config.AllowEmpty {
			p.Advance()
			return nil
		}
		return printers.NewErrorf("empty structure not allowed at position %d", p.pos)
	}

	if err := config.ElementFunc(); err != nil {
		return err
	}

	for {
		p.SkipWhitespace()

		if p.Current() == config.CloseChar {
			p.Advance()
			return nil
		}

		if p.Current() != config.Separator {
			return printers.NewErrorf("expected '%c' or '%c' at position %d", config.Separator, config.CloseChar, p.pos)
		}
		p.Advance()
		p.SkipWhitespace()

		if err := config.ElementFunc(); err != nil {
			return err
		}
	}
}