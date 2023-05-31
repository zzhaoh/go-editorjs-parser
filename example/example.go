package main

import (
	"fmt"
	"gitlab.com/zzhaoh/go-editorjs-parser/pkg/html"
	"gitlab.com/zzhaoh/go-editorjs-parser/pkg/markdown"
	"io"
	"log"
	"os"
)

func main() {

	err := html.Sample("example/example.json", "example/sample.html")
	if err != nil {
		log.Println("It was not possible to parse json to html file\n", err)
		return
	}

	err = html.Bootstrap("example/example.json", "example/bootstrap.html")
	if err != nil {
		log.Println("It was not possible to parse json to html file\n", err)
		return
	}

	err = html.Bulma("example/example.json", "example/bulma.html")
	if err != nil {
		log.Println("It was not possible to parse json to html file\n", err)
		return
	}

	err = html.Custom("example/example.json", "example/custom.html", "example/test.json")
	if err != nil {
		log.Println("It was not possible to parse json to html file\n", err)
		return
	}

	err = markdown.Parser("./example/example.json", "example/example.md")
	if err != nil {
		log.Println("It was not possible to parse json to markdown file\n", err)
		return
	}

	open, err := os.Open("./example/example.json")
	if err != nil {
		log.Println("open file error\n", err)
		return
	}
	all, err := io.ReadAll(open)
	if err != nil {
		log.Println("io.ReadAll error\n", err)
		return
	}
	str, err := html.SampleStr(string(all))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)

	parserStr, err := markdown.ParserStr(string(all))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parserStr)
}
