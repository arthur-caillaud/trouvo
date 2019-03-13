package display

import "trouvo/parser"

// Display struct is used to display documents
type Display struct {
	docDict *map[int]*parser.Document
}

// New creates a new indexer
func New(docDict *map[int]*parser.Document) *Display {
	return &Display{docDict}
}
