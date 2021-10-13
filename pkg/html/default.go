package html

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/config"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/domain"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	result       []string
	styles       []string
	scripts      []string
	style        string
)

func Default(jsonFilePath, outputFilePath string) (err error) {
	if reflect.DeepEqual(support.SM, domain.StyleMap{}) {
		support.LoadStyleMap(config.DefaultJSONFile)
	}

	if support.SM.StyleName == config.DefaultStyleName {
		styles = append(styles, string(support.MinifyStyle(config.DefaultCSSFile, "css")))
	} else if support.SM.StyleName == config.Bootstrap5StyleName {
		styles = append(styles, string(support.LoadStyle(config.Bootstrap5CSSFile)))
	} else {
		styleMinified, err := support.MinifyExternalStyle(support.SM.LibraryPath)
		if string(styleMinified) != "" && err == nil {
			styles = append(styles, string(styleMinified))
		} else {
			log.Println("It was not possible to read the external style file\n", err)
		}
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
			result = append(result, html.Header(content.(*domain.EditorJSDataHeader)))
		case "paragraph":
			result = append(result, html.Paragraph(content.(*domain.EditorJSDataParagraph)))
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
			appendBlockScript(imgScript)
		}

	}

	result = append(result, support.Separator(support.SM.SpaceBetweenBlocks))

	style = "\n<style>\n" + strings.Join(styles[:], "\n") + "\n</style>\n\n"
	script := "\n\n<script>\n" + strings.Join(scripts[:], "\n") + "\n</script>\n\n"

	content := style + strings.Join(result[:], "\n\n") + script

	err = support.WriteOutputFile(outputFilePath, content, "html")
	if err != nil {
		log.Println("It was not possible to write the output html file\n", err)
	}

	return
}

func appendLibs(block domain.EditorJSBlock) {
	libName := strings.ToLower(block.Type)
	libPath := config.LibsPath +libName + "/"
	if _, err := os.Stat(libPath); !os.IsNotExist(err) {
		styleMinified := string(support.MinifyLib(libName + "/" + libName + ".css", "css"))
		if styleMinified != "" {
			styles = append(styles, styleMinified)
		}

		scriptMinified := string(support.MinifyLib(libName + "/" + libName + ".js", "js"))
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
