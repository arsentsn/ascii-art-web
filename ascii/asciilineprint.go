package ascii

// asciiLinePrint takes a sub-string (without line breaks) and returns its ascii art equivalent
func asciiLinePrint(text string) string {

	var lineOfText string
	for height := 0; height < 8; height++ {
		for _, r := range text {
			lineOfText += (banner[((r-' ')*9)+1+rune(height)])

		}

		lineOfText += "\n"

	}
	return lineOfText

}
