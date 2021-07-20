package conversion

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
)

// Request 转化回传参数
type Request struct {
	// EventType 回传的事件，例如”激活“、”付费“。详细列表见附录
	EventType   enum.ConversionEventType `json:"event_type,omitempty"`
	EventWeight float64                  `json:"event_weight,omitempty"` // 广告主可以针对每一次转化，回传一个这次转化的价值，单位是 RMB。
	Context     *Context                 `json:"context,omitempty"`      // 包含一些关键的上下文信息
	Properties  *Properties              `json:"properties,omitempty"`   // 对于上报事件的附加属性，详情请见自定义属性的介绍。
	Timestamp   int64                    `json:"timestamp,omitempty"`    // 事件发生的毫秒级时间戳
}

// Context 包含一些关键的上下文信
type Context struct {
	Ad     *ContextAd     `json:"ad,omitempty"`     // 包含一些关键的广告相关信息
	Device *ContextDevice `json:"device,omitempty"` // 传递一些归因的设备信息
}

// ContextAd
type ContextAd struct {
	Callback  string                     `json:"callback,omitempty"`   // callback 字段有两个获取途径，对于监测链接归因的方式，需要从监测链接的__CALLBACK_PARAM__这个宏获取这个字段值；对于落地页或小程序归因的方式，需要从 url 中的 clickid 参数获取值。
	MatchType enum.ConversionAdMatchType `json:"match_type,omitempty"` // 这个参数可以帮助广告主标识是通过哪种渠道进行的归因，如果不传，就会默认为点击归因。
}

// ContextDevice
type ContextDevice struct {
	Imei string `json:"imei,omitempty"` // 归因上的设备 imei 的 MD5 值
	Idfa string `json:"idfa,omitempty"` // 归因上的设备的 idfa 的原值
	Oaid string `json:"oaid,omitempty"` // 归因上的设备的 oaid 的原值
}

// Properties
type Properties struct {
	PayAmount     float64                      `json:"pay_amount,omitempty"`
	BillTrackType enum.ConversionBillTrackType `json:"bill_track_type,omitempty"`
}

func (r Request) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
