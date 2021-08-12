package event

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// AssetsGetRequest 获取推广内容API Request
type AssetsGetRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetType 资产类型，允许值：THIRD_EXTERNAL：三方落地页
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值10，最大30
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤条件
	Filtering *AssetsGetFiltering `json:"filtering,omitempty"`
}

// AssetsGetFiltering 过滤条件
type AssetsGetFiltering struct {
	// LandingPage 三方落地页数据
	LandingPage *LandingPage `json:"landing_page,omitempty"`
}

// Encode implement GetRequest interface
func (r AssetsGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("asset_type", string(r.AssetType))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	return values.Encode()
}

// AssetsGetResponse 获取推广内容 API Response
type AssetsGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AssetsGetResponseData `json:"data,omitempty"`
}

// AssetsGetResponseData json返回值
type AssetsGetResponseData struct {
	// LandingPages 三方数据集合
	LandingPages []LandingPage `json:"landing_pages,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
