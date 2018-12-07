package parser

import "strings"

// Document struct
type Document struct {
	docID          int
	title          string
	summary        string
	keywords       []string
	tokens         []string
	filteredTokens []string
}

// newDocument creates a new Document
func newDocument() *Document {
	var keywords []string
	var tokens []string
	var filteredTokens []string
	return &Document{0, "", "", keywords, tokens, filteredTokens}
}

func (doc *Document) GetDocID() int {
	return doc.docID
}

func (doc *Document) GetTitle() string {
	return doc.title
}

func (doc *Document) GetSummary() string {
	return doc.summary
}

func (doc *Document) GetKeywords() []string {
	return doc.keywords
}

func (doc *Document) GetTokens() []string {
	return doc.tokens
}

func (doc *Document) GetFilteredTokens() []string {
	return doc.filteredTokens
}

func (doc *Document) SetDocID(docID int) {
	doc.docID = docID
}

func (doc *Document) SetTitle(title string) {
	doc.title = title
}

func (doc *Document) SetSummary(summary string) {
	doc.summary = summary
}

func (doc *Document) SetKeywords(keywords []string) {
	doc.keywords = keywords
}

func (doc *Document) SetTokens() {
	titleTokens := strings.Split(doc.title, " ")
	summaryTokens := strings.Split(doc.summary, " ")
	var tokens []string
	tokens = append(tokens, titleTokens...)
	tokens = append(tokens, summaryTokens...)
	tokens = append(tokens, doc.keywords...)
	for _, token := range tokens {
		token = strings.ToLower(token)
		if token != "" {
			doc.tokens = append(doc.tokens, token)
		}
	}
}

func (doc *Document) FilterTokens(stopWords []string) {
	for _, token := range doc.tokens {
		isStopWord := false
		for _, stopWord := range stopWords {
			if token == stopWord {
				isStopWord = true
				break
			}
		}
		if !isStopWord {
			doc.filteredTokens = append(doc.filteredTokens, token)
		}
	}
}
