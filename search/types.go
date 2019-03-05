package search

import "trouvo/parser"

// Engine is a SearchEngine struct
type Engine struct {
	index       *map[int]map[int]float64  // tokenID => docID => termFrequecy
	vocDict     *map[string]int           // 'cat' => tokenID
	idfDict     *map[int]float64          // tokenID => inverseDocFrequency
	docDict     *map[int]*parser.Document // docID => Document
	docNormDict *map[int]float64          // docID => docNorm
}

// BoolQueryGroup is the struct for operating boolean queries
type BoolQueryGroup struct {
	q        []string
	operator string
	result   []int
}

// NewSearchEngine creates a new SearchEngine
func NewSearchEngine(
	index *map[int]map[int]float64,
	vocDict *map[string]int,
	idfDict *map[int]float64,
	docDict *map[int]*parser.Document,
	docNormDict *map[int]float64,
) *Engine {
	return &Engine{index, vocDict, idfDict, docDict, docNormDict}
}

func newBoolQueryGroup(q []string, operator string, result []int) BoolQueryGroup {
	return BoolQueryGroup{q, operator, result}
}
