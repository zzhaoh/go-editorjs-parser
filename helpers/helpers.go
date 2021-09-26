package helpers

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/domain"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//go:embed libs/*
var libsFiles embed.FS

func ReadJsonFile(jsonFilePath string) (jsonData string, err error) {
	data, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Println("Error reading the input json file\n",err)
	}

	jsonData = string(data)

	return
}

func MinifyLib(libPath, libType string) (contentMinified []byte) {
	contentFile, err := libsFiles.ReadFile(libPath)
	if err == nil {
		contentMinified, err = MinifyContent(contentFile,libType)
		if err != nil {
			log.Println("Error minifying internal lib file\n",err)
		}
	}

	return
}

func MinifyExternalStyle(libPath string) (contentMinified []byte, err error) {
	contentFile, err := ioutil.ReadFile(libPath)
	if err == nil {
		contentMinified, err = MinifyContent(contentFile,"css")
		if err != nil {
			log.Println("Error minifying external lib file\n",err)
		}
	}

	return
}

func MinifyContent(content []byte, format string) (contentMinified []byte, err error) {

	m := minify.New()

	switch format {

	case "css":
		m.AddFunc("text/css", css.Minify)
		contentMinified, err = m.Bytes("text/css", content)
		if err != nil {
			log.Println("Error minifying CSS file\n",err)
		}

	case "js":
		m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
		contentMinified, err = m.Bytes("application/javascript", content)
		if err != nil {
			log.Println("Error minifying JS file\n",err)
		}
	}

	return
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func CreateHTMLNestedList(items []domain.NestedListItem, listStyle string) string {
	var result []string

	result = append(result, "<"+listStyle+">")

	for _, item := range items {
		result = append(result, "<li>"+item.Content)

		if len(item.Items) > 0 {
			result = append(result, CreateHTMLNestedList(item.Items, listStyle))
		}

		result = append(result, "</li>")
	}

	result = append(result, "</"+listStyle+">")

	return strings.Join(result[:], "\n")
}

func CreateMarkDownNestedList(items []domain.NestedListItem, listStyle, spaceLeft string) string {
	var result []string

	for i, item := range items {
		if listStyle == "unordered" {
			result = append(result, spaceLeft+"- "+fmt.Sprintf("%v", item.Content))
		} else {
			n := spaceLeft + strconv.Itoa(i+1) + "."
			result = append(result, fmt.Sprintf("%s %s", n, item.Content))
		}

		if len(item.Items) > 0 {
			result = append(result, CreateMarkDownNestedList(item.Items, listStyle, spaceLeft+"    "))
		}

	}

	return strings.Join(result[:], "\n")
}

func PrepareData(el domain.EditorJSBlock) (data interface{}) {
	jsonData, err := json.Marshal(el.Data)
	if err != nil {
		log.Println("Error when trying to marshall EditorJS block data\n",err)
	}

	switch el.Type {
	case "header":
		data = new(domain.EditorJSDataHeader)
	case "paragraph":
		data = new(domain.EditorJSDataParagraph)
	case "quote":
		data = new(domain.EditorJSDataQuote)
	case "warning":
		data = new(domain.EditorJSDataWarning)
	case "alert":
		data = new(domain.EditorJSDataAlert)
	case "list":
		data = new(domain.EditorJSDataList)
	case "checklist":
		data = new(domain.EditorJSDataChecklist)
	case "table":
		data = new(domain.EditorJSDataTable)
	case "AnyButton":
		data = new(domain.EditorJSDataAnyButton)
	case "code":
		data = new(domain.EditorJSDataCode)
	case "raw":
		data = new(domain.EditorJSDataRaw)
	case "image":
		data = new(domain.EditorJSDataImage)
	case "linkTool":
		data = new(domain.EditorJSDataLinkTool)
	case "attaches":
		data = new(domain.EditorJSDataAttaches)
	case "embed":
		data = new(domain.EditorJSDataEmbed)
	case "imageGallery":
		data = new(domain.EditorJSDataImageGallery)

	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println("Error when trying to unmarshall EditorJS block data\n",err)
	}

	return
}

func Separator(size int) string {
	style := `width: 100%; height: ` + strconv.Itoa(size) + `px`
	return `<div style="` + style + `"></div>`
}

func HumanFileSize(fileSize float64) string {
	var sizePrefix string
	var formattedSize float64

	if math.Log10(+fileSize) >= 6 {
		sizePrefix = "MiB"
		formattedSize = fileSize / math.Pow(2, 20)
	} else {
		sizePrefix = "KiB"
		formattedSize = fileSize / math.Pow(2, 10)
	}

	return fmt.Sprintf("%v %v", toFixed(formattedSize, 1), sizePrefix)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func WriteOutputFile(outputFilePath, outputContent, outputType string) (err error) {
	outputPath := outputFilePath

	if !strings.HasSuffix(outputFilePath, "/") && outputFilePath != "" {
		outputPath = outputFilePath + "/"
	}

	if outputType == "html" {
		outputPath += "output.html"
	} else if outputType == "markdown" {
		outputPath += "output.md"
	}

	err = os.WriteFile(outputPath, []byte(outputContent), 0644)
	if err != nil {
		log.Println("Error writing the output file\n",err)
	}

	return
}
