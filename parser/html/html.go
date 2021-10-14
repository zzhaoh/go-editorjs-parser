package html

import (
	"encoding/json"
	"fmt"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/domain"
	"log"
	"strconv"
	"strings"
)



func Header(el *domain.EditorJSDataHeader) string {
	anchor := ""
	if el.Anchor != "" {
		anchor = `id="` + strings.ToLower(strings.ReplaceAll(el.Anchor, " ", "-")) + `"`
	}

	tag := `h` + strconv.Itoa(el.Level)

	class := `class="` + support.SM.Header[tag] + `"`

	level := strconv.Itoa(el.Level)

	return fmt.Sprintf("<h%s %s %s>%s</h%s>", level, anchor, class, el.Text, level)
}

func Paragraph(el *domain.EditorJSDataParagraph) string {
	alignment := ""
	if el.Alignment != "" {
		alignment = ` class="` + support.SM.Paragraph + ` ` + support.SM.Alignment[el.Alignment] + `"`
	}

	return fmt.Sprintf("<p%s>%s</p>", alignment, el.Text)
}

func Quote(el *domain.EditorJSDataQuote) string {
	var output []string

	output = append(output, `<figure class="` + support.SM.Quote.Figure + ` ` + support.SM.Alignment[el.Alignment] + `">`,
		`<blockquote class="` + support.SM.Quote.Blockquote + `">`,
		el.Text,
		`</blockquote>`,
		`<figcaption class="` + support.SM.Quote.Figcaption + `">`,
		el.Caption,
		`</figcaption>`,
		`</figure>`)

	return strings.Join(output[:], "\n")
}

func Warning(el *domain.EditorJSDataWarning) string {
	var output []string

	output = append(output, `<div class="` + support.SM.Warning + `">`,
		`<b>`,
		el.Title,
		`</b>`,
		el.Message,
		`</div>`)

	return strings.Join(output[:], "\n")
}

func Delimiter() string {
	var output []string

	output = append(output, `<div class="` + support.SM.Delimiter + `">***</div>`)

	return strings.Join(output[:], "\n")
}

func Alert(el *domain.EditorJSDataAlert) string {
	var output []string

	output = append(output, `<div class="` + support.SM.Alert["box"] + ` ` + support.SM.Alert[el.Type] + `">`,
		el.Message,
		`</div>`)

	return strings.Join(output[:], "\n")
}

func List(el *domain.EditorJSDataList) string {
	var output []string

	listStyle := "ol"
	if el.Style == "unordered" {
		listStyle = "ul"
	}

	items, err := json.Marshal(el.Items)
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []domain.NestedListItem

	err = json.Unmarshal(items, &itemsList)
	if err == nil {
		output = append(output, support.CreateHTMLNestedList(itemsList, listStyle, true))
	} else {
		output = append(output, `<` + listStyle + ` class="` + support.SM.List.Group + `">`)

		for _, item := range el.Items {
			output = append(output, `<li class="` + support.SM.List.Item + `">` + fmt.Sprintf("%v", item) + `</li>`)
		}

		output = append(output, `</` + listStyle + `>`)
	}

	return strings.Join(output[:], "\n")
}

func Checklist(el *domain.EditorJSDataChecklist) string {
	var output []string

	items, err := json.Marshal(el.Items)
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []domain.ChecklistItem

	err = json.Unmarshal(items, &itemsList)
	if err == nil {
		output = append(output, `<div class="` + support.SM.Checklist.Block + `">`)

		for _, item := range itemsList {
			output = append(output, `<div class="` + support.SM.Checklist.Item + `">`)

			if item.Checked {
				output = append(output, `<span class="` + support.SM.Checklist.CheckboxChecked + `">&#10004;</span>`)
			} else {
				output = append(output, `<span class="` + support.SM.Checklist.CheckboxUnchecked + `">&nbsp;-&nbsp;</span>`)
			}

			output = append(output, `<span class="` + support.SM.Checklist.Text + `">` + item.Text + `</span>`,
				`</div>`)
		}

		output = append(output, `</div>`)
	}

	return strings.Join(output[:], "\n")
}

func Table(el *domain.EditorJSDataTable) string {
	var output []string

	output = append(output, `<table class="` + support.SM.Table.Table + `">`)

	for index, line := range el.Content {
		output = append(output, `<tr class="` + support.SM.Table.Row + `">`)

		tag := `td`
		tagClass := `class="` + support.SM.Table.CellTD + `"`
		if el.WithHeadings && index == 0 {
			tag = `th`
			tagClass = `class="` + support.SM.Table.CellTH + `"`
		}

		for _, info := range line {
			output = append(output, `<` + tag + ` ` + tagClass + `>` + info + `</` + tag + `>`)
		}

		output = append(output, `</tr>`)
	}

	output = append(output, `</table>`)

	return strings.Join(output[:], "\n")
}

func AnyButton(el *domain.EditorJSDataAnyButton) string {
	var output []string

	output = append(output, `<a class="` + support.SM.AnyButton + `" href="` + el.Link + `">` + el.Text + `</a>`)

	return strings.Join(output[:], "\n")
}

func Code(el *domain.EditorJSDataCode) string {
	var output []string

	output = append(output, `<pre class="` + support.SM.Code.Pre + `">`,
		`<code class="` + support.SM.Code.Code + `">` + el.Code,
		`</code></pre>`)

	return strings.Join(output[:], "\n")
}

