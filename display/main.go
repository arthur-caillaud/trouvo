package display

import (
	"fmt"
	"time"
	"trouvo/search"
)

// Show displays a summary of the document with docID
func (display *Display) Show(results []*search.Result, elapsed time.Duration) {
	fmt.Println(len(results), "results found in", elapsed)
	fmt.Println("----")
	for k, res := range results {
		if k < 10 { // We only show 10 results (10RPP)
			doc := (*display.docDict)[res.GetDocID()]
			docTitle := doc.GetTitle()
			fmt.Println(docTitle)
		} else {
			break
		}
	}
	fmt.Println("----")
}
