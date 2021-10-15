package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/config"
)

func Sample(jsonFilePath, outputFilePath string) (err error) {
	support.LoadStyleMap(config.SampleJSONFile)
	return html.Parser(jsonFilePath, outputFilePath)
}
