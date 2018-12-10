package main

import (
	"fmt"
	"gogole/indexer"
	"gogole/parser"
	"gogole/search"
)

const (
	pathName          = "/Users/arthur/go/src/gogole/Data/CACM/cacm.all"
	stopWordsPathName = "/Users/arthur/go/src/gogole/Data/CACM/common_words"
)

func main() {
	p := parser.New(pathName, stopWordsPathName)
	col := p.GetCollection()
	docs := col.GetDocs()
	stopWords := p.GetStopWords()
	for _, doc := range docs {
		doc.SetTokens()
		doc.FilterTokens(stopWords)
	}
	col.BuildVocabulary()
	indexer := indexer.NewIndexer(col)
	indexer.Build()
	engine := search.NewSearchEngine(indexer.GetIndex(), indexer.GetVocDict(), indexer.GetDocDict())
	res := engine.BoolSearch("slip || hypothesis")
	fmt.Println(res)
}
