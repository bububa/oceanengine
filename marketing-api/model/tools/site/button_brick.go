package site

import (
	"encoding/json"
)

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
	FontSize float64 `json:"font_size,omitempty"`
	// BgColor 背景颜色信息，比如“rgba(255,255,255,1)”,默认为无
	BgColor string `json:"bg_color,omitempty"`
	// BgImage 背景图片，bg_type 为image时，有效且必填;可自行上传图片url，或从获取图片素材获取图片ID：id
	BgImage string `json:"bg_image,omitempty"`
	// BorderColor 边框颜色信息，比如“rgba(255,255,255,1)”,默认为无
	BorderColor string `json:"border_color,omitempty"`
	// BorderWidth 边框粗细，默认0
	BorderWidth float64 `json:"border_width,omitempty"`
	// BorderRadius 边框圆角，默认为0，范围：radius >= 0
	BorderRadius float64 `json:"border_radius,omitempty"`
	// Events 事件列表信息
	Events []ButtonEvent `json:"events,omitempty"`
	// PackageInfo 应用下载样式
	PackageInfo *PackageInfo `json:"package_info,omitempty"`
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
	Behavior Event `json:"behavior,omitempty"`
}

type tmpButtonEvent struct {
	// Trigger 触发
	Trigger *EventTrigger `json:"trigger,omitempty"`
	// Behavior 事件行为描述，behavior目前已开放DownloadEvent 和LinkEvent以及TelephoneEvent
	Behavior json.RawMessage `json:"behavior,omitempty"`
	Name     string          `json:"name,omitempty"`
	Method   EventType       `json:"method,omitempty"`
	Payload  json.RawMessage `json:"payload,omitempty"`
}

func (e *ButtonEvent) UnmarshalJSON(b []byte) error {
	var tmp tmpButtonEvent
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	event := ButtonEvent{
		Trigger: tmp.Trigger,
	}
	var (
		eventType EventType
		payload   json.RawMessage
	)
	if tmp.Payload != nil {
		eventType = tmp.Method
		payload = tmp.Payload
	} else {
		var base BaseEvent
		if err := json.Unmarshal(tmp.Behavior, &base); err != nil {
			return err
		}
		eventType = base.Type()
		payload = tmp.Behavior
	}
	switch eventType {
	case EventType_DownloadEvent:
		var data DownloadEvent
		if err := json.Unmarshal(payload, &data); err != nil {
			return err
		}
		event.Behavior = &data
	case EventType_LinkEvent:
		var data LinkEvent
		if err := json.Unmarshal(payload, &data); err != nil {
			return err
		}
		event.Behavior = &data
	case EventType_TelephoneEvent:
		var data TelephoneEvent
		if err := json.Unmarshal(payload, &data); err != nil {
			return err
		}
		event.Behavior = &data
	}
	*e = event
	return nil
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

// PackageInfo 应用下载样式
type PackageInfo struct {
	// Type 展示样式:imageText 图片文字、image 图片
	Type string `json:"type,omitepty"`
	// Image type为“image”样式时有效，可自行上传图片url，或从获取图片素材获取图片ID：id
	Image string `json:"image,omitempty"`
	// Background 背景信息，type为imageText样式时有效
	Background *Background `json:"background,omitempty"`
	// Description 应用包描述，type为imageText样式时有效
	Description *Description `json:"description,omitempty"`
	// Title 标题配置，type为imageText样式时有效
	Title *Title `json:"title,omitempty"`
}

// Background 背景信息，type为imageText样式时有效
type Background struct {
	// Type 背景类型：color 颜色、image 图片
	Type string `json:"type,omitempty"`
	// Color 景颜色信息，比如“rgba(255,255,255,1)”,默认为无
	Color string `json:"color,omitempty"`
	// Image 背景图片，background的type 为image时，可自行上传图片url，或从获取图片素材获取图片ID：id
	Image string `json:"image,omitempty"`
}

// Description 应用包描述，type为imageText样式时有效
type Description struct {
	// Type 描述展示类型:label 标签、text 文本
	Type string `json:"type,omitempty"`
	// Labels 标签内容，description的type为label时有效
	Labels []string `json:"labels,omitempty"`
	// Color 文字颜色，比如“rgba(255,255,255,1)”,默认为无
	Color string `json:"color,omitempty"`
	// FontSize 文字大小，默认15
	FontSize float64 `json:"font_size,omitempty"`
}

// Title 标题配置，type为imageText样式时有效
type Title struct {
	// Color 文字颜色，比如“rgba(255,255,255,1)”,默认为无
	Color string `json:"color,omitempty"`
	// FontSize 文字大小，默认15
	FontSize float64 `json:"font_size,omitempty"`
}
