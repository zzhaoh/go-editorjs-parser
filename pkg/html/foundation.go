package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/config"
)

func Foundation(jsonFilePath, outputFilePath string) (err error) {
	support.LoadStyleMap(config.FoundationJSONFile)
	return html.Parser(jsonFilePath, outputFilePath)
}
