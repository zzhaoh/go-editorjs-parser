package bulma

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/domain"
	"strconv"

	"testing"

	"github.com/stretchr/testify/assert"
)

var obj = Init()

func TestHeaderBlock(t *testing.T) {
	for i := 1; i <= 6; i++ {
		level := strconv.Itoa(i)

		input1 := `{
    "blocks": [
        {
            "type": "header",
            "data": {
                "level": `+level+`,
                "text": "Level `+level+` Header"
            }
        }
    ]
}`

		editorJSON1 := support.ParseEditorJSON(input1)
		obj.Data = support.PrepareData(editorJSON1.Blocks[0]).(*domain.EditorJSDataHeader)

		expected1 := `<h`+level+`>Level `+level+` Header</h`+level+`>`
		actual1 := obj.Header()
		assert.Equal(t, expected1, actual1)
	}

	for i := 1; i <= 6; i++ {
		level := strconv.Itoa(i)

		input1 := `{
    "blocks": [
        {
            "type": "header",
            "data": {
                "level": `+level+`,
                "text": "Level `+level+` Header",
				"anchor": "Anchor Text `+level+`"
            }
        }
    ]
}`

		editorJSON1 := support.ParseEditorJSON(input1)
		obj.Data = support.PrepareData(editorJSON1.Blocks[0])

		expected1 := `<h`+level+` id="anchor-text-`+level+`">Level `+level+` Header</h`+level+`>`
		actual1 := obj.Header()
		assert.Equal(t, expected1, actual1)
	}
}

func TestParagraphBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
			"type": "paragraph",
            "data": {
                "text": "I am a paragraph!"
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<p>I am a paragraph!</p>`
	actual1 := obj.Paragraph()
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
			"type": "paragraph",
            "data": {
                "alignment": "center",
                "text": "I am a paragraph!"
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<p style="text-align:center">I am a paragraph!</p>`
	actual2 := obj.Paragraph()
	assert.Equal(t, expected2, actual2)
}

func TestQuoteBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
			"type": "quote",
            "data": {
                "alignment": "center",
                "caption": "Lao Tzu",
                "text": "The journey of a thousand miles begins with one step."
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<figure class="quote">
<blockquote style="text-align: center">
The journey of a thousand miles begins with one step.
</blockquote>
<figcaption style="text-align: center">
&mdash; Lao Tzu
</figcaption>
</figure>`
	actual := obj.Quote()
	assert.Equal(t, expected, actual)
}

func TestWarningBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
			"type": "warning",
            "data": {
                "message": "Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.",
                "title": "Note:"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="warning-msg">
<b>
Note:
</b>
Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.
</div>`
	actual := obj.Warning()
	assert.Equal(t, expected, actual)
}

func TestDelimiterBlock(t *testing.T) {
	expected := `<div class="ce-delimiter cdx-block"></div>`
	actual := obj.Delimiter()
	assert.Equal(t, expected, actual)
}

func TestAlertBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
			"type": "alert",
            "data": {
                "message": "Something happened that you should know about.",
                "type": "primary"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="cdx-alert cdx-alert-primary">
<div class="cdx-alert__message">
Something happened that you should know about.
</div>
</div>`
	actual := obj.Alert()
	assert.Equal(t, expected, actual)
}

func TestListBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
			"type": "list",
            "data": {
                "items": [
                    "This is a block-styled editor",
                    "Clean output data",
                    "Simple and powerful API"
                ],
                "style": "unordered"
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<ul>
<li>This is a block-styled editor</li>
<li>Clean output data</li>
<li>Simple and powerful API</li>
</ul>`
	actual1 := obj.List()
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
			"type": "list",
            "data": {
                "items": [
                    {
                        "content": "Cars",
                        "items": [
                            {
                                "content": "BMW",
                                "items": [
                                    {
                                        "content": "Z3",
                                        "items": []
                                    },
                                    {
                                        "content": "Z4",
                                        "items": []
                                    }
                                ]
                            },
                            {
                                "content": "Audi",
                                "items": [
                                    {
                                        "content": "A3",
                                        "items": []
                                    },
                                    {
                                        "content": "A1",
                                        "items": []
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "content": "Motorcycle",
                        "items": [
                            {
                                "content": "Ducati",
                                "items": [
                                    {
                                        "content": "916",
                                        "items": []
                                    }
                                ]
                            },
                            {
                                "content": "Yamanha",
                                "items": [
                                    {
                                        "content": "DT 180",
                                        "items": []
                                    }
                                ]
                            },
                            {
                                "content": "Honda",
                                "items": [
                                    {
                                        "content": "VFR 750R",
                                        "items": []
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "style": "ordered"
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<ol>
<li>Cars
<ol>
<li>BMW
<ol>
<li>Z3
</li>
<li>Z4
</li>
</ol>
</li>
<li>Audi
<ol>
<li>A3
</li>
<li>A1
</li>
</ol>
</li>
</ol>
</li>
<li>Motorcycle
<ol>
<li>Ducati
<ol>
<li>916
</li>
</ol>
</li>
<li>Yamanha
<ol>
<li>DT 180
</li>
</ol>
</li>
<li>Honda
<ol>
<li>VFR 750R
</li>
</ol>
</li>
</ol>
</li>
</ol>`
	actual2 := obj.List()
	assert.Equal(t, expected2, actual2)
}

func TestChecklistBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "checklist",
            "data": {
                "items": [
                    {
                        "checked": true,
                        "text": "This is a block-styled editor"
                    },
                    {
                        "checked": false,
                        "text": "Clean output data"
                    },
                    {
                        "checked": true,
                        "text": "Simple and powerful API"
                    }
                ]
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="cdx-block cdx-checklist">
<div class="cdx-checklist__item cdx-checklist__item--checked">
<span class="cdx-checklist__item-checkbox"></span>
<div class="cdx-checklist__item-text">This is a block-styled editor</div>
</div>
<div class="cdx-checklist__item cdx-checklist">
<span class="cdx-checklist__item-checkbox"></span>
<div class="cdx-checklist__item-text">Clean output data</div>
</div>
<div class="cdx-checklist__item cdx-checklist__item--checked">
<span class="cdx-checklist__item-checkbox"></span>
<div class="cdx-checklist__item-text">Simple and powerful API</div>
</div>
</div>`
	actual := obj.Checklist()
	assert.Equal(t, expected, actual)
}

func TestTableBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
            "type": "table",
            "data": {
                "content": [
                    [
                        "Kine",
                        "Pigs",
                        "Chicken"
                    ],
                    [
                        "1 pcs",
                        "3 pcs",
                        "12 pcs"
                    ],
                    [
                        "100$",
                        "200$",
                        "150$"
                    ]
                ],
                "withHeadings": true
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<table>
<tr>
<th>Kine</th>
<th>Pigs</th>
<th>Chicken</th>
</tr>
<tr>
<td>1 pcs</td>
<td>3 pcs</td>
<td>12 pcs</td>
</tr>
<tr>
<td>100$</td>
<td>200$</td>
<td>150$</td>
</tr>
</table>`
	actual1 := obj.Table()
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
            "type": "table",
            "data": {
                "content": [
                    [
                        "Kine",
                        "1 pcs",
                        "100$"
                    ],
                    [
                        "Pigs",
                        "3 pcs",
                        "200$"
                    ],
                    [
                        "Chickens",
                        "12 pcs",
                        "150$"
                    ]
                ]
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<table>
<tr>
<td>Kine</td>
<td>1 pcs</td>
<td>100$</td>
</tr>
<tr>
<td>Pigs</td>
<td>3 pcs</td>
<td>200$</td>
</tr>
<tr>
<td>Chickens</td>
<td>12 pcs</td>
<td>150$</td>
</tr>
</table>`
	actual2 := obj.Table()
	assert.Equal(t, expected2, actual2)
}

func TestAnyButtonBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "AnyButton",
            "data": {
                "link": "https://editorjs.io/",
                "text": "editorjs official"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<a class="AnyButton" href="https://editorjs.io/">editorjs official</a>`
	actual := obj.AnyButton()
	assert.Equal(t, expected, actual)
}

func TestCodeBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
            "type": "code",
            "data": {
                "code": "body {\n font-size: 14px;\n line-height: 16px;\n}"
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<pre><code class="language-">body {
 font-size: 14px;
 line-height: 16px;
}
</code></pre>`
	actual1 := obj.Code()
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
            "type": "code",
            "data": {
                "code": "body {\n font-size: 14px;\n line-height: 16px;\n}",
                "languageCode": "css"
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<pre><code class="language-css">body {
 font-size: 14px;
 line-height: 16px;
}
</code></pre>`
	actual2 := obj.Code()
	assert.Equal(t, expected2, actual2)
}

func TestRawBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "raw",
            "data": {
                "html": "<div style=\"background: #000; color: #fff; font-size: 30px; padding: 50px;\">Any HTML code</div>"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<pre class="preRaw"><code>&lt;div style="background: #000; color: #fff; font-size: 30px; padding: 50px;"&gt;Any HTML code&lt;/div&gt;
</code></pre>`
	actual := obj.Raw()
	assert.Equal(t, expected, actual)
}

func TestImageBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
            "type": "image",
            "data": {
                "caption": "Roadster // tesla.com",
                "file": {
                    "url": "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg"
                },
                "stretched": false,
                "withBackground": true,
                "withBorder": true
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<div class="img-with-border img-with-background " ><img src="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" alt="Roadster // tesla.com" title="Roadster // tesla.com" /></div>`
	actual1 := obj.Image()
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
            "type": "image",
            "data": {
                "caption": "Roadster // tesla.com",
                "stretched": true,
                "url": "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg",
                "withBackground": false,
                "withBorder": false
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<div class="img-stretched " ><img src="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" alt="Roadster // tesla.com" title="Roadster // tesla.com" /></div>`
	actual2 := obj.Image()
	assert.Equal(t, expected2, actual2)
}

func TestLinkToolBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "linkTool",
            "data": {
                "link": "https://codex.so",
                "meta": {
                    "description": "Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.",
                    "image": {
                        "url": "https://codex.so/public/app/img/meta_img.png"
                    },
                    "site_name": "CodeX",
                    "title": "CodeX Team"
                }
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<a href="https://codex.so" target="_Blank" rel="nofollow noindex noreferrer" class="linkTool_anchor">
<div class="linkTool_content">
<div class="linkTool_left_side">
<div class="linkTool_title">
CodeX Team
</div>
<div class="linkTool_description">
Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.
</div>
<div class="linkTool_title_anchor">
codex.so
</div>
</div>
<div class="linkTool_image" style="background: url(https://codex.so/public/app/img/meta_img.png)"></div>
</div>
</a>`
	actual := obj.LinkTool()
	assert.Equal(t, expected, actual)
}

func TestAttachesBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "attaches",
            "data": {
                "file": {
                    "extension": "jpg",
                    "name": "hero.jpg",
                    "size": 260096,
                    "url": "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg"
                },
                "title": "Hero"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<a href="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" rel="noopener noreferrer" target="_blank" class="attaches-anchor">
<div class="attaches-content">
<div class="attaches-left" style="background: url(https://i.ibb.co/K7Myr2k/file-icon.png)"></div>
<div class="attaches-center">
<div class="attaches-filename">
hero.jpg
</div>
<div class="attaches-size">
254 KiB
</div>
</div>
<div class="attaches-right" style="background: url(https://i.ibb.co/VYyHr6C/download-icon.png)"></div>
</div>
</a>`
	actual := obj.Attaches()
	assert.Equal(t, expected, actual)
}

func TestEmbedBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "embed",
            "data": {
                "caption": "Lamborghini Aventador SVJ",
                "embed": "https://www.youtube.com/embed/viW44cUfxCE",
                "height": 315,
                "service": "Youtube",
                "source": "https://www.youtube.com/watch?v=viW44cUfxCE",
                "width": 560
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="embed-block" style="max-width: 560px">
<div class="embed-title">Lamborghini Aventador SVJ</div>
<iframe width="560" height="315" src="https://www.youtube.com/embed/viW44cUfxCE" title="Lamborghini Aventador SVJ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
<div class="embed-bottom">
<a href="https://www.youtube.com/watch?v=viW44cUfxCE" target="_Blank">Watch on Youtube</a>
</div>
</div>`
	actual := obj.Embed()
	assert.Equal(t, expected, actual)
}

func TestImageGalleryBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "imageGallery",
            "data": {
                "bkgMode": false,
                "editImages": true,
                "layoutDefault": true,
                "layoutHorizontal": false,
                "layoutSquare": false,
                "layoutWithFixedSize": false,
                "layoutWithGap": false,
                "urls": [
                    "https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg",
                    "https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg",
                    "https://wallpapercave.com/wp/6L4TVMP.jpg",
                    "https://wallpapercave.com/wp/wp9810772.jpg",
                    "https://wallpapercave.com/wp/wp9121482.jpg",
                    "https://wallpapercave.com/wp/wp9100484.jpg",
                    "https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg"
                ]
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="gg-container">
<div class="gg-box" id="">
<img src="https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg" id="gg-image-0" />
<img src="https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg" id="gg-image-1" />
<img src="https://wallpapercave.com/wp/6L4TVMP.jpg" id="gg-image-2" />
<img src="https://wallpapercave.com/wp/wp9810772.jpg" id="gg-image-3" />
<img src="https://wallpapercave.com/wp/wp9121482.jpg" id="gg-image-4" />
<img src="https://wallpapercave.com/wp/wp9100484.jpg" id="gg-image-5" />
<img src="https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg" id="gg-image-6" />
</div>
</div>`
	actual, _ := obj.ImageGallery()
	assert.Equal(t, expected, actual)
}