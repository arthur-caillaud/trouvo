package parser

import (
	"regexp"
	"strings"
)

const wordRegExp = `\b[A-Za-z]+\b`

// Tokenize extract all the words in a Document and transform them into tokens
func (doc *Document) Tokenize() {
	re := regexp.MustCompile(wordRegExp)
	titleTokens := re.FindAllString(doc.title, -1)
	summaryTokens := re.FindAllString(doc.summary, -1)
	keywordTokens := re.FindAllString(doc.keywords, -1)
	var tokens []string
	tokens = append(tokens, titleTokens...)
	tokens = append(tokens, summaryTokens...)
	tokens = append(tokens, keywordTokens...)
	for _, token := range tokens {
		token = strings.ToLower(token)
		if token != "" {
			doc.tokens = append(doc.tokens, token)
		}
	}
}

// FilterTokens filter the tokens using the stopWords list provided
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
