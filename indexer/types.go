package indexer

import (
	"trouvo/parser"
)

// Indexer struct
type Indexer struct {
	col     *parser.Collection
	index   map[int]map[int]float64 // tokenID => docID => termFrequecy
	vocDict map[string]int          // 'cat' => tokenID
	idfDict map[int]float64         // tokenID => inverseDocFrequency
	docDict map[int]*parser.Document
}

// New creates a new indexer
func New(col *parser.Collection) *Indexer {
	index := make(map[int]map[int]float64)
	idfDict := make(map[int]float64)
	vocDict := make(map[string]int)
	docDict := make(map[int]*parser.Document)
	return &Indexer{col, index, vocDict, idfDict, docDict}
}

// GetIndex returns the pointer of the index
func (indexer Indexer) GetIndex() *map[int]map[int]float64 {
	return &indexer.index
}

// GetVocDict returns the pointer of the vocDict
func (indexer Indexer) GetVocDict() *map[string]int {
	return &indexer.vocDict
}

// GetDocDict returns the pointer of the docDict
func (indexer Indexer) GetDocDict() *map[int]*parser.Document {
	return &indexer.docDict
}

// GetIdfDict returns the pointer of the idfDict
func (indexer Indexer) GetIdfDict() *map[int]float64 {
	return &indexer.idfDict
}
