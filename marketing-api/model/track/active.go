package track

import "github.com/bububa/oceanengine/marketing-api/enum"

// ActiveRequest 线索-API上报数据 API Request
type ActiveRequest struct {
	// Link 广告被打开的落地页的原始真实 url
	Link string `json:"link,omitempty"`
	// Callback 点击检测下发的 callback
	Callback string `json:"callback,omitempty"`
	// EventType 事件类型
	EventType enum.TrackEventType `json:"event_type,omitempty"`
	// ConvTime UTC 时间戳，单位：秒
	ConvTime int64 `json:"conv_time,omitempty"`
	// Source 数据来源，比如来自 talkingdata的激活回调, 可以填写 TD
	Source string `json:"source,omitempty"`
	// MatchType 归因方式
	MatchType enum.TrackMatchType `json:"match_type,omitempty"`
	// Imei 安卓手机 imei 的 md5 摘要
	Imei string `json:"imei,omitempty"`
	// Idfa ios 手机的 idfa 原值
	Idfa string `json:"idfa,omitempty"`
	// Muid 安卓：imei号取md5sum摘要；IOS：取idfa原值
	Muid string `json:"muid,omitempty"`
	// Oaid Android Q 版本的 oaid 原值
	Oaid string `json:"oaid,omitempty"`
	// OaidMd5 Android Q 版本的 oaid 原值的md5摘要
	OaidMd5 string `json:"oaid_md5,omitempty"`
	// Os 客户端的操作系统类型
	Os enum.TrackOS `json:"os,omitempty"`
	// Caid1 不同版本版本的中国广告协会互联网广告标识，caid1为最新版本，caid2为老版本
	Caid1 string `json:"caid1,omitempty"`
	// Caid2 不同版本版本的中国广告协会互联网广告标识，caid1为最新版本，caid2为老版本
	Caid2 string `json:"caid2,omitempty"`
	// Ext 补充数据
	Ext map[string]string `json:"ext,omitempty"`
}
