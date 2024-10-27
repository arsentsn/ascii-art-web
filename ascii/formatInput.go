package ascii

import (
	"strings"
)

// formatInput receives a string and returns a slice of strings with line breaks as separate elements

func formatInput(message string) []string {
	// replace escaped line breaks with actual line breaks
	// so the test cases work as intended
	message = strings.ReplaceAll(message, "\r\n", "\n")

	var textLines []string
	linebreak := false
	for i, r := range message {
		if r == '\n' { // check for line breaks, do not append the first instance of several line breaks in a row
			if linebreak || i == 0 {
				textLines = append(textLines, string(r))
			}
			linebreak = true
		} else if r >= ' ' && r <= '~' { // only allow printable ASCII
			if len(textLines) == 0 || linebreak { // append if no lines or previous was a line break
				textLines = append(textLines, string(r))
			} else {
				textLines[len(textLines)-1] += string(r) // append to the last line
			}
			linebreak = false // Reset linebreak since we appended valid character
		}
	}

	return textLines
}
