package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"trouvo/display"
	"trouvo/indexer"
	"trouvo/parser"
	// "trouvo/persist"
	"trouvo/measure"
	"trouvo/search"
)

func mainCACM() {
	indexer := buildCACM()
	measureCACM(indexer)
}

func buildCACM() *indexer.Indexer {
	p := parseCACM()
	col := tokenizeCACM(p)
	indexer := indexCACM(col)
	return indexer
}

func runCACM(indexer *indexer.Indexer) {
	engine := search.NewSearchEngine(
		indexer.GetIndex(),
		indexer.GetVocDict(),
		indexer.GetIdfDict(),
		indexer.GetDocDict(),
		indexer.GetDocNormDict(),
	)
	disp := display.New(indexer.GetDocDict())

	// Main infinite loop used to let the user query our search engine
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Type query :")
		text, _ := reader.ReadString('\n')
		start := time.Now()
		text = strings.TrimSpace(text)
		res := engine.VectSearch(text)
		end := time.Now()
		elapsed := end.Sub(start).Round(time.Microsecond)
		disp.Show(res, elapsed)
	}
}

func measureCACM(indexer *indexer.Indexer) {
	fmt.Println("Measuring relevance...")
	// Build the SearchEngine that will be measured
	engine := search.NewSearchEngine(
		indexer.GetIndex(),
		indexer.GetVocDict(),
		indexer.GetIdfDict(),
		indexer.GetDocDict(),
		indexer.GetDocNormDict(),
	)
	// parse trueResults and queries from query.text and qrels.text
	queries := measure.ParseQueries()
	trueResults := measure.ParseResults()
	colSize := len(*indexer.GetDocDict())
	// We store the scores to compute averages
	precisionScores, recallScores, accuracyScores := collectMeasuresCACM(engine, queries, trueResults, colSize)
	// Print the averages of these scores
	avgPrecision := measure.AVG(precisionScores)
	avgRecall := measure.AVG(recallScores)
	avgAccuracy := measure.AVG(accuracyScores)
	fmt.Println("Average Precision :", avgPrecision, "%")
	fmt.Println("Average Recall :", avgRecall, "%")
	fmt.Println("Average Accuracy :", avgAccuracy, "%")
}

func parseCACM() *parser.Parser {
	start := time.Now()
	p := parser.New(pathNameCACM, stopWordsPathName)
	p.Run() // Parsing...
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Parsed in", elapsed.Round(time.Millisecond))
	fmt.Println("----")
	return p
}

func tokenizeCACM(p *parser.Parser) *parser.Collection {
	start := time.Now()
	col := p.GetCollection()
	docs := col.GetDocs()
	stopWords := p.GetStopWords()
	for _, doc := range docs {
		doc.Tokenize()
		doc.FilterTokens(stopWords)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Tokenized in", elapsed.Round(time.Millisecond))
	fmt.Println("----")
	start = time.Now()
	col.BuildVocabulary()
	fmt.Println(len(col.GetVocabulary()), "words in vocabulary")
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("Vocabulary built in", elapsed.Round(time.Millisecond))
	fmt.Println("----")
	return col
}

func indexCACM(col *parser.Collection) *indexer.Indexer {
	start := time.Now()
	indexer := indexer.New(col)
	indexer.Build()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Indexed in", elapsed.Round(time.Millisecond))
	fmt.Println("Index is", indexer.GetIndexSize(), "kB large.")
	fmt.Println("----")
	return indexer
}

func collectMeasuresCACM(
	engine *search.Engine,
	queries []string,
	trueResults map[int][]int,
	colSize int,
) ([]float64, []float64, []float64) {
	precisionScores := []float64{}
	recallScores := []float64{}
	accuracyScores := []float64{}
	// Run each measure query
	for qID, q := range queries {
		engineResults := engine.VectSearch(q)
		ourRes := []int{}
		for _, res := range engineResults {
			ourRes = append(ourRes, res.GetDocID())
		}
		trueRes := trueResults[qID+1]
		// Computing tp, fp, fn, tn by compring ourResults and trueResults
		tp, fp, fn := measure.CompareResults(ourRes, trueRes)
		tn := colSize - tp - fp - fn
		// Computing precision, recall, accuract
		precision := measure.PrecisionMeasure(tp, fp)
		recall := measure.RecallMeasure(tp, fn)
		accuracy := measure.AccuracyMeasure(tp, tn, fn, fp)
		// Append the results in their slices
		precisionScores = append(precisionScores, precision)
		recallScores = append(recallScores, recall)
		accuracyScores = append(accuracyScores, accuracy)
	}
	return precisionScores, recallScores, accuracyScores
}
