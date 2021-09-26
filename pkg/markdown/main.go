package markdown

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/domain"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/helpers"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/markdown"
	"log"
	"strings"
)

func Parser(jsonFilePath, outputFilePath string) (err error) {
	var result []string

	input, err := helpers.ReadJsonFile(jsonFilePath)
	if err != nil {
		log.Println("It was not possible to read the input json file\n",err)
	}

	editorJSAST := parser.ParseEditorJSON(input)

	for _, el := range editorJSAST.Blocks {

		content := helpers.PrepareData(el)

		switch el.Type {

		case "header":
			result = append(result, markdown.Header(content.(*domain.EditorJSDataHeader)))
		case "paragraph":
			result = append(result, markdown.Paragraph(content.(*domain.EditorJSDataParagraph)))
		case "quote":
			result = append(result, markdown.Quote(content.(*domain.EditorJSDataQuote)))
		case "warning":
			result = append(result, markdown.Warning(content.(*domain.EditorJSDataWarning)))
		case "delimiter":
			result = append(result, markdown.Delimiter())
		case "alert":
			result = append(result, markdown.Alert(content.(*domain.EditorJSDataAlert)))
		case "list":
			result = append(result, markdown.List(content.(*domain.EditorJSDataList)))
		case "checklist":
			result = append(result, markdown.Checklist(content.(*domain.EditorJSDataChecklist)))
		case "table":
			result = append(result, markdown.Table(content.(*domain.EditorJSDataTable)))
		case "AnyButton":
			result = append(result, markdown.AnyButton(content.(*domain.EditorJSDataAnyButton)))
		case "code":
			result = append(result, markdown.Code(content.(*domain.EditorJSDataCode)))
		case "raw":
			result = append(result, markdown.Raw(content.(*domain.EditorJSDataRaw)))
		case "image":
			result = append(result, markdown.Image(content.(*domain.EditorJSDataImage)))
		case "linkTool":
			result = append(result, markdown.LinkTool(content.(*domain.EditorJSDataLinkTool)))
		case "attaches":
			result = append(result, markdown.Attaches(content.(*domain.EditorJSDataAttaches)))
		case "embed":
			result = append(result, markdown.Embed(content.(*domain.EditorJSDataEmbed)))
		case "imageGallery":
			result = append(result, markdown.ImageGallery(content.(*domain.EditorJSDataImageGallery)))
		}

	}

	content := strings.Join(result[:], "\n\n")

	err = helpers.WriteOutputFile(outputFilePath, content, "markdown")
	if err != nil {
		log.Println("It was not possible to write the output markdown file\n",err)
	}

	return
}
