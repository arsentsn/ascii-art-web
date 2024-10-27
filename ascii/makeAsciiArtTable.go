package ascii

import (
	"os"
	"strings"
)

var banner []string

// MakeAsciiArtTable reads from a txt file in the parent folder
// and stores the contents line by line in a slice of strings
func MakeAsciiArtTable(bannerFile string) {
	// Read the banner file into a slice of bytes
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		panic("Couldn't read banner file")

	}

	fileString := string(file)
	fileString = strings.ReplaceAll(fileString, "\r\n", "\n")

	banner = strings.Split(fileString, "\n")
}
