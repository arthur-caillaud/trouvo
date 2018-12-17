package main

import (
	"fmt"
	"time"
	"trouvo/indexer"
	"trouvo/parser"
	"trouvo/search"
)

const (
	pathName          = "/Users/arthur/go/src/trouvo/Data/CACM/cacm.all"
	stopWordsPathName = "/Users/arthur/go/src/trouvo/Data/CACM/common_words"
)

func main() {

	fmt.Println("----")
	start := time.Now()
	p := parser.New(pathName, stopWordsPathName)
	col := p.GetCollection()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Parsed in", elapsed)
	fmt.Println("----")

	start = time.Now()
	docs := col.GetDocs()
	stopWords := p.GetStopWords()
	for _, doc := range docs {
		doc.SetTokens()
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
	fmt.Println("----")

	start = time.Now()
	engine := search.NewSearchEngine(indexer.GetIndex(), indexer.GetVocDict(), indexer.GetDocDict())
	res := engine.BoolSearch("cat&&dog||house&&platform||a&&!a&&b")
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("Results found in", elapsed)
	fmt.Println("----")
	fmt.Println(res)
}
