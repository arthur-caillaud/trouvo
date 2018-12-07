package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Document struct {
	docId    int
	title    string
	summary  string
	keywords []string
}

func newDocument() Document {
	var docId int
	var title string
	var summary string
	var keywords []string
	return Document{docId, title, summary, keywords}
}

func openFile(pathName string) string {
	dat, err := ioutil.ReadFile(pathName)
	check(err)
	f := string(dat)
	return f
}

func readSeparator(line string) string {
	if len(line) >= 2 {
		sep := line[:2]
		return sep
	}
	return ""
}

func main() {
	pathName := "/Users/arthur/Documents/Centrale/Moteur de recherche/project/Data/CACM/cacm.all"
	separators := [8]string{".I", ".T", ".W", ".B", ".A", ".N", ".X", ".K"}
	f := openFile(pathName)
	docs := strings.Split(f, "\n")
	var Docs []Document
	doc := newDocument()
	nextLineHas := ""
	for _, line := range docs {
		sep := readSeparator(line)
		isSep := false
		for _, _sep := range separators {
			if sep == _sep {
				isSep = true
			}
		}
		if isSep {
			switch sep {
			case ".I":
				Docs = append(Docs, doc)
				doc = newDocument()
				docId, _ := strconv.Atoi(line[3:])
				doc.docId = docId
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
				doc.title += line
			case "W":
				doc.summary += line + " "
			case "K":
				doc.keywords = append(doc.keywords, strings.Split(line, " ")...)
			}
		}
	}
	Docs = Docs[1:]
}
