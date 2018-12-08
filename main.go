package main

import (
	"fmt"
	"gogole/parser"
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
	fmt.Println("Vocabulary :", col.GetVocabulary())
}
