package domain

type EditorJSMethods interface {
	SetData(d interface{})
	LoadLibrary() (styles []string)
	Header() string
	Paragraph() string
	Quote() string
	Warning() string
	Delimiter() string
	Alert() string
	List() string
	Checklist() string
	Table() string
	AnyButton() string
	Code() string
	Raw() string
	Image() string
	LinkTool() string
	Attaches() string
	Embed() string
	ImageGallery() (string, string)
}
