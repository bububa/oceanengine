package site

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CopyRequest 建站工具-建站复制 API Request
type CopyRequest struct {
	// AdvertiserID 广告主id，数字范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteIDs 站点id列表, min_size=1, max_size=20
	SiteIDs []string `json:"site_ids,omitempty"`
}

// implement PostRequest interface
func (r CopyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CopyResponse 建站工具-建站复制 API Response
type CopyResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CopyResponseData `json:"data,omitempty"`
}

// CopyResponseData json 返回值
type CopyResponseData struct {
	// SuccessList 复制成功列表，整体失败不返回该列表
	SuccessList []CopyResult `json:"success_list,omitempty"`
	// ErrorList 复制失败列表，整体成功不返回该列表
	ErrorList []CopyResult `json:"error_list,omitempty"`
}

// CopyResult 复制结果
type CopyResult struct {
	// OriginSiteID 返回复制成功的原站点id
	OriginSiteID model.Uint64 `json:"origin_site_id,omitempty"`
	// SiteID 返回复制成功后生成的新站点id
	SiteID model.Uint64 `json:"site_id,omitempty"`
}
