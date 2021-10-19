# Custom style
You can create your own style. You only have to follow the steps below.

#### Import
If you imported go-editorjs-parser into your project, use the code below:
```golang
err = html.Custom("YOUR/EDITORJS/OUTPUT/FILE.json", "YOUR/OUTPUT/FILE.html", "YOUR/STYLE/FILE.json")
if err != nil {
    log.Fatal(err)
    return
}
```

#### Build
If you downloaded and build the project, execute with the command:
```shell
./goEditorjsParser -j Input.json -s YOUR/STYLE/FILE.json -o output/ -t html
```

To create yor own style, you have to choose between 3 pre-configured styles ([sample](/html/sample?id=configuration), [bootstrap](/html/bootstrap?id=configuration) and [bulma](/html/bulma?id=configuration)). This is necessary to determinate in which format the blocks will be generated.
Copy the JSON file and change the classes values. Do not change the style name, this value will be used to generate the HTML blocks.

At the table below, you can check which variable is used for each available styles.

||Sample|Bootstrap|Bulma||
|---|---|---|---|---|
|styleName|X|X|X|string|
|libraryPaths|X|X|X|string|
|spaceBetweenBlocks|X|X|X|string|
|**alignment**|
|left|X|X|X|string|
|center|X|X|X|string|
|right|X|X|X|string|
|**blocks**|
|**header**|
|h1|X|X|X|string|
|h2|X|X|X|string|
|h3|X|X|X|string|
|h4|X|X|X|string|
|h5|X|X|X|string|
|h6|X|X|X|string|
|**paragraph**|X|X|X|string|
|**quote**|
|figure|X|X|X|string|
|blockquote|X|X|X|string|
|figcaption|X|X|X|string|
|author|||X|string|
|**warning**|
|block|X|X|X|string|
|title|X|X|X|string|
|closeButton|||X|bool|
|**delimiter**|X|X|X|string|
|**alert**|
|block|X|X|X|string|
|closeButton|||X|bool|
|types|
|primary|X|X|X|string|
|secondary|X|X|X|string|
|info|X|X|X|string|
|success|X|X|X|string|
|warning|X|X|X|string|
|danger|X|X|X|string|
|light|X|X|X|string|
|dark|X|X|X|string|
|**list**|
|group|X|X|X|string|
|nestedGroup|X|X|X|string|
|item|X|X|X|string|
|**checklist**|
|block|X|X|X|string|
|item|X|X|X|string|
|text|X|X|X|string|
|checkboxChecked|X|X|X|string|
|checkboxUnchecked|X|X|X|string|
|**table**|
|table|X|X|X|string|
|row|X|X|X|string|
|cellTH|X|X|X|string|
|cellTD|X|X|X|string|
|**anyButton**|X|X|X|string|
|**code**|
|pre|X|X|X|string|
|code|X|X|X|string|
|**raw**|
|pre|X|X|X|string|
|code|X|X|X|string|
|**image**|
|block|||X|string|
|image|||X|string|
|border|X|X|X|string|
|stretched|X|X|X|string|
|background|X|X|X|string|
|**linkTool**|
|link|X|X|X|string|
|container|X|X|X|string|
|row|X|X||string|
|leftColumn|X|X|X|string|
|rightColumn|X|X|X|string|
|title|X|X|X|string|
|description|X|X|X|string|
|linkDescription|X|X|X|string|
|image|X|X|X|string|
|**attaches**|
|link|X|X|X|string|
|container|X|X|X|string|
|row|X|X||string|
|leftColumn|X|X|X|string|
|centerColumn|X|X|X|string|
|rightColumn|X|X|X|string|
|filename|X|X|X|string|
|size|X|X|X|string|
|leftImage|X|X|X|string|
|rightImage|X|X|X|string|
|**embed**|
|block|X|X|X|string|
|title|X|X|X|string|
|bottom|X|X|X|string|
|link|X|X|X|string|