func Raw(el *domain.EditorJSDataRaw) string {
	var output []string

	content := strings.ReplaceAll(el.Html,"<","&lt;")
	content = strings.ReplaceAll(content,">","&gt;")

	output = append(output, `<pre class="` + support.SM.Raw.Pre + `">`,
		`<code class="` + support.SM.Raw.Code + `">` + content,
		`</code></pre>`)

	return strings.Join(output[:], "\n")
}

func Image(el *domain.EditorJSDataImage) string {
	classes := ""
	classDiv := ""
	url := ""

	if el.File.URL != "" {
		url = el.File.URL
	} else {
		url = el.URL
	}

	if el.WithBorder {
		classes += support.SM.Image.Border + " "
	}

	if el.Stretched {
		classes += support.SM.Image.Stretched
	}

	if el.WithBackground {
		classDiv = support.SM.Image.Background
	}

	return fmt.Sprintf(`<div class="%s" ><img class="%s" src="%s" alt="%s" title="%s" /></div>`, classDiv, classes, url, el.Caption, el.Caption)
}

func LinkTool(el *domain.EditorJSDataLinkTool) string {
	var output []string

	output = append(output, `<a href="`+el.Link+`" target="_Blank" rel="nofollow noindex noreferrer" class="` + support.SM.LinkTool.Link + `">`,
		`<div class="` + support.SM.LinkTool.Container + `">`,
		`<div class="` + support.SM.LinkTool.Row + `">`,
		`<div class="` + support.SM.LinkTool.LeftColumn + `">`,
		`<div class="` + support.SM.LinkTool.Title + `">`,
		el.Meta.Title,
		`</div>`,
		`<div class="` + support.SM.LinkTool.Description + `">`,
		el.Meta.Description,
		`</div>`,
		`<div class="` + support.SM.LinkTool.LinkDescription + `">`,
		strings.ReplaceAll(strings.ReplaceAll(el.Link, "https://", ""), "http://", ""),
		`</div>`,
		`</div>`,
		`<div class="` + support.SM.LinkTool.RightColumn + `">`,
		`<img class="` + support.SM.LinkTool.Image + `" src="` + el.Meta.Image.URL + `" />`,
		`</div>`,
		`</div>`,
		`</div>`,
		`</a>`)

	return strings.Join(output[:], "\n")
}

func Attaches(el *domain.EditorJSDataAttaches) string {
	var output []string

	output = append(output, `<a href="` + el.File.URL + `" rel="noopener noreferrer" target="_blank" class="` + support.SM.Attaches.Link + `">`,
		`<div class="` + support.SM.Attaches.Container + `">`,
		`<div class="` + support.SM.Attaches.Row + `" >`,
		`<div class="` + support.SM.Attaches.LeftColumn + `" >`,
		`<img class="` + support.SM.Attaches.LeftImage + `" src="https://i.ibb.co/K7Myr2k/file-icon.png" />`,
		`</div>`,
		`<div class="` + support.SM.Attaches.CenterColumn + `">`,
		`<div class="` + support.SM.Attaches.Filename + `">`,
		el.File.Name,
		`</div>`,
		`<div class="` + support.SM.Attaches.Size + `">`,
		support.HumanFileSize(el.File.Size),
		`</div>`,
		`</div>`,
		`<div class="` + support.SM.Attaches.RightColumn + `" >`,
		`<img class="` + support.SM.Attaches.RightImage + `" src="https://i.ibb.co/VYyHr6C/download-icon.png" />`,
		`</div>`,
		`</div>`,
		`</div>`,
		`</a>`)

	return strings.Join(output[:], "\n")
}

func Embed(el *domain.EditorJSDataEmbed) string {
	var output []string

	output = append(output, `<div class="` + support.SM.Embed.Block + `" style="max-width: `+strconv.Itoa(el.Width)+`px">`,
		`<div class="` + support.SM.Embed.Title + `">` + el.Caption + `</div>`,
		`<iframe width="` + strconv.Itoa(el.Width) + `" height="` + strconv.Itoa(el.Height) + `" src="` + el.Embed + `" title="` + el.Caption + `" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`,
		`<div class="` + support.SM.Embed.Bottom + `">`,
		`<a class="` + support.SM.Embed.Link + `" href="`+el.Source+`" target="_Blank">Watch on `+el.Service+`</a>`,
		`</div>`,
		`</div>`)
	return strings.Join(output[:], "\n")
}

func ImageGallery(el *domain.EditorJSDataImageGallery) (string, string) {
	galleryScript := `
    const gg = new GalleryGrid();
    gg.loadGallery();
`
	galleryId := ""
	galleryClass := "gg-box"
	galleryHTML := `<div class="gg-container">`

	if el.BkgMode {
		galleryClass += " dark"
		galleryScript += `
	gg.galleryOptions({
	    selector: ".dark",
	    darkMode: true
	});`
	}

	if el.LayoutDefault {
		galleryId = ""
	} else if el.LayoutHorizontal {
		galleryId = "horizontal"
		galleryScript += `
	gg.galleryOptions({
		selector: "#horizontal",
		layout: "horizontal"
	});`
	} else if el.LayoutSquare {
		galleryId = "square"
		galleryScript += `
	gg.galleryOptions({
		selector: "#square",
		layout: "square"
	});`
	} else if el.LayoutWithGap {
		galleryId = "gap"
		galleryScript += `
	gg.galleryOptions({
		selector: "#gap",
		gapLength: 10
	});`
	} else if el.LayoutWithFixedSize {
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

	for index, url := range el.URLs {
		galleryHTML += fmt.Sprintf(`
<img src="%s" id="gg-image-%s" />`, url, strconv.Itoa(index))
	}

	galleryHTML += `
</div>
</div>`

	return galleryHTML, galleryScript
}
