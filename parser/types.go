package parser

// Parser struct
type Parser struct {
	pathName          string
	stopWordsPathName string
	stopWords         []string
	data              string
	lines             []string
	docs              *Collection
}

// Document struct
type Document struct {
	docID          int
	title          string
	summary        string
	keywords       string
	tokens         []string
	filteredTokens []string
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
	return p.docs
}

// GetStopWords return the list of the stopwords
func (p *Parser) GetStopWords() []string {
	return p.stopWords
}

// newDocument creates a new Document
func newDocument() *Document {
	var tokens []string
	var filteredTokens []string
	return &Document{0, "", "", "", tokens, filteredTokens}
}

// GetDocID returns doc.docID
func (doc *Document) GetDocID() int {
	return doc.docID
}

// GetTitle returns doc.title
func (doc *Document) GetTitle() string {
	return doc.title
}

// GetSummary returns doc.summary
func (doc *Document) GetSummary() string {
	return doc.summary
}

// GetKeywords returns doc.keywords
func (doc *Document) GetKeywords() string {
	return doc.keywords
}

// GetTokens returns doc.tokens
func (doc *Document) GetTokens() []string {
	return doc.tokens
}

// GetFilteredTokens returns doc.filteredTokens
func (doc *Document) GetFilteredTokens() []string {
	return doc.filteredTokens
}

// SetDocID sets doc.docID to docID
func (doc *Document) SetDocID(docID int) {
	doc.docID = docID
}

// SetTitle sets doc.title to title
func (doc *Document) SetTitle(title string) {
	doc.title = title
}

// SetSummary sets doc.summary to summary
func (doc *Document) SetSummary(summary string) {
	doc.summary = summary
}

// SetKeywords sets doc.keywords to keywords
func (doc *Document) SetKeywords(keywords string) {
	doc.keywords = keywords
}
