package parser

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var separators = [8]string{".I", ".T", ".W", ".B", ".A", ".N", ".X", ".K"}

// Run the parser so that it parse all the documents
func (p *Parser) Run() {
	p.openFile()
	p.loadStopWords()
	p.parseFile()
	p.parseDocuments()
}

func (p *Parser) openFile() {
	data, _ := ioutil.ReadFile(p.pathName)
	p.data = string(data)
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
	doc := NewDocument()
	nextLineHas := ""
	for _, line := range p.lines {
		sep := readSeparator(line)
		isSep := isSeparator(sep)
		if isSep {
			switch sep {
			case ".I":
				docs = append(docs, doc)
				doc = NewDocument()
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
				doc.SetTitle(doc.GetTitle() + line + " ")
			case "W":
				doc.SetSummary(doc.GetSummary() + line + " ")
			case "K":
				doc.SetKeywords(doc.GetKeywords() + line + " ")
			}
		}
	}
	docs = docs[1:]
	p.docs = NewCollection(docs)
}
