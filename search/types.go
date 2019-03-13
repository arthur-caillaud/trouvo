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

// SuperEngine enables to perform requests on multiple collection simultaneously
type SuperEngine struct {
	engines []*Engine
}

// BoolQueryGroup is the struct for operating boolean queries
type BoolQueryGroup struct {
	q        []string
	operator string
	result   []int
}

// Result is the struct returned by our engines
type Result struct {
	docID int
	score float64
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

// NewSuperEngine creates a new SuperEngine
func NewSuperEngine(engines []*Engine) *SuperEngine {
	return &SuperEngine{engines}
}

func newBoolQueryGroup(q []string, operator string, result []int) BoolQueryGroup {
	return BoolQueryGroup{q, operator, result}
}

func newResult(docID int, score float64) *Result {
	return &Result{docID, score}
}

func (res *Result) GetDocID() int {
	return res.docID
}

func (res *Result) GetScore() float64 {
	return res.score
}
