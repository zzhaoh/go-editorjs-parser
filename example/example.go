package main

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/pkg/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/pkg/markdown"
	"log"
)

func main() {

	err := html.Default("./example/example.json", "","default")
	if err != nil {
		log.Println("It was not possible to parse json to html file\n",err)
		return
	}

	err = markdown.Parser("./example/example.json", "")
	if err != nil {
		log.Println("It was not possible to parse json to markdown file\n",err)
		return
	}

}