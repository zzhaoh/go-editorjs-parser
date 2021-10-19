# Sample style
This style was created by me based on the plugins preview.

#### Import
If you imported go-editorjs-parser into your project, use the code below:
```golang
err = html.Sample("YOUR/EDITORJS/OUTPUT/FILE.json", "YOUR/OUTPUT/FILE.html")
if err != nil {
    log.Fatal(err)
    return
}
```

#### Build
If you downloaded and build the project, execute with the command:
```shell
./goEditorjsParser -j Input.json -s sample -o output/ -t html
```

#### Configuration
Pre-configured JSON file for sample style:
```json
{
  "styleName": "sample",
  "libraryPaths": [
    "sample.min.css"
  ],
  "pageHead": [
    "<meta charset=\"utf-8\">",
    "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">"
  ],
  "spaceBetweenBlocks": "space-between-blocks",
  "alignment": {
    "left": "alignment_text_left",
    "center": "alignment_text_center",
    "right": "alignment_text_right"
  },
  "blocks": {
    "header": {
      "h1": "",
      "h2": "",
      "h3": "",
      "h4": "",
      "h5": "",
      "h6": ""
    },
    "paragraph": "",
    "quote": {
      "figure": "quote_figure",
      "blockquote": "quote_blockquote",
      "figcaption": "quote_figcaption"
    },
    "warning": {
      "block": "warning_msg",
      "title": ""
    },
    "delimiter": "delimiter_block",
    "alert": {
      "block": "alert_box",
      "types": {
        "primary": "alert_primary",
        "secondary": "alert_secondary",
        "info": "alert_info",
        "success": "alert_success",
        "warning": "alert_warning",
        "danger": "alert_danger",
        "light": "alert_light",
        "dark": "alert_dark"
      }
    },
    "list": {
      "group": "list_group",
      "nestedGroup": "",
      "item": "list_item"
    },
    "checklist": {
      "block": "checklist_block",
      "item": "checklist_item",
      "text": "checklist_item_text",
      "checkboxChecked": "checklist_item_checkbox checklist_checked",
      "checkboxUnchecked": "checklist_item_checkbox"
    },
    "table": {
      "table": "table_block",
      "row": "table_tr",
      "cellTH": "table_th",
      "cellTD": "table_td"
    },
    "anyButton": "anyButton",
    "code": {
      "pre": "code_pre",
      "code": "code_block"
    },
    "raw": {
      "pre": "raw_pre",
      "code": "raw_block"
    },
    "image": {
      "border": "image_with_border",
      "stretched": "image_stretched",
      "background": "image_with_background"
    },
    "linkTool": {
      "link": "",
      "container": "linkTool_content",
      "row": "",
      "leftColumn": "linkTool_left_side",
      "rightColumn": "linkTool_image_block",
      "title": "linkTool_title",
      "description": "linkTool_description",
      "linkDescription": "linkTool_anchor",
      "image": "linkTool_image"
    },
    "attaches": {
      "link": "",
      "container": "attaches_content",
      "row": "",
      "leftColumn": "attaches_left",
      "centerColumn": "attaches_center",
      "rightColumn": "attaches_right",
      "filename": "attaches_filename",
      "size": "attaches_size",
      "leftImage": "attaches_image",
      "rightImage": "attaches_image"
    },
    "embed": {
      "block": "embed_block",
      "title": "embed_title",
      "bottom": "embed_bottom",
      "link": "embed_link"
    }
  }
}
```

#### Blocks

##### Header
*Input*
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
*Output*

![](_media/sample/header.png)
```html
<h4 id="anchor-text">I don't know why Telegram is the best messenger</h4>
```

&nbsp;

##### Paragraph
*Input*
```json
{
  "type": "paragraph",
  "data": {
    "text": "Check out our projects on a <a href=\"https://github.com/codex-team\">GitHub page</a>.",
    "alignment": "center"
  }
}
```
*Output*

![](_media/sample/paragraph.png)
```html
<p class=" alignment_text_center">Check out our projects on a <a href="https://github.com/codex-team">GitHub page</a>.</p>
```

