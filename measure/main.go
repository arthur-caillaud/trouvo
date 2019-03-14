package measure

import (
	"io/ioutil"
	"strconv"
	"strings"
	"trouvo/parser"
)

const (
	qrelsPathName     = "/Users/arthur/go/src/trouvo/Data/CACM/qrels.text"
	queryPathName     = "/Users/arthur/go/src/trouvo/Data/CACM/query.text"
	stopWordsPathName = "/Users/arthur/go/src/trouvo/Data/CACM/common_words"
)

func ParseQueries() []string {
	// Parse the query.text
	p := parser.New(queryPathName, stopWordsPathName)
	p.Run()
	docs := p.GetCollection().GetDocs()
	stopWords := p.GetStopWords()
	queries := []string{}
	// Tokenize and filter the queries
	for _, doc := range docs {
		doc.Tokenize()
		doc.FilterTokens(stopWords)
		// Build the query from the tokens
		q := ""
		for k, token := range doc.GetFilteredTokens() {
			if k == len(doc.GetFilteredTokens())-1 {
				q += token
			} else {
				q += token + " "
			}
		}
		// Append the resulting query to the queries
		queries = append(queries, q)
	}
	return queries
}

func ParseResults() map[int][]int {
	results := make(map[int][]int)
	data, _ := ioutil.ReadFile(qrelsPathName)
	fileContent := string(data)
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		if len(words) >= 2 {
			qID, _ := strconv.Atoi(words[0])
			docID, _ := strconv.Atoi(words[1])
			if _, ok := results[qID]; !ok {
				results[qID] = []int{}
			}
			results[qID] = append(results[qID], docID)
		}
	}
	return results
}

func CompareResults(trueRes []int, ourRes []int) (tp int, fp int, fn int) {
	truePositives := []int{}
	falsePositives := []int{}
	for _, res := range trueRes {
		for _, _res := range ourRes {
			if res == _res {
				truePositives = append(truePositives, _res)
				break
			}
		}
		falsePositives = append(falsePositives, res)
	}
	// Compute tp, fp, fn
	tp = len(truePositives)
	fp = len(falsePositives)
	fn = len(trueRes) - len(truePositives)
	return
}

func PrecisionMeasure(tp, fp int) float64 {
	return float64(tp) / float64(tp+fp)
}

func RecallMeasure(tp, fn int) float64 {
	return float64(tp) / float64(tp+fn)
}

func AccuracyMeasure(tp, tn, fn, fp int) float64 {
	return float64(tp+tn) / float64(tp+fn+fp+tn)
}

func EMeasure(alpha float64) float64 {
	// return 1 - 1 / (alpha * 1/PrecisionMeasure() + (1-alpha) * 1/RecallMeasure())
	return 0
}

func FMeasure(alpha float64) float64 {
	// return 1 - EMeasure(alpha)
	return 0
}

func AVG(scores []float64) (avg float64) {
	for _, score := range scores {
		avg += score
	}
	return 100 * avg / float64(len(scores))
}
