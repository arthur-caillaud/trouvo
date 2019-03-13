package indexer

import (
	"math"
)

// Build builds all the dicts and index of the indexer
func (indexer *Indexer) Build() {
	indexer.buildDocDict()
	indexer.buildVocDict()
	indexer.buildIndex()
	indexer.buildIdfDict()
	indexer.buildDocNormDict()
}

func (indexer *Indexer) buildIndex() {
	for docID, doc := range indexer.docDict {
		docLength := len(doc.GetFilteredTokens())
		for _, token := range doc.GetFilteredTokens() {
			tokenID := indexer.vocDict[token]
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
}

func (indexer *Indexer) buildIdfDict() {
	docDictLength := len(indexer.docDict)
	for tokenID, postings := range indexer.index {
		postingsLength := len(postings)
		idfRatio := float64(docDictLength) / float64(postingsLength)
		indexer.idfDict[tokenID] = math.Log10(idfRatio)
	}
}

func (indexer *Indexer) buildDocNormDict() {
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

func (indexer *Indexer) buildVocDict() {
	voc := indexer.col.GetVocabulary()
	tokenID := 0
	for word := range voc {
		indexer.vocDict[word] = tokenID
		tokenID++
	}
}

func (indexer *Indexer) buildDocDict() {
	docs := indexer.col.GetDocs()
	for k, doc := range docs {
		indexer.docDict[k] = doc
	}
}
