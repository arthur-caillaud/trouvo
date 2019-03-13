package cs276parser

import (
	"fmt"
	"io/ioutil"
	"strings"
	"trouvo/parser"
)

// Run runs the CS276Parser on the doc collection
func (p *Parser) Run() {
	dirs := p.openDirectory(p.rootPath)
	docID := 0
	for _, dirPath := range dirs {
		docs := []*parser.Document{}
		files := p.openDirectory(dirPath)
		for _, filePath := range files {
			fileContent := p.openFile(filePath)
			fileName := strings.Split(filePath, dirPath+"/")[1]
			tokens := p.parseFile(fileContent)
			doc := parser.NewDocument()
			doc.SetTitle(fileName)
			doc.SetDocID(docID)
			doc.SetFilteredTokens(tokens)
			docs = append(docs, doc)
			docID++
		}
		col := parser.NewCollection(docs)
		p.collections = append(p.collections, col)
	}
}

func (p *Parser) openDirectory(pathName string) []string {
	files, err := ioutil.ReadDir(pathName)
	filePaths := []string{}
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		filePaths = append(filePaths, pathName+"/"+f.Name())
	}

	return filePaths
}

func (p *Parser) openFile(pathName string) string {
	data, _ := ioutil.ReadFile(pathName)
	return string(data)
}

func (p *Parser) parseFile(fileContent string) []string {
	return strings.Split(fileContent, " ")
}
