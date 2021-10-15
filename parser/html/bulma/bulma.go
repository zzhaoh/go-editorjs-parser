package bulma

import (
	"encoding/json"
	"fmt"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/domain"
	"log"
	"strconv"
	"strings"
)

type Object struct {
	Data    interface{}
}

func Init() (framework Object) {
	return framework
}

func (o *Object) SetData(data interface{})  {
	o.Data = data
}

func (o *Object) LoadLibrary() (styles []string){
	for _, l := range support.SM.LibraryPaths {
		styles = append(styles, `<link rel="stylesheet" href="`+l+`">`)
	}

	return
}

func (o *Object) Header() string {
	obj := o.Data.(*domain.EditorJSDataHeader)

	anchor := ""
	if obj.Anchor != "" {
		anchor = `id="` + strings.ToLower(strings.ReplaceAll(obj.Anchor, " ", "-")) + `"`
	}

	tag := `h` + strconv.Itoa(obj.Level)

	class := `class="` + support.SM.Header[tag] + `"`

	level := strconv.Itoa(obj.Level)

	return fmt.Sprintf("<h%s %s %s>%s</h%s>", level, anchor, class, obj.Text, level)
}

func (o *Object) Paragraph() string {
	obj := o.Data.(*domain.EditorJSDataParagraph)

	alignment := ""
	if obj.Alignment != "" {
		alignment = ` class="` + support.SM.Paragraph + ` ` + support.SM.Alignment[obj.Alignment] + `"`
	}

	return fmt.Sprintf("<p%s>%s</p>", alignment, obj.Text)
}

