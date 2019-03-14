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
	// indexer := buildCACM()
	// runCACM(indexer)
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
	engine := search.NewSearchEngine(
		indexer.GetIndex(),
		indexer.GetVocDict(),
		indexer.GetIdfDict(),
		indexer.GetDocDict(),
		indexer.GetDocNormDict(),
	)
	queries := measure.ParseQueries()
	for k, q := range queries {
		results := engine.VectSearch(q)
		if k == 0 {
			for _, res := range results {
				fmt.Println(res.GetDocID())
			}
		}
	}
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
