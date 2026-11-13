package parser

import "github.com/osteele/liquid/expressions"

// A Config holds configuration information for parsing and rendering.
type Config struct {
	expressions.Config
	Grammar Grammar
	Delims  []string
	// TagsDisabled makes the parser treat tags as text.
	TagsDisabled bool
}

// NewConfig creates a parser Config.
func NewConfig(g Grammar) Config {
	return Config{Grammar: g}
}
