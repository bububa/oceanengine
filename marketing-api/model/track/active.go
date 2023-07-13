package track

import (
	"crypto/rsa"
	"errors"
	"net/http"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/conversion"
	"github.com/bububa/oceanengine/marketing-api/util"
)

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
	// PrivateKey
	PrivateKey *rsa.PrivateKey `json:"-"`
	// Credential
	Credential enum.Credential `json:"-"`
}

// WxaActiveRequest 微信小程序线索-API上报数据 API Request
type WxaActiveRequest struct {
	// ClueToken 下发线索token
	ClueToken string `json:"clue_token,omitempty"`
	// UnionID 微信union_id
	UnionID string `json:"union_id,omitempty"`
	// OpenID 微信open_id
	OpenID string `json:"open_id,omitempty"`
	// EventType 事件类型
	EventType string `json:"event_type,omitempty"`
	// Props 参数包含pay_amount
	Props *conversion.Properties `json:"props,omitempty"`
}

// Encode implement GetRequest interface
func (r ActiveRequest) Encode() string {
	values := util.GetUrlValues()
	if r.Callback != "" {
		values.Set("callback", r.Callback)
		values.Set("os", strconv.Itoa(r.Os))
		if r.Os == enum.Track_ANDROID {
			values.Set("imei", r.Imei)
			values.Set("muid", r.Muid)
			values.Set("oaid", r.Oaid)
			if r.OaidMd5 != "" {
				values.Set("oaid_md5", r.OaidMd5)
			}
		} else {
			values.Set("idfa", r.Idfa)
		}
		if r.Caid1 != "" {
			values.Set("caid1", r.Caid1)
		}
		if r.Caid2 != "" {
			values.Set("caid2", r.Caid2)
		}
	} else {
		values.Set("link", r.Link)
	}
	values.Set("event_type", strconv.Itoa(r.EventType))
	if r.ConvTime > 0 {
		values.Set("conv_time", strconv.FormatInt(r.ConvTime, 10))
	}
	if r.Source != "" {
		values.Set("source", r.Source)
	}
	for k, v := range r.Ext {
		values.Set(k, v)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// Sign implement ConvertionRequest interface
func (r ActiveRequest) Sign(req *http.Request, content []byte) (string, error) {
	if r.PrivateKey == nil || r.Credential == "" {
		return "", errors.New("no private_key/credential")
	}
	return model.CredentialSign(req, content, r.PrivateKey, r.Credential)
}
