package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/domain"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/helpers"
	"log"
	"reflect"
	"strings"
)

func Default(jsonFilePath, outputFilePath, styleToUse string) (err error) {
	var libs []string
	var result []string
	var blockScripts []string
	var styles []string
	var scripts []string
	var style string

	if reflect.DeepEqual(helpers.SM,domain.StyleMap{}) {
		helpers.LoadStyleMap("./support/styles/default.json")
	}

	input, err := helpers.ReadJsonFile(jsonFilePath)
	if err != nil {
		log.Println("It was not possible to read the input json file\n",err)
	}

	editorJSON := helpers.ParseEditorJSON(input)

	for _, el := range editorJSON.Blocks {

		_, found := helpers.Find(libs, el.Type)
		if !found {
			libs = append(libs, strings.ToLower(el.Type))
		}

		result = append(result, helpers.Separator(helpers.SM.SpaceBetweenBlocks))

		content := helpers.PrepareData(el)

		switch el.Type {

		case "header":
			result = append(result, html.Header(content.(*domain.EditorJSDataHeader)))
		case "paragraph":
			result =
				append(result, html.Paragraph(content.(*domain.EditorJSDataParagraph)))
		case "quote":
			result = append(result, html.Quote(content.(*domain.EditorJSDataQuote)))
		case "warning":
			result = append(result, html.Warning(content.(*domain.EditorJSDataWarning)))
		case "delimiter":
			result = append(result, html.Delimiter())
		case "alert":
			result = append(result, html.Alert(content.(*domain.EditorJSDataAlert)))
		case "list":
			result = append(result, html.List(content.(*domain.EditorJSDataList)))
		case "checklist":
			result = append(result, html.Checklist(content.(*domain.EditorJSDataChecklist)))
		case "table":
			result = append(result, html.Table(content.(*domain.EditorJSDataTable)))
		case "AnyButton":
			result = append(result, html.AnyButton(content.(*domain.EditorJSDataAnyButton)))
		case "code":
			result = append(result, html.Code(content.(*domain.EditorJSDataCode)))
		case "raw":
			result = append(result, html.Raw(content.(*domain.EditorJSDataRaw)))
		case "image":
			result = append(result, html.Image(content.(*domain.EditorJSDataImage)))
		case "linkTool":
			result = append(result, html.LinkTool(content.(*domain.EditorJSDataLinkTool)))
		case "attaches":
			result = append(result, html.Attaches(content.(*domain.EditorJSDataAttaches)))
		case "embed":
			result = append(result, html.Embed(content.(*domain.EditorJSDataEmbed)))
		case "imageGallery":
			imgRes, imgScript := html.ImageGallery(content.(*domain.EditorJSDataImageGallery))
			result = append(result, imgRes)
			blockScripts = append(blockScripts, imgScript)
		}

	}

	result = append(result, helpers.Separator(helpers.SM.SpaceBetweenBlocks))

	if helpers.SM.StyleName == "default" {
		styles = append(styles, string(helpers.MinifyLib("default.css", "css")))
	} else if helpers.SM.StyleName == "bootstrap5" {
		styles = append(styles, string(helpers.LoadLib("bootstrap5.min.css")))
	} else {
		styleMinified, err := helpers.MinifyExternalStyle(styleToUse)
		if string(styleMinified) != "" && err == nil {
			styles = append(styles, string(styleMinified))
		} else {
			log.Println("It was not possible to read the external style file\n",err)
		}
	}

	style = "\n<style>\n" + strings.Join(styles[:], "\n") + "\n</style>\n\n"

	for _, lib := range libs {
		scriptMinified := string(helpers.MinifyLib(lib + "/" + lib + ".js", "js"))
		if scriptMinified != "" {
			scripts = append(scripts, scriptMinified)
		}
	}

	for _, blockScript := range blockScripts {
		blockScriptMinified, _ := helpers.MinifyContent([]byte(blockScript), "js")
		blockScriptMinifiedStr := string(blockScriptMinified)
		if blockScriptMinifiedStr != "" {
			scripts = append(scripts, blockScriptMinifiedStr)
		}
	}

	script := "\n\n<script>\n" + strings.Join(scripts[:], "\n") + "\n</script>\n\n"

	content := style + strings.Join(result[:], "\n\n") + script

	err = helpers.WriteOutputFile(outputFilePath, content, "html")
	if err != nil {
		log.Println("It was not possible to write the output html file\n",err)
	}

	return
}
