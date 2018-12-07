package parser

import (
	//"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var separators = [8]string{".I", ".T", ".W", ".B", ".A", ".N", ".X", ".K"}

// Parser struct
type Parser struct {
	pathName          string
	stopWordsPathName string
	stopWords         []string
	data              string
	lines             []string
	docs              *Collection
}

// New Parser
func New(pathName string, stopWordsPathName string) *Parser {
	var stopWords []string
	var data string
	var lines []string
	var docs Collection
	return &Parser{pathName, stopWordsPathName, stopWords, data, lines, &docs}
}

// GetCollection get the collection from the parser
func (p *Parser) GetCollection() *Collection {
	p.run()
	return p.docs
}

// GetStopWords return the list of the stopwords
func (p *Parser) GetStopWords() []string {
	return p.stopWords
}

// Run the parser so that it parse all the documents
func (p *Parser) run() {
	p.openFile()
	p.loadStopWords()
	p.parseFile()
	p.parseDocuments()
}

func (p *Parser) openFile() {
	dat, _ := ioutil.ReadFile(p.pathName)
	f := string(dat)
	p.data = f
}

func (p *Parser) loadStopWords() {
	f, _ := ioutil.ReadFile(p.stopWordsPathName)
	data := string(f)
	p.stopWords = strings.Split(data, "\n")
}

func (p *Parser) parseFile() {
	lines := strings.Split(p.data, "\n")
	p.lines = lines
}

func readSeparator(line string) string {
	if len(line) >= 2 {
		sep := line[:2]
		return sep
	}
	return ""
}

func isSeparator(sep string) bool {
	isSep := false
	for _, _sep := range separators {
		if sep == _sep {
			isSep = true
		}
	}
	return isSep
}

func (p *Parser) parseDocuments() {
	var docs []*Document
	doc := newDocument()
	nextLineHas := ""
	for _, line := range p.lines {
		sep := readSeparator(line)
		isSep := isSeparator(sep)
		if isSep {
			switch sep {
			case ".I":
				docs = append(docs, doc)
				doc = newDocument()
				docID, _ := strconv.Atoi(line[3:])
				doc.SetDocID(docID)
				nextLineHas = ""
			case ".T":
				nextLineHas = "T"
			case ".W":
				nextLineHas = "W"
			case ".K":
				nextLineHas = "K"
			default:
				nextLineHas = ""
			}
		} else {
			switch nextLineHas {
			case "T":
				doc.SetTitle(line)
			case "W":
				doc.SetSummary(doc.GetSummary() + line + " ")
			case "K":
				doc.SetKeywords(append(doc.GetKeywords(), strings.Split(line, " ")...))
			}
		}
	}
	docs = docs[1:]
	p.docs = NewCollection(docs)
}
