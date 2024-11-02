package ascii

// asciiPrint takes a string and returns a string with its ascii-art equivalent
// it handles the preprocessing for linebreaks and sends each line separately for conversion
func AsciiPrint(message string) string {

	// make a slice of strings where linebreaks are separate from the rest of the text
	textLines := formatInput(message)

	var output string

	//store each line of text
	for _, line := range textLines {
		if line == "\n" {
			output += "\n"
		} else {
			output += asciiLinePrint(line)
		}
	}
	return output
}
