package html

import (
	"github.com/zzhaoh/go-editorjs-parser/parser/html"
	"github.com/zzhaoh/go-editorjs-parser/support"
)

func Custom(jsonFilePath, outputFilePath, stylePath string) (err error) {
	support.LoadExternalStyleMap(stylePath)
	return html.Parser(jsonFilePath, outputFilePath, "custom")
}

func CustomStr(input, stylePath string) (string, error) {
	support.LoadExternalStyleMap(stylePath)
	return html.ParserStr(input, "custom")
}
