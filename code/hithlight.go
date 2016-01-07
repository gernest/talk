package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"

	"github.com/gernest/frontman"
)

func main() {
	var (
		fileName   string
		themeName  string
		outputFile string
		inline     bool
		highlight  bool
	)

	flag.StringVar(&fileName, "f", "highlight.md", "Input markdown document")
	flag.StringVar(&themeName, "t", "prettify", "Theme name")
	flag.StringVar(&outputFile, "o", "intex.html", "File to save outuput")
	flag.BoolVar(&inline, "i", true, "enable inline style")
	flag.BoolVar(&highlight, "h", true, "enable syntax highlight")
	flag.Parse()

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	out := &bytes.Buffer{}
	err = frontman.Markdown(out, bytes.NewReader(data), highlight, inline, themeName)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(outputFile, out.Bytes(), 0600)
}
