package track

import (
	"crypto/rsa"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/core"
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
	// Caid 中国广告协会互联网广告标识
	Caid string `json:"caid,omitempty"`
	// Ext 补充数据
	Ext map[string]string `json:"ext,omitempty"`
	// PrivateKey
	PrivateKey *rsa.PrivateKey `json:"-"`
	// Credential
	Credential enum.Credential `json:"-"`
	// AppAccessToken
	AppAccessToken string `json:"-"`
}

// Encode implement PostRequest interface
func (r ActiveRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// Sign implement ConvertionRequest interface
func (r ActiveRequest) Sign(req *http.Request, content []byte) (string, error) {
	if r.PrivateKey == nil || r.Credential == "" {
		return "", errors.New("no private_key/credential")
	}
	return model.CredentialSign(req, content, r.PrivateKey, r.Credential)
}

// GetAppAccessToken implement ConvertionRequest interface
func (r ActiveRequest) GetAppAccessToken() string {
	return r.AppAccessToken
}

// RequestURI implement TrackRequest interface
func (r ActiveRequest) RequestURI() string {
	return core.TRACK_URL
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
	// Token
	Token string `json:"-"`
	// Gateway
	Gateway string `json:"-"`
	// PrivateKey
	PrivateKey *rsa.PrivateKey `json:"-"`
	// Credential
	Credential enum.Credential `json:"-"`
	// AppAccessToken
	AppAccessToken string `json:"-"`
}

// Encode implement GetRequest interface
func (r WxaActiveRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// RequestURI implement TrackRequest interface
func (r WxaActiveRequest) RequestURI() string {
	values := util.GetUrlValues()
	values.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	values.Set("nonce", strconv.Itoa(rand.Intn(1000)))
	strs := make([]string, 0, 3)
	strs = append(strs, r.Token)
	strs = append(strs, values.Get("timestamp"))
	strs = append(strs, values.Get("nonce"))
	sort.Strings(strs)
	b := util.GetBufferPool()
	for _, s := range strs {
		b.WriteString(s)
	}
	h := sha1.New()
	h.Write(b.Bytes())
	values.Set("signature", hex.EncodeToString(h.Sum(nil)))
	util.PutBufferPool(b)
	link := util.StringsJoin(r.Gateway, "?", values.Encode())
	util.PutUrlValues(values)
	return link
}

// Sign implement ConvertionRequest interface
func (r WxaActiveRequest) Sign(req *http.Request, content []byte) (string, error) {
	if r.PrivateKey == nil || r.Credential == "" {
		return "", errors.New("no private_key/credential")
	}
	return model.CredentialSign(req, content, r.PrivateKey, r.Credential)
}

// GetAppAccessToken implement ConvertionRequest interface
func (r WxaActiveRequest) GetAppAccessToken() string {
	return r.AppAccessToken
}
