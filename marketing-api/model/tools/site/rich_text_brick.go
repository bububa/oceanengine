package site

// FontStyle 文字样式
type FontStyle string

const (
	// FontStyle_BOLD 黑体
	FontStyle_BOLD FontStyle = "bold"
	// FontStyle_ITALIC 斜体
	FontStyle_ITALIC FontStyle = "italic"
	// FontStyle_UNDERLINE 下划线
	FontStyle_UNDERLINE FontStyle = "underline"
)

// TextAlign 对齐方式
type TextAlign string

const (
	// TextAlign_LEFT 左对齐
	TextAlign_LEFT TextAlign = "left"
	// TextAlign_RIGHT 右对齐
	TextAlign_RIGHT TextAlign = "right"
	// TextAlign_CENTER 居中
	TextAlign_CENTER TextAlign = "center"
)

// RichTextBrick 富文本组件
type RichTextBrick struct {
	BaseBrick
	// Content 文本内容，长度至少为1
	Content string `json:"content,omitempty"`
	// Link 跳转链接信息
	Link *Link `json:"link,omitempty"`
	// FontFamily 文本字体
	// 文本字体支持如下： font-x
	// x 的值可为：
	// 0（系统默认字体）
	// 4（方正仿宋简体）
	// 7（方正黑体简体）
	// 10（方正楷体简体）
	// 13（方正书宋简体）
	// 14（华文琥珀）
	// 15（华文隶书）
	// 16（华文新魏）
	// 17（华文行楷）
	// 23（濑户字体）
	// 24（庞门正道标题体）
	// 32（书体坊颜体）
	// 90（文泉驿正黑）
	// 94（杨任东竹石体）
	// 95（站酷高端黑）
	// 96（站酷酷黑）
	// 97（站酷快乐体）
	// 100（装甲明朝体）
	// 101（思源黑体粗体）
	// 102（思源黑体黑体）
	// 103（思源黑体细体）
	// 104（思源黑体纤细）
	// 105（思源黑体中黑）
	// 106（思源黑体常规）
	// 107（思源黑体极细）
	FontFamily int `json:"font_family,omitempty"`
	// FontSize 字号,默认值为14
	FontSize float64 `json:"font_size,omitempty"`
	// FontStyle 文字样式，可选值：
	// 黑体bold
	// 斜体italic
	// 下划线underline,
	// 默认为无样式。可随意组合
	FontStyle FontStyle `json:"font_style,omitempty"`
	// LetterSpacing 字间距，默认值为0
	LetterSpacing int `json:"letter_spacing,omitempty"`
	// LineHeight 行高，例如：1.3
	LineHeight float64 `json:"line_height,omitempty"`
	// TextAlign 对齐方式，可选值：
	// left(左对齐)
	// right(左对齐)
	// center(居中)
	// 默认为左对齐
	TextAlign TextAlign `json:"text_align,omitempty"`
}

// Type implement Brick interface
func (b RichTextBrick) Type() BrickType {
	return XrRichText
}
