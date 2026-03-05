package parser

import "io"

// Result represents the outcome of parsing a page.
type Result struct {
	Title string
	Links []string
}

// Parser defines the interface for extracting information from HTML.
type Parser interface {
	Parse(body io.Reader) (*Result, error)
}
