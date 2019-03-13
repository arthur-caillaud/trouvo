package cs276parser

import "trouvo/parser"

type Parser struct {
	rootPath    string
	collections []*parser.Collection
}

func New(pathName string) *Parser {
	collections := []*parser.Collection{}
	return &Parser{pathName, collections}
}
