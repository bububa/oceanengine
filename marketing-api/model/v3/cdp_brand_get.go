package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CdpBrandGetRequest 获取关联云图的广告主账户信息 API Request
type CdpBrandGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r CdpBrandGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CdpBrandGetResponse 获取关联云图的广告主账户信息 API Response
type CdpBrandGetResponse struct {
	// Data
	Data *CdpBrandGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type CdpBrandGetResult struct {
	// BrandCategoryInfo 品牌所属类别
	BrandCategoryInfo *CdpBrandCategory `json:"brand_category_info,omitempty"`
	// BrandInfoData 品牌信息
	BrandInfoData []CdpBrandInfo `json:"brand_info_data,omitempty"`
}

// CdpBrandCategory 品牌所属类别
type CdpBrandCategory struct {
	// Children
	Children *CdpBrandCategory `json:"children,omitempty"`
	// ID 一级类别id: yuntu_category_id
	ID string `json:"id,omitempty"`
	// Label 一级类别标签
	Label string `json:"label,omitempty"`
}

// CdpBrandInfo 品牌信息
type CdpBrandInfo struct {
	// SubBrandMap 子品牌信息名称和id
	SubBrandMap *CdpBrandInfo `json:"sub_brand_map,omitempty"`
	// CdpBrandName cdp品牌信息
	CdpBrandName string `json:"cdp_brand_name,omitempty"`
	// CdpBrandID cdp品牌id
	CdpBrandID int64 `json:"cdp_brand_id,omitempty"`
	// EcomBrandID 电商品牌id
	EcomBrandID int64 `json:"ecom_brand_id,omitempty"`
	// BrandNameID 云图品牌id
	BrandNameID int64 `json:"brand_name_id,omitempty"`
}
