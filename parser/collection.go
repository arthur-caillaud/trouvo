package parser

// Collection struct
type Collection struct {
	docs []*Document
	voc  []string
}

// NewCollection creates a new Collection
func NewCollection(docs []*Document) *Collection {
	var voc []string
	return &Collection{docs, voc}
}

// BuildVocabulary looks at all the words in the collection and build the vocabulary
func (col *Collection) BuildVocabulary() {
	for _, doc := range col.docs {
		for _, word := range doc.filteredTokens {
			if !col.isInVoc(word) {
				col.voc = append(col.voc, word)
			}
		}
	}
}

func (col *Collection) isInVoc(word string) bool {
	for _, _word := range col.voc {
		if word == _word {
			return true
		}
	}
	return false
}

// GetVocabulary get the vocabulary of the collection
func (col *Collection) GetVocabulary() []string {
	return col.voc
}

// GetDocs get the docs in the collection
func (col *Collection) GetDocs() []*Document {
	return col.docs
}
