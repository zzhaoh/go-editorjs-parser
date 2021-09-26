<div align="center">
<img width="300" src="https://i.ibb.co/nCDpvSH/logo-small.png" alt="go-editorjs-parser logo">
<br><br>
<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" >
<img src="https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E" >
<img src="https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white" >
<img src="https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white" >
<br><br><br>
<a href="#"><img src="https://img.shields.io/badge/build-passing-green"></a>
<a href="#"><img src="https://img.shields.io/badge/test-success-green"></a>
<a href="#"><img src="https://img.shields.io/badge/Editor.js-2.22.2-blue"></a>
<a href="https://paypal.me/rodrigoodhin"><img src="https://img.shields.io/badge/donate-PayPal-blue"></a>
</div>


# EditorJS Tool for Editor.js

A Golang library which converts **[Editor.js](https://editorjs.io)** JSON output to pure **Markdown** or **HTML**.

&nbsp;
&nbsp;
&nbsp;

## Usage

Is it possible to import in a project

```go
import (
    "gitlab.com/rodrigoodhin/go-editorjs-parser/pkg/html"
    "gitlab.com/rodrigoodhin/go-editorjs-parser/pkg/markdown"
)

func main() {

    // Parse a EditorJS JSOn file to HTML file
    html.Parser("JSON/FILE/PATH", "NEW/HTML/FILE/PATH", "CSS/STYLE/FILE/PATH", 20)

    // Parse a EditorJS JSOn file to MARKDOWN file
    markdown.Parser("JSON/FILE/PATH", "NEW/MARKDOWN/FILE/PATH")
}
```

or you can build and execute
```
go build goEditorjsParser.go
```

Usage of goEditorjsParser
```
./goEditorjsParser -j <JSONFilePath> -s <SeparatorSize> -c <StylePath> -o <OutputFilePath> -t <OutputFileType>

-j = Path to a JSON file. MANDATORY
-s = Separator size (in pixels) between blocks. Default 20
-c = Style CSS path to be used. If not set, the default style will be used
-o = Output file path. If not set, root path will be used
-t = Output file type. Possible values: markdown or html
```

Example:
```
./goEditorjsParser -j Input.json
./goEditorjsParser -j Input.json -s 30 
./goEditorjsParser -j Input.json -s 15 -c Style.css
./goEditorjsParser -j Input.json -s 10 -c Style.css -o output/
./goEditorjsParser -j Input.json -s 10 -c Style.css -o output/ -t html
```

&nbsp;
&nbsp;
&nbsp;

## Supported Plugins

### Block Tools

#### Text and typography

- [@editorjs/paragraph](https://github.com/editor-js/paragraph) - Paragraph
- [@editorjs/header](https://github.com/editor-js/header) - Header
- [@editorjs/quote](https://github.com/editor-js/quote) - Quote
- [@editorjs/warning](https://github.com/editor-js/warning) - Warning
- [@editorjs/delimiter](https://github.com/editor-js/delimiter) - Delimiter
- [editorjs-alert](https://github.com/vishaltelangre/editorjs-alert) - Alert
- [paragraph-with-alignment](https://github.com/kaaaaaaaaaaai/paragraph-with-alignment) - Paragraph with alignment
- [header-with-anchor ](https://github.com/Aleksst95/header-with-anchor) - Header with anchor

#### Lists

- [@editorjs/list](https://github.com/editor-js/list) - List
- [@editorjs/nested-list](https://github.com/editor-js/nested-list) - Nested List
- [@editorjs/checklist](https://github.com/editor-js/checklist) - Checklist

#### Media & Embed

- [@editorjs/image](https://github.com/editor-js/image) - Image
- [@editorjs/simple-image](https://github.com/editor-js/simple-image) - Add images by pasting image URLs. no uploader required
- [@editorjs/link](https://github.com/editor-js/link) - Link with preview
- [@editorjs/attaches](https://github.com/editor-js/attaches) - Attach files
- [@editorjs/embed](https://github.com/editor-js/embed) - Embed content
- [simple-image-editorjs](https://github.com/PaulKinlan/simple-image) - Image
- [editorjs-github-gist-plugin](https://github.com/ranemihir/editorjs-github-gist-plugin) - @TODO
- [editorjs-social-post-plugin](https://github.com/ranemihir/editorjs-social-post-plugin) - @TODO
- [editorjs-inline-image](https://github.com/kommitters/editorjs-inline-image) - @TODO
- [mr8bit/carousel-editorjs](https://github.com/mr8bit/carousel-editorjs) - @TODO
- [mdgaziur/EditorJS-LaTeX](https://github.com/mdgaziur/EditorJS-LaTeX) - @TODO
- [rodrigoodhin/editorjs-image-gallery](https://github.com/rodrigoodhin/editorjs-image-gallery) - Image Gallery

#### Table

- [@editorjs/table](https://github.com/editor-js/table) - Table
- [editorjs-table](https://github.com/4rw44z/editorjs-table) - Table

#### Code

- [@editorjs/code](https://github.com/editor-js/code) - Code examples
- [@editorjs/raw](https://github.com/editor-js/raw) - Raw HTML code
- [editor-js-code ](https://github.com/paraswaykole/editor-js-code) - Code examples with language name
- [editorjs-codemirror](https://github.com/alexiej/editorjs-codemirror) - Code mirror
- [@bomdi/codebox](https://github.com/BomdiZane/codebox) - @TODO

#### Button

- [editorjs-button](https://github.com/kaaaaaaaaaaai/editorjs-button) - Button with link and text

&nbsp;
&nbsp;
&nbsp;

## HTML code generated

*Based on these HTML outputs, you can create your own css file and use it to generate your html page.*

### header

###### Input

```json
{
  "type": "header",
  "data": {
    "text": "I don't know why Telegram is the best messenger",
    "level": 4,
    "anchor": "Anchor Text"
  }
}
```

###### Output

```html
<h4 id="anchor-text">I don't know why Telegram is the best messenger</h4>
```

### paragraph

###### Input

```json
{
  "type": "paragraph",
  "data": {
    "text": "Check out our projects on a <a href=\"https://github.com/codex-team\">GitHub page</a>.",
    "alignment": "center"
  }
}
```

###### Output

```html
<p style="text-align:center">Check out our projects on a <a href="https://github.com/codex-team">GitHub page</a>.</p>
```

### quote

###### Input

```json
{
  "type": "quote",
  "data": {
    "text": "The journey of a thousand miles begins with one step.",
    "caption": "Lao Tzu",
    "alignment": "center"
  }
}
```

###### Output

```html
<figure class="quote">
    <blockquote style="text-align: center">
        The journey of a thousand miles begins with one step.
    </blockquote>
    <figcaption style="text-align: center">
        &mdash; Lao Tzu
    </figcaption>
</figure>
```

### warning

###### Input

```json
{
  "type": "warning",
  "data": {
    "title": "Note:",
    "message": "Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff."
  }
}
```

###### Output

```html
<div class="warning-msg">
    <b>
        Note:
    </b>
    Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.
</div>
```

### delimiter

###### Input

```json
{
  "type": "delimiter",
  "data": {}
}
```

###### Output

```html
<div class="ce-delimiter cdx-block"></div>
```

### alert

###### Input

```json
{
  "type": "alert",
  "data": {
    "type": "info",
    "message": "We did something <strong>new</strong>..."
  }
}
```

###### Output

```html
<div class="cdx-alert cdx-alert-info">
    <div class="cdx-alert__message">
        We did something <strong>new</strong>...
    </div>
</div>
```

### list

###### Input

```json
{
  "type": "list",
  "data": {
    "style": "ordered",
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
    ]
  }
}
```

###### Output

```html
<ol>
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
</ol>
```

### checklist

###### Input

```json
{
  "type": "checklist",
  "data": {
    "items": [
      {
        "text": "This is a block-styled editor",
        "checked": true
      },
      {
        "text": "Clean output data",
        "checked": false
      },
      {
        "text": "Simple and powerful API",
        "checked": true
      }
    ]
  }
}
```

###### Output

```html
<div class="cdx-block cdx-checklist">
    <div class="cdx-checklist__item cdx-checklist__item--checked">
        <span class="cdx-checklist__item-checkbox"></span>
        <div class="cdx-checklist__item-text" contenteditable="true">This is a block-styled editor</div>
    </div>
    <div class="cdx-checklist__item cdx-checklist">
        <span class="cdx-checklist__item-checkbox"></span>
        <div class="cdx-checklist__item-text" contenteditable="true">Clean output data</div>
    </div>
    <div class="cdx-checklist__item cdx-checklist__item--checked">
        <span class="cdx-checklist__item-checkbox"></span>
        <div class="cdx-checklist__item-text" contenteditable="true">Simple and powerful API</div>
    </div>
</div>
```

### table

###### Input

```json
{
  "type" : "table",
  "data" : {
    "withHeadings": true,
    "content" : [ [ "Kine", "Pigs", "Chicken" ], [ "1 pcs", "3 pcs", "12 pcs" ], [ "100$", "200$", "150$" ] ]
  }
}
```

###### Output

```html
<table>
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
</table>
```

### AnyButton

###### Input

```json
{
  "type" : "AnyButton",
  "data" : {
    "link" : "https://editorjs.io/",
    "text" : "editorjs official"
  }
}
```

###### Output

```html
<a class="AnyButton" href="https://editorjs.io/">editorjs official</a>
```

### code

###### Input

```json
{
  "type" : "code",
  "data" : {
    "code": "body {\n font-size: 14px;\n line-height: 16px;\n}",
    "languageCode": "css"
  }
}
```

###### Output

```html
<pre><code class="language-css">body {
 font-size: 14px;
 line-height: 16px;
}
</code></pre>
```

### raw

###### Input

```json
{
  "type" : "raw",
  "data" : {
    "html": "<div style=\"background: #000; color: #fff; font-size: 30px; padding: 50px;\">Any HTML code</div>"
  }
}
```

###### Output

```html
<pre class="preRaw"><code>&lt;div style="background: #000; color: #fff; font-size: 30px; padding: 50px;"&gt;Any HTML code&lt;/div&gt;
</code></pre>
```

### image

###### Input

```json
{
  "type" : "image",
  "data" : {
    "file": {
      "url" : "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg"
    },
    "caption" : "Roadster // tesla.com",
    "withBorder" : true,
    "withBackground" : true,
    "stretched" : true
  }
}
```
```json
{
  "type" : "image",
  "data" : {
    "url" : "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg",
    "caption" : "Roadster // tesla.com",
    "withBorder" : true,
    "withBackground" : true,
    "stretched" : true
  }
}
```

###### Output

```html
<div class="img-with-border img-with-background img-stretched " ><img src="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" alt="Roadster // tesla.com" title="Roadster // tesla.com" /></div>
```

### linkTool

###### Input

```json
{
  "type" : "linkTool",
  "data" : {
    "link" : "https://codex.so",
    "meta" : {
      "title" : "CodeX Team",
      "site_name" : "CodeX",
      "description" : "Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.",
      "image" : {
        "url" : "https://codex.so/public/app/img/meta_img.png"
      }
    }
  }
}
```

###### Output

```html
<a href="https://codex.so" target="_Blank" rel="nofollow noindex noreferrer" class="linkTool-anchor">
    <div class="linkTool-content">
        <div class="linkTool-left-side">
            <div class="linkTool-title">
                CodeX Team
            </div>
            <div class="linkTool-description">
                Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.
            </div>
            <div class="linkTool-title-anchor">
                codex.so
            </div>
        </div>
        <div class="linkTool-image" style="background: url(https://codex.so/public/app/img/meta_img.png)"></div>
    </div>
</a>
```

### attaches

###### Input

```json
{
  "type" : "attaches",
  "data" : {
    "file": {
      "url" : "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg",
      "size": 260096,
      "name": "hero.jpg",
      "extension": "jpg"
    },
    "title": "Hero"
  }
}
```

###### Output

```html
<a href="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" rel="noopener noreferrer" target="_blank" class="attaches-anchor">
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
</a>
```

### embed

###### Input

```json
{
  "type" : "embed",
  "data" : {
    "service" : "Youtube",
    "source" : "https://www.youtube.com/watch?v=viW44cUfxCE",
    "embed" : "https://www.youtube.com/embed/viW44cUfxCE",
    "width" : 560,
    "height" : 315,
    "caption" : "Lamborghini Aventador SVJ"
  }
}
```

###### Output

```html
<div class="embed-block" style="max-width: 560px">
    <div class="embed-title">Lamborghini Aventador SVJ</div>
    <iframe width="560" height="315" src="https://www.youtube.com/embed/viW44cUfxCE" title="Lamborghini Aventador SVJ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
    <div class="embed-bottom">
        <a href="https://www.youtube.com/watch?v=viW44cUfxCE" target="_Blank">Watch on Youtube</a>
    </div>
</div>
```

### imageGallery

###### Input

```json
{
  "type": "imageGallery",
  "data": {
    "urls": [
      "https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg",
      "https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg",
      "https://wallpapercave.com/wp/6L4TVMP.jpg",
      "https://wallpapercave.com/wp/wp9810772.jpg",
      "https://wallpapercave.com/wp/wp9121482.jpg",
      "https://wallpapercave.com/wp/wp9100484.jpg",
      "https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg"
    ],
    "editImages": true,
    "bkgMode": false,
    "layoutDefault": true,
    "layoutHorizontal": false,
    "layoutSquare": false,
    "layoutWithGap": false,
    "layoutWithFixedSize": false
  }
}
```

###### Output

```html
<div class="gg-container">
    <div class="gg-box" id="">
        <img src="https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg" id="gg-image-0" />
        <img src="https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg" id="gg-image-1" />
        <img src="https://wallpapercave.com/wp/6L4TVMP.jpg" id="gg-image-2" />
        <img src="https://wallpapercave.com/wp/wp9810772.jpg" id="gg-image-3" />
        <img src="https://wallpapercave.com/wp/wp9121482.jpg" id="gg-image-4" />
        <img src="https://wallpapercave.com/wp/wp9100484.jpg" id="gg-image-5" />
        <img src="https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg" id="gg-image-6" />
    </div>
</div>
```

&nbsp;
&nbsp;
&nbsp;

## LICENSE

[MIT License](/LICENSE)

