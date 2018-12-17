package search

import (
	// "fmt"
	"regexp"
	"trouvo/parser"
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
	re := regexp.MustCompile("||")
	return re.Split(q, -1)
}

func parseBoolAnd(q string) []string {
	re := regexp.MustCompile("&&")
	return re.Split(q, -1)
}

func parseBoolNot(q string) []string {
	re := regexp.MustCompile("!")
	token := re.Split(q, -1)[1]
	return []string{token}
}

// BoolSearch executes a boolean query string to find documents
func (engine *Engine) BoolSearch(q string) (foundDocs []*parser.Document) {
	foundDocIDs := engine.boolSearchOr(q)
	foundDocs = engine.foundDuplicates(foundDocIDs)
	return
}

// boolSearchOr executes a boolean OR query string to find documents
func (engine *Engine) boolSearchOr(q string) (foundDocIDs []int) {
	qTokens := parseBoolOr(q)
	for _, qToken := range qTokens {
		tokenID := (*engine.vocDict)[qToken]
		foundDocIDs = append(foundDocIDs, (*engine.index)[tokenID]...)
	}
	return
}

// boolSearchAnd executes a boolean AND query string to find documents
func (engine *Engine) boolSearchAnd(q string) (foundDocIDs []int) {
	qTokens := parseBoolAnd(q)
	for i, qToken := range qTokens {
		tokenID := (*engine.vocDict)[qToken]
		if i == 0 {
			foundDocIDs = append(foundDocIDs, (*engine.index)[tokenID]...)
		} else {
			foundDocIDs = intersect(foundDocIDs, (*engine.index)[tokenID])
		}
	}
	return
}

// boolSearchNot executes a boolean NOT query string to find documents
func (engine *Engine) boolSearchNot(q string) (foundDocIDs []int) {
	qToken := parseBoolNot(q)[0]
	tokenID := (*engine.vocDict)[qToken]
	var allDocIDs []int
	for docID := range *engine.docDict {
		allDocIDs = append(allDocIDs, docID)
	}
	for _, docID := range allDocIDs {
		for _, _docID := range (*engine.index)[tokenID] {
			if docID == _docID {
				break
			}
		}
		foundDocIDs = append(foundDocIDs, docID)
	}
	return
}

func (engine *Engine) foundDuplicates(foundDocIDs []int) (foundDocs []*parser.Document) {
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

func intersect(a []int, b []int) (inter []int) {
	for _, elA := range a {
		for _, elB := range b {
			if elA == elB {
				inter = append(inter, elA)
				break
			}
		}
	}
	return
}
