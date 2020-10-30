package cmd

import (
	"flag"
	"fmt"
	"jack-watts/empty-dci-tt/tt"
	"os"
)

var (
	text      bool
	image     bool
	track     bool
	encrypt   bool
	reel      int
	duration  int
	display   int
	frameRate string
	template  string
	title     string
	lang      string
	output    string
)

// Execute ...
func Execute() {

	flag.BoolVar(&text, "text", true, "- Inidcate that text profile is to be used.")
	flag.BoolVar(&image, "image", false, "- Inidcate that image profile is to be used.")
	flag.BoolVar(&track, "T", false, "- write MXF trackfile, requires '-d'")
	flag.BoolVar(&encrypt, "e", false, "- encrypt trackfile")
	flag.IntVar(&duration, "d", 24, "- set the duration of the track file. Default is '24'")
	flag.StringVar(&frameRate, "p", "24", "- set the frame rate of the track file. Default is '24'")
	flag.IntVar(&display, "m", 0, "- set the DisplayType.'0'=MainSubtitle,'1'=ClosedCaption. Default ='0'")
	flag.IntVar(&reel, "r", 1, "- set the ReelNumber, Default ='1'")
	flag.StringVar(&lang, "l", "en", "- set the RFC 5646 Language subtag")
	flag.StringVar(&title, "t", "No Title", "- set the ContentTitleText value.")
	flag.StringVar(&template, "x", "", "- path to 428-7 XML to use as template")
	flag.StringVar(&output, "o", "", "- set the output path, Default is StdOut")
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Println("check command expression")
		os.Exit(1)
	}
	if err := tt.CreateXML(text, image, track, encrypt, reel, display, duration, frameRate, lang, title, template, output); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}
