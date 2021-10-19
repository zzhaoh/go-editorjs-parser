# Markdown
To genrerate a markdown file, follow the steps below.

#### Import
If you imported go-editorjs-parser into your project, use the code below:
```golang
err = html.Markdown("YOUR/EDITORJS/OUTPUT/FILE.json", "YOUR/OUTPUT/FILE.md")
if err != nil {
    log.Fatal(err)
    return
}
```

#### Build
If you downloaded and build the project, execute with the command:
```shell
./goEditorjsParser -j Input.json -o output/ -t markdown
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
```markdown
#### I don't know why Telegram is the best messenger
```

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
```markdown
Check out our projects on a <a href="https://github.com/codex-team">GitHub page</a>.
```

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
```markdown
> The journey of a thousand miles begins with one step.
>
> --- Lao Tzu
```

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
```markdown
> Note:
>
> --- Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.
```

##### Delimiter
*Input*
```json
{
  "type": "delimiter",
  "data": {}
}
```
*Output*
```markdown
***
```

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
```markdown
| Something happened that you should know about. |
| --- |
```

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
```markdown
1. Cars
    1. BMW
        1. Z3
        2. Z4
    2. Audi
        1. A3
        2. A1
2. Motorcycle
    1. Ducati
        1. 916
    2. Yamanha
        1. DT 180
    3. Honda
        1. VFR 750R
```

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
```markdown
- [x] This is a block-styled editor
- [ ] Clean output data
- [x] Simple and powerful API
```

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
```markdown
| Kine | Pigs | Chicken |
|---|---|---|
| 1 pcs | 3 pcs | 12 pcs |
| 100$ | 200$ | 150$ |
```

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
```markdown
[editorjs official](https://editorjs.io/)
```

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
```markdown
body {
font-size: 14px;
line-height: 16px;
}
```

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
```markdown
<div style="background: #000; color: #fff; font-size: 30px; padding: 50px;">Any HTML code</div>
```

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
```markdown
![Mountain](https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg)
```

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
```markdown
---

# CodeX Team

![CodeX Team](https://pbs.twimg.com/profile_images/993612654861344768/wMPEM5XW_400x400.jpg)

*Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.*

[codex.so](https://codex.so)

---
```

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
```markdown
---

### hero.jpg

###### 254 KiB

[Download](https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg)

---
```

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
```markdown
---

### Lamborghini Aventador SVJ

[Watch on Youtube](https://www.youtube.com/watch?v=viW44cUfxCE)

---
```

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
```markdown
---
![Image 0](https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg)
![Image 1](https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg)
![Image 2](https://wallpapercave.com/wp/6L4TVMP.jpg)
![Image 3](https://wallpapercave.com/wp/wp9810772.jpg)
![Image 4](https://wallpapercave.com/wp/wp9121482.jpg)
![Image 5](https://wallpapercave.com/wp/wp9100484.jpg)
![Image 6](https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg)
---
```
