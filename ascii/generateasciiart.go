package ascii

import "log"

func GenerateAsciiArt(text, bannerFile string) string {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("%v:recovered", r)
			return
		}
	}()

	// read banner file
	MakeAsciiArtTable("banners/" + bannerFile + ".txt")

	return AsciiPrint(text)
}
