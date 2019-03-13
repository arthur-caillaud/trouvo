package parser

// Collection struct
type Collection struct {
	docs []*Document
	voc  map[string]bool
}

// NewCollection creates a new Collection
func NewCollection(docs []*Document) *Collection {
	voc := make(map[string]bool)
	return &Collection{docs, voc}
}

// BuildVocabulary looks at all the words in the collection and build the vocabulary
func (col *Collection) BuildVocabulary() {
	for _, doc := range col.docs {
		for _, word := range doc.filteredTokens {
			if !col.isInVoc(word) {
				col.voc[word] = true
			}
		}
	}
}

func (col *Collection) isInVoc(word string) bool {
	_, ok := col.voc[word]
	return ok
}

// GetVocabulary get the vocabulary of the collection
func (col *Collection) GetVocabulary() map[string]bool {
	return col.voc
}

// GetDocs get the docs in the collection
func (col *Collection) GetDocs() []*Document {
	return col.docs
}
