package display

import (
	"fmt"
	"time"
)

// Show displays a summary of the document with docID
func (display *Display) Show(res []int, elapsed time.Duration) {
	fmt.Println(len(res), "results found in", elapsed)
	fmt.Println("----")
	for k, docID := range res {
		if k < 10 { // We only show 10 results (10RPP)
			doc := (*display.docDict)[docID]
			docTitle := doc.GetTitle()
			fmt.Println(docTitle)
		} else {
			break
		}
	}
	fmt.Println("----")
}
