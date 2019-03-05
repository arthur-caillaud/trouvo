package search

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

// Run the search engine that reads the console input to process the query
func (engine *Engine) Run() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Type query :")
		text, _ := reader.ReadString('\n')
		start := time.Now()
		res := engine.VectSearch(text)
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Println(len(res), "results found in", elapsed)
		fmt.Println(res)
		fmt.Println("----")
	}
}

// BoolSearch runs a boolean query with the SearchEngine
func (engine *Engine) BoolSearch(q string) (res []int) {
	q = strings.TrimSpace(q)
	firstBoolQueryGroup := newBoolQueryGroup([]string{q}, "RET", res)
	res = (*engine).recursiveBoolSearch(firstBoolQueryGroup).result
	return res
}

// VectSearch runs a vectorial query with the SearchEngine and cos measure
func (engine *Engine) VectSearch(q string) (res []int) {
	qWords := splitWords(q)
	index := *engine.index
	vocDict := *engine.vocDict
	idfDict := *engine.idfDict
	docNormDict := *engine.docNormDict
	var qNormFactor float64
	s := make(map[int]float64)

	for _, qWord := range qWords {
		termOccurence := 0
		for _, _qWord := range qWords {
			if _qWord == qWord {
				termOccurence++
			}
		}
		termFrequency := float64(termOccurence) / float64(len(qWords))

		if tokenID, ok := vocDict[qWord]; ok {
			inverseDocFrequency := idfDict[tokenID]
			qTermWeight := termFrequency * inverseDocFrequency
			qNormFactor += qTermWeight * qTermWeight

			for docID, termFrequency := range index[tokenID] {
				docTermWeight := docNormDict[docID] * termFrequency * inverseDocFrequency
				s[docID] = docTermWeight * docTermWeight
			}
		}
	}

	for docID, docScore := range s {
		if docScore != 0 {
			s[docID] = docScore / (math.Sqrt(qNormFactor) * math.Sqrt(docNormDict[docID]))
			res = append(res, docID)
		}
	}

	sort.Slice(res, makeSortDocClosure(s))

	return
}

func (engine *Engine) recursiveBoolSearch(b BoolQueryGroup) BoolQueryGroup {
	subResults := [][]int{}
	for _, q := range b.q {
		if isOr(q) {
			subResults = append(subResults, (*engine).getQuerySubResults(q, "OR"))
		} else if isAnd(q) {
			subResults = append(subResults, (*engine).getQuerySubResults(q, "AND"))
		} else if isNot(q) {
			subResults = append(subResults, (*engine).getQuerySubResults(q, "NOT"))
		} else {
			if vocID, ok := (*engine.vocDict)[q]; ok {
				postings := (*engine.index)[vocID]
				subResult := []int{}
				for docID := range postings {
					subResult = append(subResult, docID)
				}
				subResults = append(subResults, subResult)
			} else {
				subResults = append(subResults, []int{})
			}
		}
	}
	res := engine.processSubResults(subResults, b.operator)
	return newBoolQueryGroup(b.q, b.operator, res)
}

func (engine *Engine) getQuerySubResults(q string, op string) []int {
	subQueries := parse(q, op)
	subQueryGroup := newBoolQueryGroup(subQueries, op, []int{})
	res := (*engine).recursiveBoolSearch(subQueryGroup)
	return res.result
}

func (engine *Engine) processSubResults(subResults [][]int, op string) (res []int) {
	switch op {
	case "AND":
		res = intersect(subResults...)
	case "OR":
		res = union(subResults...)
	case "NOT":
		allDocs := []int{}
		for docID := range *engine.docDict {
			allDocs = append(allDocs, docID)
		}
		res = subtract(allDocs, subResults...)
	case "RET":
		for _, subResult := range subResults {
			res = append(res, subResult...)
		}
	}
	return
}
