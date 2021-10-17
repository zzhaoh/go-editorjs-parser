package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/bootstrap5"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
)

func Bootstrap5(jsonFilePath, outputFilePath string) (err error) {
	support.LoadStyleMap(bootstrap5.MapFile)
	return html.Parser(jsonFilePath, outputFilePath)
}
