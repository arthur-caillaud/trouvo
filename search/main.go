package search

import (
	"bufio"
	"fmt"
	"os"
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
		res := engine.BoolSearch(text)
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Println(len(res), "results found in", elapsed)
		fmt.Println(res)
		fmt.Println("----")
	}
}

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
			subResults = append(subResults, (*engine).getQuerySubResults(q, "OR"))
		} else if isAnd(q) {
			subResults = append(subResults, (*engine).getQuerySubResults(q, "AND"))
		} else if isNot(q) {
			subResults = append(subResults, (*engine).getQuerySubResults(q, "NOT"))
		} else {
			subResults = append(subResults, (*engine.index)[(*engine.vocDict)[q]])
		}
	}
	res := engine.processSubResults(subResults, b.operator)
	return newBoolQueryGroup(b.q, b.operator, res)
}

func (engine *Engine) getQuerySubResults(q string, op string) []int {
	subQueries := parse(q, op)
	subQueryGroup := newBoolQueryGroup(subQueries, "OR", []int{})
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
