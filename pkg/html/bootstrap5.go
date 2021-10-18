package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/bootstrap"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
)

func Bootstrap(jsonFilePath, outputFilePath string) (err error) {
	support.LoadStyleMap(bootstrap.MapFile)
	return html.Parser(jsonFilePath, outputFilePath)
}
