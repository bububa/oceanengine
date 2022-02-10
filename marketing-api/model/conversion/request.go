package conversion

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
)

// Request 转化回传参数
type Request struct {
	// EventType 回传的事件，例如”激活“、”付费“。详细列表见附录
	EventType enum.ConversionEventType `json:"event_type,omitempty"`
	// EventWeight 广告主可以针对每一次转化，回传一个这次转化的价值，单位是 RMB。
	EventWeight float64 `json:"event_weight,omitempty"`
	// Context 包含一些关键的上下文信息
	Context *Context `json:"context,omitempty"`
	// Properties 对于上报事件的附加属性，详情请见自定义属性的介绍。
	Properties *Properties `json:"properties,omitempty"`
	// Timestamp 事件发生的毫秒级时间戳
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Context 包含一些关键的上下文信
type Context struct {
	Ad     *ContextAd     `json:"ad,omitempty"`     // 包含一些关键的广告相关信息
	Device *ContextDevice `json:"device,omitempty"` // 传递一些归因的设备信息
}

// ContextAd 广告相关信息
type ContextAd struct {
	// Callback callback 字段有两个获取途径，对于监测链接归因的方式，需要从监测链接的__CALLBACK_PARAM__这个宏获取这个字段值；对于落地页或小程序归因的方式，需要从 url 中的 clickid 参数获取值。
	Callback string `json:"callback,omitempty"`
	// MatchType 这个参数可以帮助广告主标识是通过哪种渠道进行的归因，如果不传，就会默认为点击归因。
	MatchType enum.ConversionAdMatchType `json:"match_type,omitempty"`
}

// ContextDevice 归因设备信息
type ContextDevice struct {
	// Imei 归因上的设备 imei 的 MD5 值
	Imei string `json:"imei,omitempty"`
	// Idfa 归因上的设备的 idfa 的原值
	Idfa string `json:"idfa,omitempty"`
	// Oaid 归因上的设备的 oaid 的原值
	Oaid string `json:"oaid,omitempty"`
}

// Properties 附加属性
type Properties struct {
	// PayAmount 支付金额
	PayAmount float64 `json:"pay_amount,omitempty"`
	// BillTrackType 支付方式
	BillTrackType enum.ConversionBillTrackType `json:"bill_track_type,omitempty"`
}

// Encode implement GetRequest interface
func (r Request) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
