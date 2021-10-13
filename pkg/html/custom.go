package html

import "gitlab.com/rodrigoodhin/go-editorjs-parser/support"

func Custom(jsonFilePath, outputFilePath, stylePath string) (err error) {
	support.LoadExternalStyleMap(stylePath)
	return Default(jsonFilePath, outputFilePath)
}
