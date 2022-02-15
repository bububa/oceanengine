package site

// IconType 图标信息
type IconType string

const (
	// IconType_NONE 缺省
	IconType_NONE IconType = "none"
	// IconType_CONSULT 咨询图标
	IconType_CONSULT IconType = "consult"
	// IconType_TELEPHONE 电话图标
	IconType_TELEPHONE IconType = "telephone"
	// IconType_DOWNLOAD 下载图标
	IconType_DOWNLOAD IconType = "download"
)

// ButtonBgType 背景类型
type ButtonBgType string

const (
	// ButtonBgType_COLOR 默认值
	ButtonBgType_COLOR ButtonBgType = "color"
	// ButtonBgType_IMAGE 图片
	ButtonBgType_IMAGE ButtonBgType = "image"
)

// ButtonBrick 按钮组件
type ButtonBrick struct {
	BaseBrick
	// Icon 图标信息，icon类型值：
	// none(缺省)
	// consult（咨询图标）
	// telephone(电话图标)
	// download(下载图标)
	Icon IconType `json:"icon,omitempty"`
	// BgType 背景类型，可选值：color(默认值)，image(图片)
	BgType ButtonBgType `json:"bg_type,omitempty"`
	// Text 文案，bg_type 为color时，有效且必填
	Text string `json:"text,omitempty"`
	// Color 文案颜色信息，默认值rgba(255,255,255,1)
	Color string `json:"color,omitempty"`
	// FontSize 文案字号,默认值16
	FontSize int `json:"font_size,omitempty"`
	// BgColor 背景颜色信息，比如“rgba(255,255,255,1)”,默认为无
	BgColor string `json:"bg_color,omitempty"`
	// BgImage 背景图片，bg_type 为image时，有效且必填;可自行上传图片url，或从获取图片素材获取图片ID：id
	BgImage string `json:"bg_image,omitempty"`
	// BorderColor 边框颜色信息，比如“rgba(255,255,255,1)”,默认为无
	BorderColor string `json:"border_color,omitempty"`
	// BorderWidth 边框粗细，默认0
	BorderWidth int `json:"border_width,omitempty"`
	// BorderRadius 边框圆角，默认为0，范围：radius >= 0
	BorderRadius float64 `json:"border_radius,omitempty"`
	// Events 事件列表信息
	Events []ButtonEvent `json:"events,omitempty"`
}

// Type implement Brick interface
func (b ButtonBrick) Type() BrickType {
	return XrButton
}

// ButtonEvent 事件信息
type ButtonEvent struct {
	// Trigger 触发
	Trigger *EventTrigger `json:"trigger,omitempty"`
	// Behavior 事件行为描述，behavior目前已开放DownloadEvent 和LinkEvent以及TelephoneEvent
	Behavior *Event `json:"behavior,omitempty"`
}

// TriggerType 触发事件
type TriggerType string

const (
	// TriggerType_CLICK 点击
	TriggerType_CLICK TriggerType = "click"
	// TriggerType_LONG_TOUCH 长按
	TriggerType_LONG_TOUCH TriggerType = "longTouch"
)

// EventTrigger 事件触发条件
type EventTrigger struct {
	// Type 触发事件，可选值：
	// click（点击）
	// longTouch（长按）
	Type TriggerType `json:"type,omitempty"`
	// TouchSeconds 长按时间（秒）：触发时间为longTouch类型，必填;长按时间 >= 1
	TouchSeconds int64 `json:"touch_seconds,omitempty"`
}
