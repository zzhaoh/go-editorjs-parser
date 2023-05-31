package html

import (
	"gitlab.com/zzhaoh/go-editorjs-parser/parser/html/bootstrap"
	"gitlab.com/zzhaoh/go-editorjs-parser/parser/html/bulma"
	"gitlab.com/zzhaoh/go-editorjs-parser/parser/html/sample"
	sup "gitlab.com/zzhaoh/go-editorjs-parser/support"
	"gitlab.com/zzhaoh/go-editorjs-parser/support/config"
	"gitlab.com/zzhaoh/go-editorjs-parser/support/domain"
	"log"
	"os"
	"reflect"
	"strings"
)

func Parser(jsonFilePath, outputFilePath, styleName string) error {

	input, err := sup.ReadJsonFile(jsonFilePath)
	if err != nil {
		log.Println("It was not possible to read the input json file\n", err)
	}

	output, err := parser1(input, styleName)
	if err != nil {
		log.Println("It was not possible to write the output html file\n", err)
		return err
	}

	err = sup.WriteOutputFile(outputFilePath, output, "html")
	if err != nil {
		log.Println("It was not possible to write the output html file\n", err)
		return err
	}

	return nil
}

func ParserStr(input, styleName string) (string, error) {
	output, err := parser1(input, styleName)
	if err != nil {
		log.Println("It was not possible to write the output html file\n", err)
		return "", err
	}

	return output, nil
}

func parser1(input, styleName string) (output string, err error) {

	useDefault := true
	if styleName == "custom" {
		useDefault = false
		styleName = sup.SM.StyleName
	}

	var f domain.EditorJSMethods
	switch styleName {
	case sample.StyleName:
		samplePkg := sample.Init(useDefault)
		f = &samplePkg
	case bootstrap.StyleName:
		bootstrapPkg := bootstrap.Init(useDefault)
		f = &bootstrapPkg
	case bulma.StyleName:
		bulmaPkg := bulma.Init(useDefault)
		f = &bulmaPkg
	}

	if reflect.DeepEqual(sup.SM, domain.StyleMap{}) {
		log.Fatal("Style map is empty\n", err)
	}

	if !sup.IsValidStyle(sup.SM.StyleName) {
		log.Fatal("Invalid style name: "+sup.SM.StyleName+"\n", err)
	}

	f.LoadLibrary()

	editorJSON := sup.ParseEditorJSON(input)

	for _, el := range editorJSON.Blocks {

		styles, scripts := appendLibs(el)
		f.SetStyles(styles)
		f.SetScripts(scripts)
		f.Separator()
		f.SetData(sup.PrepareData(el))

		switch el.Type {

		case "header":
			f.Header()
		case "paragraph":
			f.Paragraph()
		case "quote":
			f.Quote()
		case "warning":
			f.Warning()
		case "delimiter":
			f.Delimiter()
		case "alert":
			f.Alert()
		case "list":
			f.List()
		case "checklist":
			f.Checklist()
		case "table":
			f.Table()
		case "AnyButton":
			f.AnyButton()
		case "code":
			f.Code()
		case "raw":
			f.Raw()
		case "image":
			f.Image()
		case "linkTool":
			f.LinkTool()
		case "attaches":
			f.Attaches()
		case "embed":
			f.Embed()
		case "imageGallery":
			f.ImageGallery()
		}

	}

	f.Separator()
	output = f.CreatePage()

	return
}

func appendLibs(block domain.EditorJSBlock) (styles []string, scripts []string) {
	libName := strings.ToLower(block.Type)
	libPath := config.LibsPath + libName + "/"
	if _, err := os.Stat(libPath); !os.IsNotExist(err) {
		styleMinified := string(sup.MinifyLib(libName+"/"+libName+".css", "css"))
		if styleMinified != "" {
			styles = append(styles, `<style>`+styleMinified+`</style>`)
		}

		scriptMinified := string(sup.MinifyLib(libName+"/"+libName+".js", "js"))
		if scriptMinified != "" {
			scripts = append(scripts, scriptMinified)
		}
	}

	return
}
