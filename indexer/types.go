package indexer

import (
	"math"
	"trouvo/parser"
)

// Indexer struct
type Indexer struct {
	col         *parser.Collection       // Collection
	index       map[int]map[int]float64  // tokenID => docID => termFrequecy
	vocDict     map[string]int           // 'cat' => tokenID
	idfDict     map[int]float64          // tokenID => inverseDocFrequency
	docDict     map[int]*parser.Document // docID => Document
	docNormDict map[int]float64          // docID => docNorm
}

// New creates a new indexer
func New(col *parser.Collection) *Indexer {
	index := make(map[int]map[int]float64)
	idfDict := make(map[int]float64)
	vocDict := make(map[string]int)
	docDict := make(map[int]*parser.Document)
	docNormDict := make(map[int]float64)
	return &Indexer{col, index, vocDict, idfDict, docDict, docNormDict}
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

// GetDocNormDict returns the pointer of the docNormDict
func (indexer Indexer) GetDocNormDict() *map[int]float64 {
	return &indexer.docNormDict
}

// GetIndexSize computes the total size of the index in bytes
func (indexer Indexer) GetIndexSize() int {
	index := *indexer.GetIndex()
	size := 8 * len(index)
	for _, postings := range index {
		size += 16 * len(postings)
	}
	kb := int(math.Pow(2, 10))
	size = int(size / kb)
	return size
}
