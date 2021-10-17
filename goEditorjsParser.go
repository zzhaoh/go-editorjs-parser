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
	style := flag.String("s", "sample", "Style CSS to be used. If not set, the default style will be used. Possible values: \"default\", \"bootstrap5\" or \"PATH/TO/YOUR/CUSTOM/STYLE/MAP\"")
	outputFile := flag.String("o", "", "Output file path. If not set, root path will be used")
	outputType := flag.String("t", "", "Output file type. Possible values are: markdown or html")

	flag.Parse()

	var err error
	if *jsonPath == "" {
		usage()
		return
	}

	if *outputType == "" || *outputType == "html" {
		switch *style {
		case "default":
			err = html.Sample(*jsonPath, *outputFile)
		case "bootstrap5":
			err = html.Bootstrap5(*jsonPath, *outputFile)
		default:
			err = html.Custom(*jsonPath, *outputFile, *style)
		}
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
	fmt.Printf("\n%s -j <JSONFilePath> -s <Style> -o <OutputFilePath> -t <OutputFileType>\n\n", os.Args[0])
	fmt.Println("-j = Path to a JSON file. MANDATORY")
	fmt.Println("-s = Style CSS to be used. If not set, the default style will be used. Possible values: \"default\", \"bootstrap5\" or \"PATH/TO/YOUR/CUSTOM/STYLE/MAP\"")
	fmt.Println("-o = Output file path. If not set, root path will be used")
	fmt.Printf("-t = Output file type. Possible values: markdown or html\n\n")
}