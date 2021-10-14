package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/config"
)

func Bootstrap5(jsonFilePath, outputFilePath string) (err error) {
	support.LoadStyleMap(config.Bootstrap5JSONFile)
	return html.Parser(jsonFilePath, outputFilePath)
}
