package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/bootstrap5"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/bulma"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/sample"
	sup "gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/config"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/domain"
	"log"
	"os"
	"reflect"
	"strings"
)

func Parser(jsonFilePath, outputFilePath string) (err error) {
	if reflect.DeepEqual(sup.SM, domain.StyleMap{}) {
		log.Fatal("Style map is empty\n", err)
	}

	if !sup.IsValidStyle(sup.SM.StyleName) {
		log.Fatal("Invalid style name: "+sup.SM.StyleName+"\n", err)
	}

	f := make(map[string]domain.EditorJSMethods)
	samplePkg := sample.Init()
	f[sample.StyleName] = &samplePkg

	bootstrap5Pkg := bootstrap5.Init()
	f[bootstrap5.StyleName] = &bootstrap5Pkg

	bulmaPkg := bulma.Init()
	f[bulma.StyleName] = &bulmaPkg

	f[sup.SM.StyleName].LoadLibrary()

	input, err := sup.ReadJsonFile(jsonFilePath)
	if err != nil {
		log.Println("It was not possible to read the input json file\n", err)
	}

	editorJSON := sup.ParseEditorJSON(input)

	for _, el := range editorJSON.Blocks {

		styles, scripts := appendLibs(el)
		f[sup.SM.StyleName].SetStyles(styles)
		f[sup.SM.StyleName].SetScripts(scripts)

		f[sup.SM.StyleName].SetResult(sup.Separator(sup.SM.SpaceBetweenBlocks))

		f[sup.SM.StyleName].SetData(sup.PrepareData(el))

		switch el.Type {

		case "header":
			f[sup.SM.StyleName].Header()
		case "paragraph":
			f[sup.SM.StyleName].Paragraph()
		case "quote":
			f[sup.SM.StyleName].Quote()
		case "warning":
			f[sup.SM.StyleName].Warning()
		case "delimiter":
			f[sup.SM.StyleName].Delimiter()
		case "alert":
			f[sup.SM.StyleName].Alert()
		case "list":
			f[sup.SM.StyleName].List()
		case "checklist":
			f[sup.SM.StyleName].Checklist()
		case "table":
			f[sup.SM.StyleName].Table()
		case "AnyButton":
			f[sup.SM.StyleName].AnyButton()
		case "code":
			f[sup.SM.StyleName].Code()
		case "raw":
			f[sup.SM.StyleName].Raw()
		case "image":
			f[sup.SM.StyleName].Image()
		case "linkTool":
			f[sup.SM.StyleName].LinkTool()
		case "attaches":
			f[sup.SM.StyleName].Attaches()
		case "embed":
			f[sup.SM.StyleName].Embed()
		case "imageGallery":
			f[sup.SM.StyleName].ImageGallery()
		}

	}

	f[sup.SM.StyleName].SetResult(sup.Separator(sup.SM.SpaceBetweenBlocks))

	err = sup.WriteOutputFile(outputFilePath, f[sup.SM.StyleName].CreatePage(), "html")
	if err != nil {
		log.Println("It was not possible to write the output html file\n", err)
	}

	return
}

func appendLibs(block domain.EditorJSBlock) (styles []string, scripts []string){
	libName := strings.ToLower(block.Type)
	libPath := config.LibsPath + libName + "/"
	if _, err := os.Stat(libPath); !os.IsNotExist(err) {
		styleMinified := string(sup.MinifyLib(libName+"/"+libName+".css", "css"))
		if styleMinified != "" {
			styles = append(styles, `<style>` + styleMinified + `</style>`)
		}

		scriptMinified := string(sup.MinifyLib(libName+"/"+libName+".js", "js"))
		if scriptMinified != "" {
			scripts = append(scripts, scriptMinified)
		}
	}

	return
}