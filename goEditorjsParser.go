package main

import (
	"flag"
	"fmt"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/pkg/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/pkg/markdown"
	"log"
	"os"
)

func main() {

	jsonPath := flag.String("j", "", "Path to a JSON file")
	stylePath := flag.String("c", "default", "Style CSS path to be used. If not set, the default style will be used")
	outputFile := flag.String("o", "", "Output file path. If not set, root path will be used")
	outputType := flag.String("t", "", "Output file type. Possible values are: markdown or html")

	flag.Parse()

	var err error
	if *jsonPath == "" {
		usage()
		return
	}

	if *outputType == "" || *outputType == "html" {
		err = html.Default(*jsonPath, *outputFile, *stylePath)
		if err != nil {
			log.Println("It was not possible to parse json to html file\n",err)
			return
		}
	}

	if *outputType == "" || *outputType == "markdown" {
		err = markdown.Parser(*jsonPath, *outputFile)
		if err != nil {
			log.Println("It was not possible to parse json to markdown file\n", err)
			return
		}
	}

}

func usage() {
	fmt.Printf("\n%s -j <JSONFilePath> -c <StylePath> -o <OutputFilePath>\n\n", os.Args[0])
	fmt.Println("-j = Path to a JSON file. MANDATORY")
	fmt.Println("-c = Style CSS path to be used. If not set, the default style will be used")
	fmt.Println("-o = Output file path. If not set, root path will be used")
	fmt.Printf("-t = Output file type. Possible values: markdown or html\n\n")
}