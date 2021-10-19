package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/bulma"
)

func Bulma(jsonFilePath, outputFilePath string) (err error) {
	return html.Parser(jsonFilePath, outputFilePath, bulma.StyleName)
}
