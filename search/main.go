package search

import (
	// "fmt"
	"gogole/parser"
	"regexp"
)

// Engine is a SearchEngine struct
type Engine struct {
	index   *map[int][]int
	vocDict *map[string]int
	docDict *map[int]*parser.Document
}

// NewSearchEngine creates a new SearchEngine
func NewSearchEngine(index *map[int][]int, vocDict *map[string]int, docDict *map[int]*parser.Document) *Engine {
	return &Engine{index, vocDict, docDict}
}

func parseBoolOr(q string) []string {
	re := regexp.MustCompile(" || ")
	return re.Split(q, -1)
}

// BoolSearch executes a boolean query string to find documents
func (engine *Engine) BoolSearch(q string) (foundDocs []*parser.Document) {
	qTokens := parseBoolOr(q)
	var foundDocIDs []int
	for _, qToken := range qTokens {
		tokenID := (*engine.vocDict)[qToken]
		foundDocIDs = append(foundDocIDs, (*engine.index)[tokenID]...)
	}
	for _, docID := range foundDocIDs {
		doc := (*engine.docDict)[docID]
		shouldAppend := true
		for _, _doc := range foundDocs {
			if doc == _doc {
				shouldAppend = false
				break
			}
		}
		if shouldAppend {
			foundDocs = append(foundDocs, doc)
		}
	}
	return
}
