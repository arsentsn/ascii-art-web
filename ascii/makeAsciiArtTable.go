package ascii

import (
	"fmt"
	"os"
	"strings"
)

var banner []string

// MakeAsciiArtTable reads from a txt file in the parent folder
// and stores the contents line by line in a slice of strings
func MakeAsciiArtTable(bannerFile string) error {

	const bannerLines = 856

	// Read the banner file into a slice of bytes
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		return fmt.Errorf("couldn't read banner file: %w", err)

	}

	fileString := string(file)
	fileString = strings.ReplaceAll(fileString, "\r\n", "\n")

	banner = strings.Split(fileString, "\n")

	//check for correct number of lines
	if len(banner) != bannerLines {
		return fmt.Errorf("banner file has invalid number of lines")
	}

	return nil
}
