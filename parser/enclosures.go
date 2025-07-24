package parser

import "parsing_studies/printers"

type ElementParser func() error

type EnclosedStructureConfig struct {
	OpenMarker		string
	CloseMarker		string
	Separator		string
	ElementFunc		ElementParser
	AllowEmpty		bool
	CaseSensitive	bool
	TrimWhitespace	bool
	MaxElements		int
	MinElements		int
	CustomValidator	func(elementCount int) error
}

func (p *Parser) ParseEnclosedStructure(config EnclosedStructureConfig) error {
	// Check for opening marker
	if !p.MatchMarker(config.OpenMarker, config.CaseSensitive) {
		return printers.NewErrorf("expected '%s' at position %d", config.OpenMarker, p.pos)
	}
	p.AdvanceBy(len(config.OpenMarker))

	if config.TrimWhitespace {
		p.SkipWhitespace()
	}

	elementCount := 0

	// Check for immediate closing (empty structure)
	if p.MatchMarker(config.CloseMarker, config.CaseSensitive) {
		if config.AllowEmpty && config.MinElements == 0 {
			p.AdvanceBy(len(config.CloseMarker))
			return nil
		}
		if config.MinElements > 0 {
			return printers.NewErrorf("minimum %d elements required at position %d",
				config.MinElements, p.pos)
		}
		return printers.NewErrorf("empty structure not allowed at position %d", p.pos)
	}

	// Parse first element
	if err := config.ElementFunc(); err != nil {
		return err
	}
	elementCount++

	// Parse remaining elements
	for {
		if config.TrimWhitespace {
			p.SkipWhitespace()
		}

		// Check for closing marker
		if p.MatchMarker(config.CloseMarker, config.CaseSensitive) {
			// Validate element count
			if config.MinElements > 0 && elementCount < config.MinElements {
				return printers.NewErrorf("minimum %d elements required, got %d at position %d",
					config.MinElements, elementCount, p.pos)
			}

			// Custom validation
			if config.CustomValidator != nil {
				if err := config.CustomValidator(elementCount); err != nil {
					return err
				}
			}

			p.AdvanceBy(len(config.CloseMarker))
			return nil
		}

		// Check max elements limit
		if config.MaxElements > 0 && elementCount >= config.MaxElements {
			return printers.NewErrorf("maximum %d elements allowed at position %d",
				config.MaxElements, p.pos)
		}

		// Handle separator (if required)
		if config.Separator != "" {
			if !p.MatchMarker(config.Separator, config.CaseSensitive) {
				return printers.NewErrorf("expected '%s' or '%s' at position %d",
					config.Separator, config.CloseMarker, p.pos)
			}
			p.AdvanceBy(len(config.Separator))

			if config.TrimWhitespace {
				p.SkipWhitespace()
			}
		}

		// Parse next element
		if err := config.ElementFunc(); err != nil {
			return err
		}
		elementCount++
	}
}