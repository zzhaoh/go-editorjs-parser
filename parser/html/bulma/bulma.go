package bulma

import (
	"fmt"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/parser/html/common"
	sup "gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/config"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/domain"
	"strings"
)

type Object struct {
	Data    interface{}
	Result  []string
	Styles  []string
	Scripts []string
}

const (
	StyleName  = "bulma"
	MapFile    = "bulma.json"
	ScriptFile = "bulma.js"
	ScriptType = "js"
)

func Init() (framework Object) {
	return framework
}

func (o *Object) SetData(data interface{}) {
	o.Data = data
}

func (o *Object) SetStyles(styles []string) {
	for _, style := range styles {
		o.Styles = append(o.Styles, style)
	}
}

func (o *Object) SetResult(result string) {
	o.Result = append(o.Result, result)
}

func (o *Object) SetScripts(scripts []string) {
	for _, script := range scripts {
		o.Scripts = append(o.Scripts, script)
	}
}

func (o *Object) LoadLibrary() {
	for _, l := range sup.SM.LibraryPaths {
		o.Styles = append(o.Styles, `<link rel="stylesheet" href="`+l+`">`)
	}

	o.Scripts = append(o.Scripts, string(sup.MinifyAsset(config.AssetsScriptPath+ScriptFile, ScriptType)))
}

func (o *Object) CreatePage() string {
	return common.CreatePage(o.Scripts, o.Styles, o.Result)
}

func (o *Object) Header() {
	obj := o.Data.(*domain.EditorJSDataHeader)
	o.Result = append(o.Result, fmt.Sprintf(`<div class="content">%s</div>`, common.Header(obj)))
}

func (o *Object) Paragraph() {
	obj := o.Data.(*domain.EditorJSDataParagraph)
	o.Result = append(o.Result, fmt.Sprintf(`<div class="content">%s</div>`, common.Paragraph(obj)))
}

