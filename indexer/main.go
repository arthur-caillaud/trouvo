package indexer

import (
	// "fmt"
	"gogole/parser"
)

// Indexer struct
type Indexer struct {
	col     *parser.Collection
	index   map[int][]int
	vocDict map[string]int
	docDict map[int]*parser.Document
}

// New creates a new indexer
func New(col *parser.Collection) *Indexer {
	index := make(map[int][]int)
	vocDict := make(map[string]int)
	docDict := make(map[int]*parser.Document)
	return &Indexer{col, index, vocDict, docDict}
}

func (indexer *Indexer) buildVocDict() map[string]int {
	voc := indexer.col.GetVocabulary()
	for k, word := range voc {
		indexer.vocDict[word] = k
	}
	return indexer.vocDict
}

func (indexer *Indexer) buildDocDict() map[int]*parser.Document {
	docs := indexer.col.GetDocs()
	for k, doc := range docs {
		indexer.docDict[k] = doc
	}
	return indexer.docDict
}

// Build build the index of the indexer
func (indexer *Indexer) Build() {
	docs := indexer.buildDocDict()
	voc := indexer.buildVocDict()
	for docID, doc := range docs {
		for _, token := range doc.GetFilteredTokens() {
			tokenID := voc[token]
			postingList := indexer.index[tokenID]
			postingList = append(postingList, docID)
			indexer.index[tokenID] = postingList
		}
	}
}

// GetIndex returns the pointer of the index
func (indexer Indexer) GetIndex() *map[int][]int {
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
