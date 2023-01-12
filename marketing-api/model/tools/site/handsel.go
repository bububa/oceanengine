package site

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// HandselRequest 建站工具-建站转赠 API Request
type HandselRequest struct {
	// AdvertiserID 广告主id，数字范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteIDs 站点id列表, min_size=1, max_size=20
	SiteIDs []string `json:"site_ids,omitempty"`
	// TargetAdvertiserIDs 目标广告主id列表, min_size=1, max_size=20
	TargetAdvertiserIDs []string `json:"target_advertiser_ids,omitempty"`
}

// implement PostRequest interface
func (r HandselRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// HandselResponse 建站工具-建站转赠 API Response
type HandselResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *HandselResponseData `json:"data,omitempty"`
}

// HandselResponseData json 返回值
type HandselResponseData struct {
	// SuccessList 转赠成功列表，整体失败不返回该列表
	SuccessList []HandselResult `json:"success_list,omitempty"`
	// ErrorList 转赠失败列表，整体成功不返回该列表
	ErrorList []HandselResult `json:"error_list,omitempty"`
}

// HandselResult 转赠结果
type HandselResult struct {
	// OriginSiteID 返回转赠成功的原site_ids
	OriginSiteID model.Uint64 `json:"origin_site_id,omitempty"`
	// SiteID 转赠成功后的生成的新站点id，失败的数据无此参数返回
	SiteID model.Uint64 `json:"site_id,omitempty"`
	// TargetAdvertiserID 返回转赠成功的目标广告主id
	TargetAdvertiserID model.Uint64 `json:"target_advertiser_id,omitempty"`
}