func (o *Object) Quote() {
	obj := o.Data.(*domain.EditorJSDataQuote)
	var output []string

	output = append(output, `<div class="content">`,
		`<blockquote class="`+sup.SM.Quote.Blockquote+` `+sup.SM.Alignment[obj.Alignment]+`">`,
		obj.Text,
		`<p class="`+sup.SM.Quote.Author+`">`,
		obj.Caption,
		`</p>`,
		`</blockquote>`,
		`</div>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) Warning() {
	obj := o.Data.(*domain.EditorJSDataWarning)
	var output []string

	output = append(output, `<div class="`+sup.SM.Warning.Block+`">`)

	if sup.SM.Warning.CloseButton {
		output = append(output, `<button class="delete"></button>`)
	}

	output = append(output, `<span class="`+sup.SM.Warning.Title+`">`,
		obj.Title,
		`</span>`,
		obj.Message,
		`</div>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) Delimiter() {
	o.Result = append(o.Result, common.Delimiter())
}

func (o *Object) Alert() {
	obj := o.Data.(*domain.EditorJSDataAlert)
	var output []string

	output = append(output, `<div class="`+sup.SM.Alert.Block+` `+sup.SM.Alert.Types[obj.Type]+`">`)

	if sup.SM.Alert.CloseButton {
		output = append(output, `<button class="delete"></button>`)
	}

	output = append(output, obj.Message,
		`</div>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) List() {
	obj := o.Data.(*domain.EditorJSDataList)
	o.Result = append(o.Result, fmt.Sprintf(`<div class="content">%s</div>`, common.List(obj)))
}

func (o *Object) Checklist() {
	obj := o.Data.(*domain.EditorJSDataChecklist)
	o.Result = append(o.Result, common.Checklist(obj))
}

func (o *Object) Table() {
	obj := o.Data.(*domain.EditorJSDataTable)
	o.Result = append(o.Result, common.Table(obj))
}

func (o *Object) AnyButton() {
	obj := o.Data.(*domain.EditorJSDataAnyButton)
	o.Result = append(o.Result, common.AnyButton(obj))
}

func (o *Object) Code() {
	obj := o.Data.(*domain.EditorJSDataCode)
	o.Result = append(o.Result, common.Code(obj))
}

func (o *Object) Raw() {
	obj := o.Data.(*domain.EditorJSDataRaw)
	o.Result = append(o.Result, common.Raw(obj))
}

func (o *Object) Image() {
	obj := o.Data.(*domain.EditorJSDataImage)
	classes := ""
	classDiv := ""
	url := ""

	if obj.File.URL != "" {
		url = obj.File.URL
	} else {
		url = obj.URL
	}

	if obj.WithBorder {
		classes += sup.SM.Image.Border + " "
	}

	if obj.Stretched {
		classes += sup.SM.Image.Stretched
	}

	if obj.WithBackground {
		classDiv = sup.SM.Image.Background
	}

	o.Result = append(o.Result, fmt.Sprintf(`<figure class="%s %s" ><img class="%s %s" src="%s" alt="%s" title="%s" /></figure>`, sup.SM.Image.Block, classDiv, sup.SM.Image.Image, classes, url, obj.Caption, obj.Caption))
}

func (o *Object) LinkTool() {
	obj := o.Data.(*domain.EditorJSDataLinkTool)
	var output []string

	output = append(output, `<a href="`+obj.Link+`" target="_Blank" rel="nofollow noindex noreferrer" class="` + sup.SM.LinkTool.Link + `">`,
		`<div class="` + sup.SM.LinkTool.Container + `">`,
		`<div class="` + sup.SM.LinkTool.LeftColumn + `">`,
		`<div class="` + sup.SM.LinkTool.Title + `">`,
		obj.Meta.Title,
		`</div>`,
		`<div class="` + sup.SM.LinkTool.Description + `">`,
		obj.Meta.Description,
		`</div>`,
		`<div class="` + sup.SM.LinkTool.LinkDescription + `">`,
		strings.ReplaceAll(strings.ReplaceAll(obj.Link, "https://", ""), "http://", ""),
		`</div>`,
		`</div>`,
		`<div class="` + sup.SM.LinkTool.RightColumn + `">`,
		`<img class="` + sup.SM.LinkTool.Image + `" src="` + obj.Meta.Image.URL + `" />`,
		`</div>`,
		`</div>`,
		`</a>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) Attaches() {
	obj := o.Data.(*domain.EditorJSDataAttaches)
	var output []string

	output = append(output, `<a href="` + obj.File.URL + `" rel="noopener noreferrer" target="_blank" class="` + sup.SM.Attaches.Link + `">`,
		`<div class="` + sup.SM.Attaches.Container + `">`,
		`<div class="` + sup.SM.Attaches.LeftColumn + `" >`,
		`<img class="` + sup.SM.Attaches.LeftImage + `" src="https://i.ibb.co/K7Myr2k/file-icon.png" />`,
		`</div>`,
		`<div class="` + sup.SM.Attaches.CenterColumn + `">`,
		`<div class="` + sup.SM.Attaches.Filename + `">`,
		obj.File.Name,
		`</div>`,
		`<div class="` + sup.SM.Attaches.Size + `">`,
		sup.HumanFileSize(obj.File.Size),
		`</div>`,
		`</div>`,
		`<div class="` + sup.SM.Attaches.RightColumn + `" >`,
		`<img class="` + sup.SM.Attaches.RightImage + `" src="https://i.ibb.co/VYyHr6C/download-icon.png" />`,
		`</div>`,
		`</div>`,
		`</a>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) Embed() {
	obj := o.Data.(*domain.EditorJSDataEmbed)
	o.Result = append(o.Result, common.Embed(obj))
}

func (o *Object) ImageGallery() {
	obj := o.Data.(*domain.EditorJSDataImageGallery)
	r, s := common.ImageGallery(obj)

	o.Result = append(o.Result, r)
	o.Scripts = append(o.Scripts, sup.AppendBlockScript(s))
}
