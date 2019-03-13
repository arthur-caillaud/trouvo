package display

import "fmt"

// Show displays a summary of the document with docID
func (display *Display) Show(docID int) {
	doc := (*display.docDict)[docID]
	docTitle := doc.GetTitle()
	fmt.Println(docTitle)
}
