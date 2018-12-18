package indexer

import (
	"trouvo/parser"
)

// Build build the index of the indexer
func (indexer *Indexer) Build() {
	docs := indexer.buildDocDict()
	voc := indexer.buildVocDict()
	for docID, doc := range docs {
		for _, token := range doc.GetFilteredTokens() {
			tokenID := voc[token]
			postingList := indexer.index[tokenID]
			postingList = append(postingList, docID)
			indexer.index[tokenID] = postingList
		}
	}
}

func (indexer *Indexer) buildVocDict() map[string]int {
	voc := indexer.col.GetVocabulary()
	for k, word := range voc {
		indexer.vocDict[word] = k
	}
	return indexer.vocDict
}

func (indexer *Indexer) buildDocDict() map[int]*parser.Document {
	docs := indexer.col.GetDocs()
	for k, doc := range docs {
		indexer.docDict[k] = doc
	}
	return indexer.docDict
}
