package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/config"
)

func Bulma(jsonFilePath, outputFilePath string) (err error) {
	support.LoadStyleMap(config.BulmaJSONFile)
	return html.Parser(jsonFilePath, outputFilePath)
}
