package search

import "trouvo/parser"

// Engine is a SearchEngine struct
type Engine struct {
	index   *map[int][]int
	vocDict *map[string]int
	docDict *map[int]*parser.Document
}

// BoolQueryGroup is the struct for operating boolean queries
type BoolQueryGroup struct {
	q        []string
	operator string
	result   []int
}

// NewSearchEngine creates a new SearchEngine
func NewSearchEngine(index *map[int][]int, vocDict *map[string]int, docDict *map[int]*parser.Document) *Engine {
	return &Engine{index, vocDict, docDict}
}

func newBoolQueryGroup(q []string, operator string, result []int) BoolQueryGroup {
	return BoolQueryGroup{q, operator, result}
}
