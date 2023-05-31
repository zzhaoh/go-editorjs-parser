package html

import (
	"github.com/zzhaoh/go-editorjs-parser/parser/html"
	"github.com/zzhaoh/go-editorjs-parser/parser/html/bulma"
)

func Bulma(jsonFilePath, outputFilePath string) (err error) {
	return html.Parser(jsonFilePath, outputFilePath, bulma.StyleName)
}

func BulmaStr(input string) (string, error) {
	return html.ParserStr(input, bulma.StyleName)
}
