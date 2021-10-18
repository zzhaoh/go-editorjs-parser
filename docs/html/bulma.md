# Bula style
This style was created using [Bulma](https://bulma.io/) version 0.9.3.

#### Import
If you imported go-editorjs-parser into your project, use the code below:
```golang
err = html.Bulma("YOUR/EDITORJS/OUTPUT/FILE.json", "YOUR/OUTPUT/FILE.html")
if err != nil {
    log.Fatal(err)
    return
}
```

#### Build
If you downloaded and build the project, execute with the command:
```shell
./goEditorjsParser -j Input.json -s bulma -o output/ -t html
```

#### Configuration
Pre-configured JSON file for sample style:
```json
{
  "styleName": "bulma",
  "libraryPaths": [
    "https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css"
  ],
  "pageHead": [
    "<meta charset=\"utf-8\">",
    "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">"
  ],
  "spaceBetweenBlocks": "full",
  "alignment": {
    "left": "has-text-left",
    "center": "has-text-centered",
    "right": "has-text-right"
  },
  "blocks": {
    "header": {
      "h1": "title is-1",
      "h2": "title is-2",
      "h3": "title is-3",
      "h4": "title is-4",
      "h5": "title is-5",
      "h6": "title is-6"
    },
    "paragraph": "",
    "quote": {
      "figure": "",
      "blockquote": "",
      "figcaption": "",
      "author": "is-italic"
    },
    "warning": {
      "block": "notification is-warning is-light",
      "title": "has-text-weight-bold",
      "closeButton": true
    },
    "delimiter": "has-text-centered is-size-4",
    "alert": {
      "block": "notification",
      "closeButton": true,
      "types": {
        "primary": "is-primary",
        "secondary": "is-link",
        "info": "is-info",
        "success": "is-success",
        "warning": "is-warning",
        "danger": "is-danger",
        "light": "is-light",
        "dark": "is-dark"
      }
    },
    "list": {
      "group": "content",
      "nestedGroup": "",
      "item": ""
    },
    "checklist": {
      "block": "",
      "item": "",
      "text": "",
      "checkboxChecked": "tag is-link is-rounded",
      "checkboxUnchecked": "tag is-rounded"
    },
    "table": {
      "table": "table is-striped is-bordered",
      "row": "",
      "cellTH": "",
      "cellTD": ""
    },
    "anyButton": "button is-link",
    "code": {
      "pre": "",
      "code": ""
    },
    "raw": {
      "pre": "",
      "code": ""
    },
    "image": {
      "block": "image",
      "image": "",
      "border": "",
      "stretched": "is-fullwidth",
      "background": "has-background-primary p-5"
    },
    "linkTool": {
      "link": "has-text-black-bis",
      "container": "columns is-full p-5 has-background-light m-5",
      "leftColumn": "column is-10",
      "rightColumn": "column is-2",
      "title": "column is-12 has-text-weight-bold",
      "description": "column is-12",
      "linkDescription": "column is-12 has-text-grey-light",
      "image": "image is-96x96"
    },
    "attaches": {
      "link": "has-text-black-bis",
      "container": "columns is-full p-5 has-background-light m-5",
      "leftColumn": "column is-2",
      "centerColumn": "column is-9",
      "rightColumn": "column is-1",
      "filename": "column is-12 has-text-weight-bold",
      "size": "column is-12 has-text-grey-light",
      "leftImage": "image is-96x96",
      "rightImage": "image is-48x48"
    },
    "embed": {
      "block": "",
      "title": "has-text-grey-dark p-2 has-text-weight-bold",
      "bottom": "p-2 has-text-right is-italic",
      "link": "has-text-danger-dark"
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

![](_media/bootstrap/header.png)
```html
<div class="content">
    <h4 id="anchor-text" class="title is-4">I don't know why Telegram is the best messenger</h4>
</div>
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

![](_media/bootstrap/paragraph.png)
```html
<div class="content">
    <p class=" has-text-centered">Check out our projects on a <a href="https://github.com/codex-team">GitHub page</a>.</p>
</div>
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

![](_media/bootstrap/quote.png)
```html
<div class="content">
    <blockquote class=" has-text-centered">
        The journey of a thousand miles begins with one step.
        <p class="is-italic">
            Lao Tzu
        </p>
    </blockquote>
</div>
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

![](_media/bootstrap/warning.png)
```html
<div class="notification is-warning is-light">
    <button class="delete"></button>
    <span class="has-text-weight-bold">
    Note:
    </span>
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

![](_media/bootstrap/delimiter.png)
```html
<div class="has-text-centered is-size-4">***</div>
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

![](_media/bootstrap/alert.png)
```html
<div class="notification is-primary">
    <button class="delete"></button>
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

![](_media/bootstrap/list.png)
```html
<div class="content"><ol class="content">
    <li class="">Cars</li>
    <ol class="">
        <li class="">BMW</li>
        <ol class="">
            <li class="">Z3</li>
            </li>
            <li class="">Z4</li>
            </li>
        </ol>
        </li>
        <li class="">Audi</li>
        <ol class="">
            <li class="">A3</li>
            </li>
            <li class="">A1</li>
            </li>
        </ol>
        </li>
    </ol>
    </li>
    <li class="">Motorcycle</li>
    <ol class="">
        <li class="">Ducati</li>
        <ol class="">
            <li class="">916</li>
            </li>
        </ol>
        </li>
        <li class="">Yamanha</li>
        <ol class="">
            <li class="">DT 180</li>
            </li>
        </ol>
        </li>
        <li class="">Honda</li>
        <ol class="">
            <li class="">VFR 750R</li>
            </li>
        </ol>
        </li>
    </ol>
    </li>
</ol></div>
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

![](_media/bootstrap/checklist.png)
```html
<div class="">
    <div class="">
        <span class="tag is-link is-rounded">&#10004;</span>
        <span class="">This is a block-styled editor</span>
    </div>
    <div class="">
        <span class="tag is-rounded">&nbsp;-&nbsp;</span>
        <span class="">Clean output data</span>
    </div>
    <div class="">
        <span class="tag is-link is-rounded">&#10004;</span>
        <span class="">Simple and powerful API</span>
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

![](_media/bootstrap/table.png)
```html
<table class="table is-striped is-bordered">
    <tr class="">
        <th class="">Kine</th>
        <th class="">Pigs</th>
        <th class="">Chicken</th>
    </tr>
    <tr class="">
        <td class="">1 pcs</td>
        <td class="">3 pcs</td>
        <td class="">12 pcs</td>
    </tr>
    <tr class="">
        <td class="">100$</td>
        <td class="">200$</td>
        <td class="">150$</td>
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

![](_media/bootstrap/anybutton.png)
```html
<a class="button is-link" href="https://editorjs.io/">editorjs official</a>
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

![](_media/bootstrap/code.png)
```html
<pre class="">
    <code class="">
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

![](_media/bootstrap/raw.png)
```html
<pre class="">
    <code class="">
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
    "withBorder" : true,
    "withBackground" : true,
    "stretched" : true
  }
}
```
*Output*

![](_media/bootstrap/image.png)
```html
<figure class="image has-background-primary p-5" >
    <img class="  is-fullwidth" src="https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg" alt="Mountain" title="Mountain" />
</figure>

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

![](_media/bootstrap/linktool.png)
```html
<a href="https://codex.so" target="_Blank" rel="nofollow noindex noreferrer" class="has-text-black-bis">
    <div class="columns is-full p-5 has-background-light m-5">
        <div class="column is-10">
            <div class="column is-12 has-text-weight-bold">
                CodeX Team
            </div>
            <div class="column is-12">
                Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.
            </div>
            <div class="column is-12 has-text-grey-light">
                codex.so
            </div>
        </div>
        <div class="column is-2">
            <img class="image is-96x96" src="https://pbs.twimg.com/profile_images/993612654861344768/wMPEM5XW_400x400.jpg" />
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

![](_media/bootstrap/attaches.png)
```html
<a href="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" rel="noopener noreferrer" target="_blank" class="has-text-black-bis">
    <div class="columns is-full p-5 has-background-light m-5">
        <div class="column is-2" >
            <img class="image is-96x96" src="https://i.ibb.co/K7Myr2k/file-icon.png" />
        </div>
        <div class="column is-9">
            <div class="column is-12 has-text-weight-bold">
                hero.jpg
            </div>
            <div class="column is-12 has-text-grey-light">
                254 KiB
            </div>
        </div>
        <div class="column is-1" >
            <img class="image is-48x48" src="https://i.ibb.co/VYyHr6C/download-icon.png" />
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

![](_media/bootstrap/embed.png)
```html
<div class="" style="max-width: 560px">
    <div class="has-text-grey-dark p-2 has-text-weight-bold">Lamborghini Aventador SVJ</div>
    <iframe width="560" height="315" src="https://www.youtube.com/embed/viW44cUfxCE" title="Lamborghini Aventador SVJ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
    <div class="p-2 has-text-right is-italic">
        <a class="has-text-danger-dark" href="https://www.youtube.com/watch?v=viW44cUfxCE" target="_Blank">Watch on Youtube</a>
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

![](_media/bootstrap/imagegallery.png)
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
