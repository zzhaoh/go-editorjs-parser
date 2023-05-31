package html

import (
	"gitlab.com/zzhaoh/go-editorjs-parser/parser/html"
	"gitlab.com/zzhaoh/go-editorjs-parser/parser/html/bootstrap"
)

func Bootstrap(jsonFilePath, outputFilePath string) (err error) {
	return html.Parser(jsonFilePath, outputFilePath, bootstrap.StyleName)
}

func BootstrapStr(input string) (string, error) {
	return html.ParserStr(input, bootstrap.StyleName)
}
