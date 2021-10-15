package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/bootstrap5"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/bulma"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/sample"
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

	frameworks := make(map[string]domain.EditorJSMethods)
	samplePkg := sample.Init()
	frameworks[config.SampleStyleName] = &samplePkg

	bootstrap5Pkg := bootstrap5.Init()
	frameworks[config.Bootstrap5StyleName] = &bootstrap5Pkg

	bulmaPkg := bulma.Init()
	frameworks[config.BulmaStyleName] = &bulmaPkg

	styles = frameworks[support.SM.StyleName].LoadLibrary()

	input, err := support.ReadJsonFile(jsonFilePath)
	if err != nil {
		log.Println("It was not possible to read the input json file\n", err)
	}

	editorJSON := support.ParseEditorJSON(input)

	for _, el := range editorJSON.Blocks {

		appendLibs(el)

		result = append(result, support.Separator(support.SM.SpaceBetweenBlocks))

		frameworks[support.SM.StyleName].SetData(support.PrepareData(el))

		switch el.Type {

		case "header":
			result = append(result, frameworks[support.SM.StyleName].Header())
		case "paragraph":
			result = append(result, frameworks[support.SM.StyleName].Paragraph())
		case "quote":
			result = append(result, frameworks[support.SM.StyleName].Quote())
		case "warning":
			result = append(result, frameworks[support.SM.StyleName].Warning())
		case "delimiter":
			result = append(result, frameworks[support.SM.StyleName].Delimiter())
		case "alert":
			result = append(result, frameworks[support.SM.StyleName].Alert())
		case "list":
			result = append(result, frameworks[support.SM.StyleName].List())
		case "checklist":
			result = append(result, frameworks[support.SM.StyleName].Checklist())
		case "table":
			result = append(result, frameworks[support.SM.StyleName].Table())
		case "AnyButton":
			result = append(result, frameworks[support.SM.StyleName].AnyButton())
		case "code":
			result = append(result, frameworks[support.SM.StyleName].Code())
		case "raw":
			result = append(result, frameworks[support.SM.StyleName].Raw())
		case "image":
			result = append(result, frameworks[support.SM.StyleName].Image())
		case "linkTool":
			result = append(result, frameworks[support.SM.StyleName].LinkTool())
		case "attaches":
			result = append(result, frameworks[support.SM.StyleName].Attaches())
		case "embed":
			result = append(result, frameworks[support.SM.StyleName].Embed())
		case "imageGallery":
			imgRes, imgScript := frameworks[support.SM.StyleName].ImageGallery()
			result = append(result, imgRes)
			appendBlockScript(imgScript)
		}

	}

	result = append(result, support.Separator(support.SM.SpaceBetweenBlocks))

	err = support.WriteOutputFile(outputFilePath, createPageStructure(), "html")
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

func createPageStructure() string {
	script := "\n\n<script>\n" + strings.Join(scripts[:], "\n") + "\n</script>\n\n"

	return `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
` + strings.Join(styles[:], "\n") + ` 
  </head>
  <body>
` + strings.Join(result[:], "\n\n") + script + ` 
  </body>
</html>
`
}