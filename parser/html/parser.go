package html

import (
	"fmt"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/config"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/domain"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	result  []string
	styles  []string
	scripts []string
)

func Parser(jsonFilePath, outputFilePath string) (err error) {
	if reflect.DeepEqual(support.SM, domain.StyleMap{}) {
		log.Fatal("Style map is empty\n", err)
	}

	switch support.SM.StyleName {
	case config.DefaultStyleName:
		loadDefaultLibraryPaths()
	case config.Bootstrap5StyleName, config.FoundationStyleName:
		loadLibraryPaths()
	default:
		loadLibraryPaths()
	}

	input, err := support.ReadJsonFile(jsonFilePath)
	if err != nil {
		log.Println("It was not possible to read the input json file\n", err)
	}

	editorJSON := support.ParseEditorJSON(input)

	for _, el := range editorJSON.Blocks {

		appendLibs(el)

		result = append(result, support.Separator(support.SM.SpaceBetweenBlocks))

		content := support.PrepareData(el)

		switch el.Type {

		case "header":
			result = append(result, Header(content.(*domain.EditorJSDataHeader)))
		case "paragraph":
			result = append(result, Paragraph(content.(*domain.EditorJSDataParagraph)))
		case "quote":
			result = append(result, Quote(content.(*domain.EditorJSDataQuote)))
		case "warning":
			result = append(result, Warning(content.(*domain.EditorJSDataWarning)))
		case "delimiter":
			result = append(result, Delimiter())
		case "alert":
			result = append(result, Alert(content.(*domain.EditorJSDataAlert)))
		case "list":
			result = append(result, List(content.(*domain.EditorJSDataList)))
		case "checklist":
			result = append(result, Checklist(content.(*domain.EditorJSDataChecklist)))
		case "table":
			result = append(result, Table(content.(*domain.EditorJSDataTable)))
		case "AnyButton":
			result = append(result, AnyButton(content.(*domain.EditorJSDataAnyButton)))
		case "code":
			result = append(result, Code(content.(*domain.EditorJSDataCode)))
		case "raw":
			result = append(result, Raw(content.(*domain.EditorJSDataRaw)))
		case "image":
			result = append(result, Image(content.(*domain.EditorJSDataImage)))
		case "linkTool":
			result = append(result, LinkTool(content.(*domain.EditorJSDataLinkTool)))
		case "attaches":
			result = append(result, Attaches(content.(*domain.EditorJSDataAttaches)))
		case "embed":
			result = append(result, Embed(content.(*domain.EditorJSDataEmbed)))
		case "imageGallery":
			imgRes, imgScript := ImageGallery(content.(*domain.EditorJSDataImageGallery))
			result = append(result, imgRes)
			appendBlockScript(imgScript)
		}

	}

	result = append(result, support.Separator(support.SM.SpaceBetweenBlocks))

	script := "\n\n<script>\n" + strings.Join(scripts[:], "\n") + "\n</script>\n\n"

	content := strings.Join(styles[:], "\n") + strings.Join(result[:], "\n\n") + script

	err = support.WriteOutputFile(outputFilePath, content, "html")
	if err != nil {
		log.Println("It was not possible to write the output html file\n", err)
	}

	return
}

func appendLibs(block domain.EditorJSBlock) {
	libName := strings.ToLower(block.Type)
	libPath := config.LibsPath + libName + "/"
	if _, err := os.Stat(libPath); !os.IsNotExist(err) {
		styleMinified := string(support.MinifyLib(libName+"/"+libName+".css", "css"))
		if styleMinified != "" {
			styles = append(styles, `<style>` + styleMinified + `</style>`)
		}

		scriptMinified := string(support.MinifyLib(libName+"/"+libName+".js", "js"))
		if scriptMinified != "" {
			scripts = append(scripts, scriptMinified)
		}
	}
}

func appendBlockScript(blockScript string) {
	blockScriptMinified, _ := support.MinifyContent([]byte(blockScript), "js")
	blockScriptMinifiedStr := string(blockScriptMinified)
	if blockScriptMinifiedStr != "" {
		scripts = append(scripts, blockScriptMinifiedStr)
	}
}

func loadLibraryPaths(){
	for _, l := range support.SM.LibraryPaths {
		styles = append(styles, `<link rel="stylesheet" href="`+l+`">`)
	}
}

func loadDefaultLibraryPaths(){
	for _, l := range support.SM.LibraryPaths {
		fmt.Println(l)
		styles = append(styles, `<style>` + string(support.LoadStyle(l, "css")) + `</style>`)
	}
}