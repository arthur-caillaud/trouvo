package search

import (
	// "fmt"
	"strings"
)

// BoolSearch operates a boolean query with the SearchEngine
func (engine *Engine) BoolSearch(q string) (res []int) {
	q = strings.TrimSpace(q)
	firstBoolQueryGroup := newBoolQueryGroup([]string{q}, "RET", res)
	res = (*engine).recursiveBoolSearch(firstBoolQueryGroup).result
	return res
}

func (engine *Engine) recursiveBoolSearch(b BoolQueryGroup) BoolQueryGroup {
	subResults := [][]int{}
	for _, q := range b.q {
		if isOr(q) {
			subQueries := parseOr(q)
			subQueryGroup := newBoolQueryGroup(subQueries, "OR", []int{})
			res := (*engine).recursiveBoolSearch(subQueryGroup)
			subResults = append(subResults, res.result)
		} else if isAnd(q) {
			subQueries := parseAnd(q)
			subQueryGroup := newBoolQueryGroup(subQueries, "AND", []int{})
			res := (*engine).recursiveBoolSearch(subQueryGroup)
			subResults = append(subResults, res.result)
		} else if isNot(q) {
			subQueries := parseNot(q)
			subQueryGroup := newBoolQueryGroup(subQueries, "NOT", []int{})
			res := (*engine).recursiveBoolSearch(subQueryGroup)
			subResults = append(subResults, res.result)
		} else {
			res := (*engine.index)[(*engine.vocDict)[q]]
			subResults = append(subResults, res)
		}
	}
	res := []int{}
	switch b.operator {
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
	return newBoolQueryGroup(b.q, b.operator, res)
}
