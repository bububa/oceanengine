package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailGetRequest 获取计划详情（含创意信息） API Request
type DetailGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// RequestMaterialURL 是否需要返回素材url，只有当version=v2时有效
	RequestMaterialURL string `json:"request_material_url,omitempty"`
	// Version 版本，可选值：v2，当传入v2值，支持获取素材预览url，同时不会返回程序化创意叉乘后的创意信息，返回内容与pc一致
	Version string `json:"version,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	if r.RequestMaterialURL != "" {
		values.Set("request_material_url", r.RequestMaterialURL)
	}
	if r.Version != "" {
		values.Set("version", r.Version)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailGetResponse 获取计划详情（含创意信息） API Response
type DetailGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *Ad `json:"data,omitempty"`
}
