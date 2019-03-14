# Trouvo
A simple search engine written in Go

## Quickstart

Two ways of starting the project

```bash
# Directly runs the main source file
go run main.go
```

```bash
# At first it compiles the source code in binary
go build trouvo
# And then it runs the binary compiled version of the project
./trouvo
```

> **The data are not versioned but they need to be found in a data/ folder at the root of this repository**

## Packages

### parser

This package contains the core structures Document and Collection used in all this project.

```go
// Document struct
type Document struct {
	docID          int
	title          string
	summary        string
	keywords       string
	tokens         []string
	filteredTokens []string
}
```

```go
// Collection struct
type Collection struct {
	docs []*Document
	voc  map[string]bool
}
```
This package also contains the CACM parser used to parse the documents contained in the CACM collection.

```go
// Parser struct
type Parser struct {
	pathName          string
	stopWordsPathName string
	stopWords         []string
	data              string
	lines             []string
	docs              *Collection
}
```

### cs276parser

This package contains everything needed to parse the documents from the CS276 document collection.
```go
type Parser struct {
	rootPath    string
	collections []*parser.Collection
}
```

### indexer

This package contains the API used to build an index from a parsed documents collection
```go
// Indexer struct
type Indexer struct {
	col         *parser.Collection       // Collection
	index       map[int]map[int]float64  // tokenID => docID => termFrequecy
	vocDict     map[string]int           // 'cat' => tokenID
	idfDict     map[int]float64          // tokenID => inverseDocFrequency
	docDict     map[int]*parser.Document // docID => Document
	docNormDict map[int]float64          // docID => docNorm
}
```
The indexer builds
- reversed index under Indexer.index
- the mapping of a token to its tokenID under Indexer.vocDict
- the mapping of a tokenId to its inverse document frequecy (IDF)
- the mapping of a documentID to the pointer of the Document in memory
- the mapping of a documentID to its norm

### search

This package contains the API used to perform boolean and vectorial queries on a collection.

A *SearchEngine* needs the content of an *Indexer* to properly work.

```go
// BoolSearch runs a boolean query with the SearchEngine
func (engine *Engine) BoolSearch(q string) (res []int)
```

```go
// VectSearch runs a vectorial query with the SearchEngine and cos measure
func (engine *Engine) VectSearch(q string) (res []*Result)
```

A *SuperEngine* struct is provided to perform a *VectSearch* on multiple *SearchEngine* simultaneously

```go
// Search performs a SuperEngine search on all its sub-engines
func (superEngine *SuperEngine) Search(q string) (res []*Result)
```

### display

This package contains the API we use to display the document informations in the terminal.

```go
func (display *Display) Show(results []*search.Result, elapsed time.Duration)
```
Displays the title of the documents of up to the first 10 results in the *results* slice provided. The *elapsed* duration is used to provide information on the performance of the query that was executed.

### persist

This package contains the API used to store the index to the disk.

> ** This API is not fully working for now **

```go
// Save saves a representation of v to the file at path.
func Save(path string, v interface{}) error
```
```go
// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func Load(path string, v interface{}) error
```
