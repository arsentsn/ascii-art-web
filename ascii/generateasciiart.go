package ascii

import (
	"fmt"
	"log"
)

func GenerateAsciiArt(text, bannerFile string) (asciiArt string, err error) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			err = fmt.Errorf("internal server error: %v", r)
			asciiArt = "There was an unexpected server error"
		}
	}()

	//this is to test panic.
	//To test, uncomment this and the last two radio buttons in index.html.
	//Also comment out banner validation in mainPage
	// if bannerFile == "" {
	// 	panic("banner file is empty")
	// }

	// read banner file
	if err = MakeAsciiArtTable("banners/" + bannerFile + ".txt"); err != nil {

		asciiArt = ("Internal server error: Couldn't read banner file")
		return asciiArt, err

	}

	return AsciiPrint(text), nil
}
