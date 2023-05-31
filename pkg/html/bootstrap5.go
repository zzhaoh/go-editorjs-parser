package html

import (
	"github.com/zzhaoh/go-editorjs-parser/parser/html"
	"github.com/zzhaoh/go-editorjs-parser/parser/html/bootstrap"
)

func Bootstrap(jsonFilePath, outputFilePath string) (err error) {
	return html.Parser(jsonFilePath, outputFilePath, bootstrap.StyleName)
}

func BootstrapStr(input string) (string, error) {
	return html.ParserStr(input, bootstrap.StyleName)
}