func (o *Object) Quote() string {
	obj := o.Data.(*domain.EditorJSDataQuote)
	var output []string

	output = append(output, `<figure class="`+support.SM.Quote.Figure+` `+support.SM.Alignment[obj.Alignment]+`">`,
		`<blockquote class="`+support.SM.Quote.Blockquote+`">`,
		obj.Text,
		`</blockquote>`,
		`<figcaption class="`+support.SM.Quote.Figcaption+`">`,
		obj.Caption,
		`</figcaption>`,
		`</figure>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) Warning() string {
	obj := o.Data.(*domain.EditorJSDataWarning)
	var output []string

	output = append(output, `<div class="`+support.SM.Warning+`">`,
		`<b>`,
		obj.Title,
		`</b>`,
		obj.Message,
		`</div>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) Delimiter() string {
	var output []string

	output = append(output, `<div class="`+support.SM.Delimiter+`">***</div>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) Alert() string {
	obj := o.Data.(*domain.EditorJSDataAlert)
	var output []string

	output = append(output, `<div class="`+support.SM.Alert["box"]+` `+support.SM.Alert[obj.Type]+`">`,
		obj.Message,
		`</div>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) List() string {
	obj := o.Data.(*domain.EditorJSDataList)
	var output []string

	listStyle := "ol"
	if obj.Style == "unordered" {
		listStyle = "ul"
	}

	items, err := json.Marshal(obj.Items)
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []domain.NestedListItem

	err = json.Unmarshal(items, &itemsList)
	if err == nil {
		output = append(output, support.CreateHTMLNestedList(itemsList, listStyle, true))
	} else {
		output = append(output, `<`+listStyle+` class="`+support.SM.List.Group+`">`)

		for _, item := range obj.Items {
			output = append(output, `<li class="`+support.SM.List.Item+`">`+fmt.Sprintf("%v", item)+`</li>`)
		}

		output = append(output, `</`+listStyle+`>`)
	}

	return strings.Join(output[:], "\n")
}

func (o *Object) Checklist() string {
	obj := o.Data.(*domain.EditorJSDataChecklist)
	var output []string

	items, err := json.Marshal(obj.Items)
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []domain.ChecklistItem

	err = json.Unmarshal(items, &itemsList)
	if err == nil {
		output = append(output, `<div class="`+support.SM.Checklist.Block+`">`)

		for _, item := range itemsList {
			output = append(output, `<div class="`+support.SM.Checklist.Item+`">`)

			if item.Checked {
				output = append(output, `<span class="`+support.SM.Checklist.CheckboxChecked+`">&#10004;</span>`)
			} else {
				output = append(output, `<span class="`+support.SM.Checklist.CheckboxUnchecked+`">&nbsp;-&nbsp;</span>`)
			}

			output = append(output, `<span class="`+support.SM.Checklist.Text+`">`+item.Text+`</span>`,
				`</div>`)
		}

		output = append(output, `</div>`)
	}

	return strings.Join(output[:], "\n")
}

func (o *Object) Table() string {
	obj := o.Data.(*domain.EditorJSDataTable)
	var output []string

	output = append(output, `<table class="`+support.SM.Table.Table+`">`)

	for index, line := range obj.Content {
		output = append(output, `<tr class="`+support.SM.Table.Row+`">`)

		tag := `td`
		tagClass := `class="` + support.SM.Table.CellTD + `"`
		if obj.WithHeadings && index == 0 {
			tag = `th`
			tagClass = `class="` + support.SM.Table.CellTH + `"`
		}

		for _, info := range line {
			output = append(output, `<`+tag+` `+tagClass+`>`+info+`</`+tag+`>`)
		}

		output = append(output, `</tr>`)
	}

	output = append(output, `</table>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) AnyButton() string {
	obj := o.Data.(*domain.EditorJSDataAnyButton)
	var output []string

	output = append(output, `<a class="`+support.SM.AnyButton+`" href="`+obj.Link+`">`+obj.Text+`</a>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) Code() string {
	obj := o.Data.(*domain.EditorJSDataCode)
	var output []string

	output = append(output, `<pre class="`+support.SM.Code.Pre+`">`,
		`<code class="`+support.SM.Code.Code+`">`+obj.Code,
		`</code></pre>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) Raw() string {
	obj := o.Data.(*domain.EditorJSDataRaw)
	var output []string

	content := strings.ReplaceAll(obj.Html, "<", "&lt;")
	content = strings.ReplaceAll(content, ">", "&gt;")

	output = append(output, `<pre class="`+support.SM.Raw.Pre+`">`,
		`<code class="`+support.SM.Raw.Code+`">`+content,
		`</code></pre>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) Image() string {
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
		classes += support.SM.Image.Border + " "
	}

	if obj.Stretched {
		classes += support.SM.Image.Stretched
	}

	if obj.WithBackground {
		classDiv = support.SM.Image.Background
	}

	return fmt.Sprintf(`<div class="%s" ><img class="%s" src="%s" alt="%s" title="%s" /></div>`, classDiv, classes, url, obj.Caption, obj.Caption)
}

func (o *Object) LinkTool() string {
	obj := o.Data.(*domain.EditorJSDataLinkTool)
	var output []string

	output = append(output, `<a href="`+obj.Link+`" target="_Blank" rel="nofollow noindex noreferrer" class="`+support.SM.LinkTool.Link+`">`,
		`<div class="`+support.SM.LinkTool.Container+`">`,
		`<div class="`+support.SM.LinkTool.Row+`">`,
		`<div class="`+support.SM.LinkTool.LeftColumn+`">`,
		`<div class="`+support.SM.LinkTool.Title+`">`,
		obj.Meta.Title,
		`</div>`,
		`<div class="`+support.SM.LinkTool.Description+`">`,
		obj.Meta.Description,
		`</div>`,
		`<div class="`+support.SM.LinkTool.LinkDescription+`">`,
		strings.ReplaceAll(strings.ReplaceAll(obj.Link, "https://", ""), "http://", ""),
		`</div>`,
		`</div>`,
		`<div class="`+support.SM.LinkTool.RightColumn+`">`,
		`<img class="`+support.SM.LinkTool.Image+`" src="`+obj.Meta.Image.URL+`" />`,
		`</div>`,
		`</div>`,
		`</div>`,
		`</a>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) Attaches() string {
	obj := o.Data.(*domain.EditorJSDataAttaches)
	var output []string

	output = append(output, `<a href="`+obj.File.URL+`" rel="noopener noreferrer" target="_blank" class="`+support.SM.Attaches.Link+`">`,
		`<div class="`+support.SM.Attaches.Container+`">`,
		`<div class="`+support.SM.Attaches.Row+`" >`,
		`<div class="`+support.SM.Attaches.LeftColumn+`" >`,
		`<img class="`+support.SM.Attaches.LeftImage+`" src="https://i.ibb.co/K7Myr2k/file-icon.png" />`,
		`</div>`,
		`<div class="`+support.SM.Attaches.CenterColumn+`">`,
		`<div class="`+support.SM.Attaches.Filename+`">`,
		obj.File.Name,
		`</div>`,
		`<div class="`+support.SM.Attaches.Size+`">`,
		support.HumanFileSize(obj.File.Size),
		`</div>`,
		`</div>`,
		`<div class="`+support.SM.Attaches.RightColumn+`" >`,
		`<img class="`+support.SM.Attaches.RightImage+`" src="https://i.ibb.co/VYyHr6C/download-icon.png" />`,
		`</div>`,
		`</div>`,
		`</div>`,
		`</a>`)

	return strings.Join(output[:], "\n")
}

func (o *Object) Embed() string {
	obj := o.Data.(*domain.EditorJSDataEmbed)
	var output []string

	output = append(output, `<div class="`+support.SM.Embed.Block+`" style="max-width: `+strconv.Itoa(obj.Width)+`px">`,
		`<div class="`+support.SM.Embed.Title+`">`+obj.Caption+`</div>`,
		`<iframe width="`+strconv.Itoa(obj.Width)+`" height="`+strconv.Itoa(obj.Height)+`" src="`+obj.Embed+`" title="`+obj.Caption+`" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`,
		`<div class="`+support.SM.Embed.Bottom+`">`,
		`<a class="`+support.SM.Embed.Link+`" href="`+obj.Source+`" target="_Blank">Watch on `+obj.Service+`</a>`,
		`</div>`,
		`</div>`)
	return strings.Join(output[:], "\n")
}

func (o *Object) ImageGallery() (string, string) {
	obj := o.Data.(*domain.EditorJSDataImageGallery)
	galleryScript := `
    const gg = new GalleryGrid();
    gg.loadGallery();
`
	galleryId := ""
	galleryClass := "gg-box"
	galleryHTML := `<div class="gg-container">`

	if obj.BkgMode {
		galleryClass += " dark"
		galleryScript += `
	gg.galleryOptions({
	    selector: ".dark",
	    darkMode: true
	});`
	}

	if obj.LayoutDefault {
		galleryId = ""
	} else if obj.LayoutHorizontal {
		galleryId = "horizontal"
		galleryScript += `
	gg.galleryOptions({
		selector: "#horizontal",
		layout: "horizontal"
	});`
	} else if obj.LayoutSquare {
		galleryId = "square"
		galleryScript += `
	gg.galleryOptions({
		selector: "#square",
		layout: "square"
	});`
	} else if obj.LayoutWithGap {
		galleryId = "gap"
		galleryScript += `
	gg.galleryOptions({
		selector: "#gap",
		gapLength: 10
	});`
	} else if obj.LayoutWithFixedSize {
		galleryId = "heightWidth"
		galleryScript += `
	gg.galleryOptions({
		selector: "#heightWidth",
		rowHeight: 180,
		columnWidth: 280
	});`
	}

	galleryHTML += fmt.Sprintf(`
<div class="%s" id="%s">`, galleryClass, galleryId)

	for index, url := range obj.URLs {
		galleryHTML += fmt.Sprintf(`
<img src="%s" id="gg-image-%s" />`, url, strconv.Itoa(index))
	}

	galleryHTML += `
</div>
</div>`

	return galleryHTML, galleryScript
}

