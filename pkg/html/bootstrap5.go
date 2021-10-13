package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/helpers"
)

func Bootstrap5(jsonFilePath, outputFilePath, styleToUse string) (err error) {
	helpers.LoadStyleMap("./support/styles/bootstrap5.json")
	return Default(jsonFilePath, outputFilePath, styleToUse)
}
