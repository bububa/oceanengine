package video

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CheckAvailableAnchorRequest 查询视频是否挂载下载类锚点 API Request
type CheckAvailableAnchorRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ItemIDs 抖音视频ID ，通过【获取抖音主页视频】 或 【获取抖音授权关系】获取主页视频或合作授权视频
	ItemIDs []string `json:"item_ids,omitempty"`
	// LandingApp 推广目的 可选值:APP: APP
	LandingApp enum.LandingType `json:"landing_app,omitempty"`
	// ExternalAction 转化目标 可选值:
	// AD_CONVERT_TYPE_ACTIVE: 激活、AD_CONVERT_TYPE_ACTIVE_REGISTER: 注册、AD_CONVERT_TYPE_APP_ORDER: app内下单、AD_CONVERT_TYPE_APP_UV: app内访问、AD_CONVERT_TYPE_APP_PAY: app内付费、AD_CONVERT_TYPE_PAY: 激活且付费
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
}

// Encode implement GetRequest interface
func (r CheckAvailableAnchorRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	buf, _ := json.Marshal(r.ItemIDs)
	values.Set("item_ids", string(buf))
	values.Set("landing_app", string(r.LandingApp))
	values.Set("external_action", string(r.ExternalAction))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CheckAvailableAnchorResponse 查询视频是否挂载下载类锚点 API Response
type CheckAvailableAnchorResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List .
		List []AvailableAnchor `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// AvailableAnchor 下载类锚点
type AvailableAnchor struct {
	// ItemID 抖音视频ID
	ItemID string `json:"item_id,omitempty"`
	// Valid 是否支持双库存
	Valid bool `json:"valid,omitempty"`
}