&nbsp;

##### Quote
*Input*
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
*Output*

![](_media/sample/quote.png)
```html
<figure class="quote_figure alignment_text_center">
    <blockquote class="quote_blockquote">
        The journey of a thousand miles begins with one step.
    </blockquote>
    <figcaption class="quote_figcaption">
        Lao Tzu
    </figcaption>
</figure>
```

&nbsp;

##### Warning
*Input*
```json
{
  "type": "warning",
  "data": {
    "title": "Note:",
    "message": "Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff."
  }
}
```
*Output*

![](_media/sample/warning.png)
```html
<div class="warning_msg">
    <b>
        Note:
    </b>
    Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.
</div>
```

&nbsp;

##### Delimiter
*Input*
```json
{
  "type": "delimiter",
  "data": {}
}
```
*Output*

![](_media/sample/delimiter.png)
```html
<div class="delimiter_block">***</div>
```

&nbsp;

##### Alert
*Input*
```json
{
  "type": "alert",
  "data": {
    "type": "primary",
    "message": "Something happened that you should know about."
  }
}
```
*Output*

![](_media/sample/alert.png)
```html
<div class="alert_box alert_primary">
    Something happened that you should know about.
</div>
```

&nbsp;

##### List
*Input*
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
*Output*

![](_media/sample/list.png)
```html
<ol class="list_group">
    <li class="list_item">Cars</li>
    <ol class="">
        <li class="list_item">BMW</li>
        <ol class="">
            <li class="list_item">Z3</li>
            </li>
            <li class="list_item">Z4</li>
            </li>
        </ol>
        </li>
        <li class="list_item">Audi</li>
        <ol class="">
            <li class="list_item">A3</li>
            </li>
            <li class="list_item">A1</li>
            </li>
        </ol>
        </li>
    </ol>
    </li>
    <li class="list_item">Motorcycle</li>
    <ol class="">
        <li class="list_item">Ducati</li>
        <ol class="">
            <li class="list_item">916</li>
            </li>
        </ol>
        </li>
        <li class="list_item">Yamanha</li>
        <ol class="">
            <li class="list_item">DT 180</li>
            </li>
        </ol>
        </li>
        <li class="list_item">Honda</li>
        <ol class="">
            <li class="list_item">VFR 750R</li>
            </li>
        </ol>
        </li>
    </ol>
    </li>
</ol>
```

&nbsp;

##### Checklist
*Input*
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
*Output*

![](_media/sample/checklist.png)
```html
<div class="checklist_block">
    <div class="checklist_item">
        <span class="checklist_item_checkbox checklist_checked">&#10004;</span>
        <span class="checklist_item_text">This is a block-styled editor</span>
    </div>
    <div class="checklist_item">
        <span class="checklist_item_checkbox">&nbsp;-&nbsp;</span>
        <span class="checklist_item_text">Clean output data</span>
    </div>
    <div class="checklist_item">
        <span class="checklist_item_checkbox checklist_checked">&#10004;</span>
        <span class="checklist_item_text">Simple and powerful API</span>
    </div>
</div>
```

&nbsp;

##### Table
*Input*
```json
{
  "type" : "table",
  "data" : {
    "withHeadings": true,
    "content" : [ [ "Kine", "Pigs", "Chicken" ], [ "1 pcs", "3 pcs", "12 pcs" ], [ "100$", "200$", "150$" ] ]
  }
}
```
*Output*

![](_media/sample/table.png)
```html
<table class="table_block">
    <tr class="table_tr">
        <th class="table_th">Kine</th>
        <th class="table_th">Pigs</th>
        <th class="table_th">Chicken</th>
    </tr>
    <tr class="table_tr">
        <td class="table_td">1 pcs</td>
        <td class="table_td">3 pcs</td>
        <td class="table_td">12 pcs</td>
    </tr>
    <tr class="table_tr">
        <td class="table_td">100$</td>
        <td class="table_td">200$</td>
        <td class="table_td">150$</td>
    </tr>
</table>
```

&nbsp;

