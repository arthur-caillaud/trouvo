package measure

import (
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

func PrecisionMeasure() float64 {
	// return truePositive / (truePositive + falsePositive)
	return 0
}

func RecallMeasure() float64 {
	// return truePositive / (truePositive + falseNegative)
	return 0
}

func AccuracyMeasure() float64 {
	// return (truePositive + trueNegative) / (truePositive + falseNegative + falsePositive + trueNegative)
	return 0
}

func EMeasure(alpha float64) float64 {
	// return 1 - 1 / (alpha * 1/PrecisionMeasure() + (1-alpha) * 1/RecallMeasure())
	return 0
}

func FMeasure(alpha float64) float64 {
	// return 1 - EMeasure(alpha)
	return 0
}
