package indexer

import (
	"math"
	"trouvo/parser"
)

// Build build the index of the indexer
func (indexer *Indexer) Build() {
	docDict := indexer.buildDocDict()
	docDictLength := len(docDict)
	vocDict := indexer.buildVocDict()
	for docID, doc := range docDict {
		docLength := len(doc.GetFilteredTokens())
		for _, token := range doc.GetFilteredTokens() {
			tokenID := vocDict[token]
			postingList := indexer.index[tokenID]
			if postingList == nil {
				postingList = make(map[int]float64)
			}
			if _, ok := postingList[docID]; !ok {
				postingList[docID] = 0
			}
			postingList[docID] += 1 / float64(docLength)
			indexer.index[tokenID] = postingList
		}
	}
	// Building idfDict
	for tokenID, postings := range indexer.index {
		postingsLength := len(postings)
		idfRatio := float64(docDictLength) / float64(postingsLength)
		indexer.idfDict[tokenID] = math.Log10(idfRatio)
	}
	// Building docNormDict
	for docID := range indexer.docDict {
		for tokenID, postings := range indexer.index {
			for _docID, termFrequency := range postings {
				if docID == _docID { // This token is in this document
					inverseDocFrequency := indexer.idfDict[tokenID]
					termWeight := termFrequency * inverseDocFrequency
					indexer.docNormDict[docID] += termWeight * termWeight
					break
				}
			}
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
