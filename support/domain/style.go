package domain

type StyleMap struct {
	StyleName          string            `json:"styleName"`
	SpaceBetweenBlocks string            `json:"spaceBetweenBlocks"`
	Alignment          map[string]string `json:"alignment"`
	Quote              QuoteStyle        `json:"quote"`
	Warning            string            `json:"warning"`
	Delimiter          string            `json:"delimiter"`
	Alert              map[string]string `json:"alert"`
	List               ListStyle         `json:"list"`
	Checklist          ChecklistStyle    `json:"checklist"`
	Table              TableStyle        `json:"table"`
	AnyButton          string            `json:"anyButton"`
	Code               CodeStyle         `json:"code"`
	Raw                CodeStyle         `json:"raw"`
	Image              ImageStyle        `json:"image"`
	LinkTool           LinkToolStyle     `json:"linkTool"`
	Attaches           AttachesStyle     `json:"attaches"`
	Embed              EmbedStyle        `json:"embed"`
}

type QuoteStyle struct {
	Figure     string `json:"figure"`
	Blockquote string `json:"blockquote"`
	Figcaption string `json:"figcaption"`
}

type ListStyle struct {
	Group       string `json:"group"`
	NestedGroup string `json:"nestedGroup"`
	Item        string `json:"item"`
}

type ChecklistStyle struct {
	Block             string `json:"block"`
	Item              string `json:"item"`
	Text              string `json:"text"`
	CheckboxChecked   string `json:"checkboxChecked"`
	CheckboxUnchecked string `json:"checkboxUnchecked"`
}

type TableStyle struct {
	Table  string `json:"table"`
	Row    string `json:"row"`
	CellTH string `json:"cellTH"`
	CellTD string `json:"cellTD"`
}

type CodeStyle struct {
	Pre  string `json:"pre"`
	Code string `json:"code"`
}

type ImageStyle struct {
	Border     string `json:"border"`
	Stretched  string `json:"stretched"`
	Background string `json:"background"`
}

type LinkToolStyle struct {
	Link            string `json:"link"`
	Container       string `json:"container"`
	Row             string `json:"row"`
	LeftColumn      string `json:"leftColumn"`
	RightColumn     string `json:"rightColumn"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	LinkDescription string `json:"linkDescription"`
	Image           string `json:"image"`
}

type AttachesStyle struct {
	Link         string `json:"link"`
	Container    string `json:"container"`
	Row          string `json:"row"`
	LeftColumn   string `json:"leftColumn"`
	CenterColumn string `json:"centerColumn"`
	RightColumn  string `json:"rightColumn"`
	Filename     string `json:"filename"`
	Size         string `json:"size"`
	LeftImage    string `json:"leftImage"`
	RightImage   string `json:"rightImage"`
}

type EmbedStyle struct {
	Block  string `json:"block"`
	Title  string `json:"title"`
	Bottom string `json:"bottom"`
	Link   string `json:"link"`
}