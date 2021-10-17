package main

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/pkg/html"
	"log"
)

func main() {

	//err := html.Sample("example/example.json", "example/sample.html")
	//if err != nil {
	//	log.Println("It was not possible to parse json to html file\n",err)
	//	return
	//}
//
	//err = html.Bootstrap5("example/example.json", "example/bootstrap5.html")
	//if err != nil {
	//	log.Println("It was not possible to parse json to html file\n",err)
	//	return
	//}
//
	//err = html.Bulma("example/example.json", "example/bulma.html")
	//if err != nil {
	//	log.Println("It was not possible to parse json to html file\n",err)
	//	return
	//}

	err := html.Custom("example/example.json", "example/custom.html", "example/test.json")
	if err != nil {
		log.Println("It was not possible to parse json to html file\n",err)
		return
	}

	//err = markdown.Parser("./example/example.json", "example")
	//if err != nil {
	//	log.Println("It was not possible to parse json to markdown file\n",err)
	//	return
	//}

}