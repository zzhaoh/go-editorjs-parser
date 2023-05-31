package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/sample"
)

func Sample(jsonFilePath, outputFilePath string) (err error) {
	return html.Parser(jsonFilePath, outputFilePath, sample.StyleName)
}

func SampleStr(input string) (string, error) {
	return html.ParserStr(input, sample.StyleName)
}
