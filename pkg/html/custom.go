package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
)

func Custom(jsonFilePath, outputFilePath, stylePath string) (err error) {
	support.LoadExternalStyleMap(stylePath)
	return html.Parser(jsonFilePath, outputFilePath)
}
