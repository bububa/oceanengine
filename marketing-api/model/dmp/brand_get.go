package dmp

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// BrandGetRequest 获取广告账户关联云图账户信息 API Request
type BrandGetRequest struct {
	// AdvertiserID 广告账户d
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r BrandGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}

// BrandGetResponse 获取广告账户关联云图账户信息 API Response
type BrandGetResponse struct {
	model.BaseResponse
	Data struct {
		// BrandInfo 云图账户信息列表
		BrandInfo []BrandInfo `json:"brand_info,omitempty"`
	} `json:"data,omitempty"`
}

// BrandInfo 云图账户信息
type BrandInfo struct {
	// BrandID 云图账户id
	BrandID uint64 `json:"brand_id,omitempty"`
	// BrandName 云图账户名
	BrandName string `json:"brand_name,omitempty"`
	// VirtualAdvID 虚拟adv_id
	VirtualAdvID uint64 `json:"virtual_adv_id,omitempty"`
}
