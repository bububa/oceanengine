package track

import "github.com/bububa/oceanengine/marketing-api/enum"

type ActiveRequest struct {
	Link      string              `json:"link,omitempty"`       // 广告被打开的落地页的原始真实 url
	Callback  string              `json:"callback,omitempty"`   // 点击检测下发的 callback
	EventType enum.TrackEventType `json:"event_type,omitempty"` // 事件类型
	ConvTime  int64               `json:"conv_time,omitempty"`  // UTC 时间戳，单位：秒
	Source    string              `json:"source,omitempty"`     // 数据来源，比如来自 talkingdata的激活回调, 可以填写 TD
	MatchType enum.TrackMatchType `json:"match_type,omitempty"` // 归因方式
	Imei      string              `json:"imei,omitempty"`       // 安卓手机 imei 的 md5 摘要
	Idfa      string              `json:"idfa,omitempty"`       // ios 手机的 idfa 原值
	Muid      string              `json:"muid,omitempty"`       // 安卓：imei号取md5sum摘要；IOS：取idfa原值
	Oaid      string              `json:"oaid,omitempty"`       // Android Q 版本的 oaid 原值
	OaidMd5   string              `json:"oaid_md5,omitempty"`   // Android Q 版本的 oaid 原值的md5摘要
	Os        enum.TrackOS        `json:"os,omitempty"`         // 客户端的操作系统类型
	Caid1     string              `json:"caid1,omitempty"`      // 不同版本版本的中国广告协会互联网广告标识，caid1为最新版本，caid2为老版本
	Caid2     string              `json:"caid2,omitempty"`      // 不同版本版本的中国广告协会互联网广告标识，caid1为最新版本，caid2为老版本
	Ext       map[string]string   `json:"ext,omitempty"`        // 补充数据
}
