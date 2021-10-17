package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/sample"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
)

func Sample(jsonFilePath, outputFilePath string) (err error) {
	support.LoadStyleMap(sample.MapFile)
	return html.Parser(jsonFilePath, outputFilePath)
}
