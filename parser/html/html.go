package html

import (
	"encoding/json"
	"fmt"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/domain"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/helpers"
	"log"
	"strconv"
	"strings"
)

func Header(el *domain.EditorJSDataHeader) string {
	anchor := ""
	if el.Anchor != "" {
		anchor = ` id="` + strings.ToLower(strings.ReplaceAll(el.Anchor, " ", "-")) + `"`
	}

	level := strconv.Itoa(el.Level)

	return fmt.Sprintf("<h%s%s>%s</h%s>", level, anchor, el.Text, level)
}

func Paragraph(el *domain.EditorJSDataParagraph) string {
	alignment := ""
	if el.Alignment != "" {
		alignment = ` style="text-align:` + el.Alignment + `"`
	}

	return fmt.Sprintf("<p%s>%s</p>", alignment, el.Text)
}

func Quote(el *domain.EditorJSDataQuote) string {
	var result []string

	result = append(result, `<figure class="quote">`,
		`<blockquote style="text-align: `+el.Alignment+`">`,
		el.Text,
		`</blockquote>`,
		`<figcaption style="text-align: `+el.Alignment+`">`,
		`&mdash; `+el.Caption,
		`</figcaption>`,
		`</figure>`)

	return strings.Join(result[:], "\n")
}

func Warning(el *domain.EditorJSDataWarning) string {
	var result []string

	result = append(result, `<div class="warning-msg">`,
		`<b>`,
		el.Title,
		`</b>`,
		el.Message,
		`</div>`)

	return strings.Join(result[:], "\n")
}

func Delimiter() string {
	var result []string

	result = append(result, `<div class="ce-delimiter cdx-block"></div>`)

	return strings.Join(result[:], "\n")
}

func Alert(el *domain.EditorJSDataAlert) string {
	var result []string

	result = append(result, `<div class="cdx-alert cdx-alert-`+el.Type+`">`,
		`<div class="cdx-alert__message">`,
		el.Message,
		`</div>`,
		`</div>`)

	return strings.Join(result[:], "\n")
}

func List(el *domain.EditorJSDataList) string {
	var result []string

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
		result = append(result, helpers.CreateHTMLNestedList(itemsList, listStyle))

	} else {
		result = append(result, "<"+listStyle+">")

		for _, item := range el.Items {
			result = append(result, "<li>"+fmt.Sprintf("%v", item)+"</li>")
		}

		result = append(result, "</"+listStyle+">")
	}

	return strings.Join(result[:], "\n")
}

func Checklist(el *domain.EditorJSDataChecklist) string {
	var result []string

	items, err := json.Marshal(el.Items)
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []domain.ChecklistItem

	err = json.Unmarshal(items, &itemsList)
	if err == nil {
		result = append(result, `<div class="cdx-block cdx-checklist">`)

		for _, item := range itemsList {
			if item.Checked {
				result = append(result, `<div class="cdx-checklist__item cdx-checklist__item--checked">`)
			} else {
				result = append(result, `<div class="cdx-checklist__item cdx-checklist">`)
			}

			result = append(result, `<span class="cdx-checklist__item-checkbox"></span>`,
				`<div class="cdx-checklist__item-text">`+item.Text+`</div>`,
				`</div>`)
		}

		result = append(result, `</div>`)
	}

	return strings.Join(result[:], "\n")
}

func Table(el *domain.EditorJSDataTable) string {
	var result []string

	result = append(result, `<table>`)

	for index, line := range el.Content {
		result = append(result, `<tr>`)

		tag := "td"
		if el.WithHeadings && index == 0 {
			tag = "th"
		}

		for _, info := range line {
			result = append(result, `<`+tag+`>`+info+`</`+tag+`>`)
		}

		result = append(result, `</tr>`)
	}

	result = append(result, `</table>`)

	return strings.Join(result[:], "\n")
}

func AnyButton(el *domain.EditorJSDataAnyButton) string {
	var result []string

	result = append(result, `<a class="AnyButton" href="`+el.Link+`">`+el.Text+`</a>`)

	return strings.Join(result[:], "\n")
}

func Code(el *domain.EditorJSDataCode) string {
	var result []string

	result = append(result, `<pre><code class="language-`+el.LanguageCode+`">`+el.Code,
		`</code></pre>`)

	return strings.Join(result[:], "\n")
}

func Raw(el *domain.EditorJSDataRaw) string {
	var result []string

	content := strings.ReplaceAll(el.Html,"<","&lt;")
	content = strings.ReplaceAll(content,">","&gt;")

	result = append(result, `<pre class="preRaw"><code>`+content,
		`</code></pre>`)

	return strings.Join(result[:], "\n")
}

func Image(el *domain.EditorJSDataImage) string {
	classes := ""
	url := ""

	if el.File.URL != "" {
		url = el.File.URL
	} else {
		url = el.URL
	}

	if el.WithBorder {
		classes += "img-with-border "
	}

	if el.Stretched {
		classes += "img-stretched "
	}

	if el.WithBackground {
		classes += "img-with-background "
	}

	return fmt.Sprintf(`<div class="%s" ><img src="%s" alt="%s" title="%s" /></div>`, classes, url, el.Caption, el.Caption)
}

func LinkTool(el *domain.EditorJSDataLinkTool) string {
	var result []string

	result = append(result, `<a href="`+el.Link+`" target="_Blank" rel="nofollow noindex noreferrer" class="linkTool-anchor">`,
		`<div class="linkTool-content">`,
		`<div class="linkTool-left-side">`,
		`<div class="linkTool-title">`,
		el.Meta.Title,
		`</div>`,
		`<div class="linkTool-description">`,
		el.Meta.Description,
		`</div>`,
		`<div class="linkTool-title-anchor">`,
		strings.ReplaceAll(strings.ReplaceAll(el.Link, "https://", ""), "http://", ""),
		`</div>`,
		`</div>`,
		`<div class="linkTool-image" style="background: url(`+el.Meta.Image.URL+`)"></div>`,
		`</div>`,
		`</a>`)

	return strings.Join(result[:], "\n")
}

func Attaches(el *domain.EditorJSDataAttaches) string {
	var result []string

	result = append(result, `<a href="`+el.File.URL+`" rel="noopener noreferrer" target="_blank" class="attaches-anchor">`,
		`<div class="attaches-content">`,
		`<div class="attaches-left" style="background: url(https://i.ibb.co/K7Myr2k/file-icon.png)"></div>`,
		`<div class="attaches-center">`,
		`<div class="attaches-filename">`,
		el.File.Name,
		`</div>`,
		`<div class="attaches-size">`,
		helpers.HumanFileSize(el.File.Size),
		`</div>`,
		`</div>`,
		`<div class="attaches-right" style="background: url(https://i.ibb.co/VYyHr6C/download-icon.png)"></div>`,
		`</div>`,
		`</a>`)

	return strings.Join(result[:], "\n")
}

func Embed(el *domain.EditorJSDataEmbed) string {
	var result []string

	result = append(result, `<div class="embed-block" style="max-width: `+strconv.Itoa(el.Width)+`px">`,
		`<div class="embed-title">`+el.Caption+`</div>`,
		`<iframe width="`+strconv.Itoa(el.Width)+`" height="`+strconv.Itoa(el.Height)+`" src="`+el.Embed+`" title="`+el.Caption+`" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`,
		`<div class="embed-bottom">`,
		`<a href="`+el.Source+`" target="_Blank">Watch on `+el.Service+`</a>`,
		`</div>`,
		`</div>`)
	return strings.Join(result[:], "\n")
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
