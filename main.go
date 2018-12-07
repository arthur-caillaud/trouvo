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
	docs := p.GetDocs()
	stopWords := p.GetStopWords()
	for _, doc := range docs {
		doc.SetTokens()
		doc.FilterTokens(stopWords)
	}
	fmt.Println("Tokens :", docs[120].GetTokens())
	fmt.Println("Filtered tokens :", docs[120].GetFilteredTokens())
}
