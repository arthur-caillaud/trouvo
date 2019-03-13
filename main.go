package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"trouvo/cs276parser"
	"trouvo/display"
	"trouvo/indexer"
	"trouvo/parser"
	"trouvo/persist"
	"trouvo/search"
)

const (
	pathNameCACM       = "/Users/arthur/go/src/trouvo/Data/CACM/cacm.all"
	indexPathNameCACM  = "/Users/arthur/go/src/trouvo/Data/CACM/index.idx"
	pathNameCS276      = "/Users/arthur/go/src/trouvo/Data/CS276"
	indexPathNameCS276 = "/Users/arthur/go/src/trouvo/Data/CS276/index.idx"
	stopWordsPathName  = "/Users/arthur/go/src/trouvo/Data/CACM/common_words"
)

func main() {
	mainCACM()
}

func mainCACM() {
	// load the indexer if stored on disk
	var indexer indexer.Indexer
	if err := persist.Load(indexPathNameCACM, indexer); err != nil {
		fmt.Println(err)
		buildCACM()
	} else {
		runCACM(indexer)
	}
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
	engines := []*search.Engine{}
	indexSize := 0
	docDict := make(map[int]*parser.Document)
	for _, col := range cols {
		indexer := indexer.New(col)
		indexer.Build()
		indexSize += indexer.GetIndexSize()

		engine := search.NewSearchEngine(
			indexer.GetIndex(),
			indexer.GetVocDict(),
			indexer.GetIdfDict(),
			indexer.GetDocDict(),
			indexer.GetDocNormDict(),
		)
		for docID, doc := range *indexer.GetDocDict() {
			docDict[docID] = doc
		}
		engines = append(engines, engine)
	}
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("Indexed in", elapsed)
	fmt.Println("Index is", indexSize, "kB large.")
	fmt.Println("----")

	superEngine := search.NewSuperEngine(engines)
	disp := display.New(&docDict)

	// Main infinite loop used to let the user query our search engine
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Type query :")
		text, _ := reader.ReadString('\n')
		start := time.Now()
		text = strings.TrimSpace(text)
		res := superEngine.Search(text)
		end := time.Now()
		elapsed := end.Sub(start)
		disp.Show(res, elapsed)
	}
}

func buildCACM() {
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

	if err := persist.Save(indexPathNameCACM, &indexer); err != nil {
		fmt.Println(err)
	}
}

func runCACM(indexer indexer.Indexer) {
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
		elapsed := end.Sub(start)
		disp.Show(res, elapsed)
	}
}
