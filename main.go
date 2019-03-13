package main

import (
	"fmt"
	"time"
	"trouvo/cs276parser"
	"trouvo/display"
	"trouvo/indexer"
	"trouvo/parser"
	"trouvo/search"
)

const (
	pathNameCACM      = "/Users/arthur/go/src/trouvo/Data/CACM/cacm.all"
	pathNameCS276     = "/Users/arthur/go/src/trouvo/Data/CS276"
	stopWordsPathName = "/Users/arthur/go/src/trouvo/Data/CACM/common_words"
)

func main() {
	mainCS276()
}

func mainCACM() {

	start := time.Now()
	p := parser.New(pathNameCACM, stopWordsPathName)
	p.Run() // Parsing...
	col := p.GetCollection()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Parsed in", elapsed)
	fmt.Println("----")

	start = time.Now()
	docs := col.GetDocs()
	stopWords := p.GetStopWords()
	for _, doc := range docs {
		doc.Tokenize()
		doc.FilterTokens(stopWords)
	}
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("Tokenized in", elapsed)
	fmt.Println("----")

	start = time.Now()
	col.BuildVocabulary()
	fmt.Println(len(col.GetVocabulary()), "words in vocabulary")
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("Vocabulary built in", elapsed)
	fmt.Println("----")

	start = time.Now()
	indexer := indexer.New(col)
	indexer.Build()
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("Indexed in", elapsed)
	fmt.Println("Index is", indexer.GetIndexSize(), "kB large.")
	fmt.Println("----")

	engine := search.NewSearchEngine(
		indexer.GetIndex(),
		indexer.GetVocDict(),
		indexer.GetIdfDict(),
		indexer.GetDocDict(),
		indexer.GetDocNormDict(),
	)
	disp := display.New(indexer.GetDocDict())
	engine.Run(disp) // Run the SearchEngine
}

func mainCS276() {

	start := time.Now()
	p := cs276parser.New(pathNameCS276)
	p.Run() // Parsing...
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Parsed in", elapsed)
	fmt.Println("----")

	start = time.Now()
	cols := p.GetCollections()
	for _, col := range cols {
		col.BuildVocabulary()
	}
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("Vocabulary built in", elapsed)
	fmt.Println("----")

	start = time.Now()
	indexer := indexer.New(cols[0])
	indexer.Build()
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("Indexed in", elapsed)
	fmt.Println("Index is", indexer.GetIndexSize(), "kB large.")
	fmt.Println("----")

	engine := search.NewSearchEngine(
		indexer.GetIndex(),
		indexer.GetVocDict(),
		indexer.GetIdfDict(),
		indexer.GetDocDict(),
		indexer.GetDocNormDict(),
	)
	disp := display.New(indexer.GetDocDict())
	engine.Run(disp)
}
