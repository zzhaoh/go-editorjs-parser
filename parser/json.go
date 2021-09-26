package parser

import (
	"encoding/json"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/domain"
	"log"
)

func ParseEditorJSON(editorJS string) domain.EditorJS {
	var result domain.EditorJS

	err := json.Unmarshal([]byte(editorJS), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
