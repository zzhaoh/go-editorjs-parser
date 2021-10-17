package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/bulma"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
)

func Bulma(jsonFilePath, outputFilePath string) (err error) {
	support.LoadStyleMap(bulma.MapFile)
	return html.Parser(jsonFilePath, outputFilePath)
}
