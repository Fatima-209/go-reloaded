package main

import "regexp"

// cleaning up punctuations
func preprocessText(text string) string {
	// removes spaces between words and a single punctuation mark
	// and ensures there is one space after the punctuation.
	endPunct := `(\w+)\s+([.'"!?:;,]{1,})\s*`
	// a word followed by multiple spaces and followed by punctuations
	multiPunct := `(\w+)\s+([\.\!?\:;,'"]{2,})`
	// removes spaces between 2 punctuation marks
	endmultiPunct := `([\.\!?\'"])\s+([\.,\!?\'"])`
	// removes spaces bet single quotes & words
	startPunct := `(['])\s+(\w+)`

	text1 := regexp.MustCompile(startPunct)
	text = text1.ReplaceAllString(text, "$1$2")
	text1 = regexp.MustCompile(endmultiPunct)
	text = text1.ReplaceAllString(text, "$1$2")
	text2 := regexp.MustCompile(multiPunct)
	text = text2.ReplaceAllString(text, "$1$2")
	text1 = regexp.MustCompile(endPunct)
	text = text1.ReplaceAllString(text, "$1$2 ")

	return text
}
