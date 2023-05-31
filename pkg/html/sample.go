package html

import (
	"github.com/zzhaoh/go-editorjs-parser/parser/html"
	"github.com/zzhaoh/go-editorjs-parser/parser/html/sample"
)

func Sample(jsonFilePath, outputFilePath string) (err error) {
	return html.Parser(jsonFilePath, outputFilePath, sample.StyleName)
}

func SampleStr(input string) (string, error) {
	return html.ParserStr(input, sample.StyleName)
}