##### AnyButton
*Input*
```json
{
  "type" : "AnyButton",
  "data" : {
    "link" : "https://editorjs.io/",
    "text" : "editorjs official"
  }
}
```
*Output*

![](_media/sample/anybutton.png)
```html
<a class="anyButton" href="https://editorjs.io/">editorjs official</a>
```

&nbsp;

##### Code
*Input*
```json
{
  "type" : "code",
  "data" : {
    "code": "body {\n font-size: 14px;\n line-height: 16px;\n}",
    "languageCode": "css"
  }
}
```
*Output*

![](_media/sample/code.png)
```html
<pre class="code_pre">
    <code class="code_block">
        body {
         font-size: 14px;
         line-height: 16px;
        }
    </code>
</pre>
```

&nbsp;

##### Raw
*Input*
```json
{
  "type" : "raw",
  "data" : {
    "html": "<div style=\"background: #000; color: #fff; font-size: 30px; padding: 50px;\">Any HTML code</div>"
  }
}
```
*Output*

![](_media/sample/raw.png)
```html
<pre class="raw_pre">
    <code class="raw_block">
        &lt;div style="background: #000; color: #fff; font-size: 30px; padding: 50px;"&gt;Any HTML code&lt;/div&gt;
    </code>
</pre>
```

&nbsp;

##### Image
*Input*
```json
{
  "type" : "image",
  "data" : {
    "file": {
      "url" : "https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg"
    },
    "caption" : "Mountain",
    "withBorder" : false,
    "withBackground" : true,
    "stretched" : false
  }
}
```
*Output*

![](_media/sample/image.png)
```html
<div class="image_with_background" >
    <img class="" src="https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg" alt="Mountain" title="Mountain" />
</div>
```

&nbsp;

##### LinkTool
*Input*
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
        "url" : "https://pbs.twimg.com/profile_images/993612654861344768/wMPEM5XW_400x400.jpg"
      }
    }
  }
}
```
*Output*

![](_media/sample/linktool.png)
```html
<a href="https://codex.so" target="_Blank" rel="nofollow noindex noreferrer" class="">
    <div class="linkTool_content">
        <div class="">
            <div class="linkTool_left_side">
                <div class="linkTool_title">
                    CodeX Team
                </div>
                <div class="linkTool_description">
                    Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.
                </div>
                <div class="linkTool_anchor">
                    codex.so
                </div>
            </div>
            <div class="linkTool_image_block">
                <img class="linkTool_image" src="https://pbs.twimg.com/profile_images/993612654861344768/wMPEM5XW_400x400.jpg" />
            </div>
        </div>
    </div>
</a>
```

&nbsp;

##### Attaches
*Input*
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
*Output*

![](_media/sample/attaches.png)
```html
<a href="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" rel="noopener noreferrer" target="_blank" class="">
    <div class="attaches_content">
        <div class="" >
            <div class="attaches_left" >
                <img class="attaches_image" src="https://i.ibb.co/K7Myr2k/file-icon.png" />
            </div>
            <div class="attaches_center">
                <div class="attaches_filename">
                    hero.jpg
                </div>
                <div class="attaches_size">
                    254 KiB
                </div>
            </div>
            <div class="attaches_right" >
                <img class="attaches_image" src="https://i.ibb.co/VYyHr6C/download-icon.png" />
            </div>
        </div>
    </div>
</a>
```

&nbsp;

##### Embed
*Input*
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
*Output*

![](_media/sample/embed.png)
```html
<div class="embed_block" style="max-width: 560px">
    <div class="embed_title">Lamborghini Aventador SVJ</div>
    <iframe width="560" height="315" src="https://www.youtube.com/embed/viW44cUfxCE" title="Lamborghini Aventador SVJ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
    <div class="embed_bottom">
        <a class="embed_link" href="https://www.youtube.com/watch?v=viW44cUfxCE" target="_Blank">Watch on Youtube</a>
    </div>
</div>
```

&nbsp;

##### ImageGallery
*Input*
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
*Output*

![](_media/sample/imagegallery.png)
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
