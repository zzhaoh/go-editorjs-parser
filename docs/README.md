<div align="center">
<img src="https://i.ibb.co/nCDpvSH/logo-small.png" alt="go-editorjs-parser logo">
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
	html.Sample("YOUR/EDITORJS/OUTPUT/FILE.json", "YOUR/OUTPUT/FILE.html")
	html.Bootstrap("YOUR/EDITORJS/OUTPUT/FILE.json", "YOUR/OUTPUT/FILE.html")
    html.Bulma("YOUR/EDITORJS/OUTPUT/FILE.json", "YOUR/OUTPUT/FILE.html")
    html.Custom("YOUR/EDITORJS/OUTPUT/FILE.json", "YOUR/OUTPUT/FILE.html", "YOUR/STYLE/FILE.json")

    // Parse a EditorJS JSOn file to MARKDOWN file
    markdown.Parser("YOUR/EDITORJS/OUTPUT/FILE.json", "YOUR/OUTPUT/FILE.md")
}
```

or you can build and execute
```
go build goEditorjsParser.go
```

Usage of goEditorjsParser
```
./goEditorjsParser -j <JSONFilePath> -s <StylePath> -o <OutputFilePath> -t <OutputFileType>

-j = Path to a JSON file. MANDATORY
-s = Style CSS to be used. If not set, the sample style will be used. Possible values: "sample", "bootstrap", "bulma" or "PATH/TO/YOUR/CUSTOM/STYLE/MAP"
-o = Output file path. If not set, root path will be used
-t = Output file type. Possible values: markdown or html
```

Example:
```
./goEditorjsParser -j Input.json
./goEditorjsParser -j Input.json -s bulma
./goEditorjsParser -j Input.json -s bootstrap -o output/
./goEditorjsParser -j Input.json -s sample -o output/ -t html
./goEditorjsParser -j Input.json -o output/ -t markdown
```
