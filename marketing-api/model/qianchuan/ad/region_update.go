package ad

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RegionUpdateRequest 更新计划地域定向 API Request
type RegionUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 需要更新的广告计划id，一次只允许传入10个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// District 地域类型，允许值：
	// CITY 省市
	// COUNTY 区县
	// NONE 不限，默认值
	District enum.District `json:"district,omitempty"`
	// City 具体定向的城市列表当 district 为COUNTY、CITY为必填，枚举值详见【附件-city.json】
	// 省市的传法："city" : [12], "district" : "CITY"
	// 区县的传法："city" : [130102], "district" : "COUNTY"
	City []uint64 `json:"city,omitempty"`
	// LocationType 地域定向的用户状态类型，当 district 为COUNTY，CITY为必填，允许值：
	// CURRENT 正在该地区的用户
	// HOME 居住在该地区的用户
	// TRAVEL 到该地区旅行的用户
	// ALL 该地区内的所有用户
	LocationType enum.LocationType `json:"location_type,omitempty"`
	// ExcludeLimitedRegion  排除限运地区，允许值：
	// 0：否，默认值
	// 1：是
	// 注意：
	// 1、当地域定向类型为“不限”/地域定向的用户状态类型为“正在该地区的用户”才支持
	ExcludeLimitedRegion int `json:"exclude_limited_region,omitempty"`
	// ElectricFenceRegion 电子围栏定向，可选值:
	// 1 不支持电子围栏地区
	// 2 支持电子围栏地区
	ElectricFenceRegion int `json:"electric_fence_region,omitempty"`
	// RegionVersion 行政区域版本号，可选值： 2.3.3 新版本
	RegionVersion string `json:"region_version,omitempty"`
}

// Encode implement PostRequest interface
func (r RegionUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
