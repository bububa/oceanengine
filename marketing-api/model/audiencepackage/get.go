package audiencepackage

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取定向包 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤字段
	Filtering *GetFiltering `json:"filtering,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量
	PageSize int `json:"page_size,omitempty"`
}

// GetFiltering 过滤字段
type GetFiltering struct {
	// LandingType 定向包类型
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// DeliveryRange 广告投放范围
	DeliveryRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
	// AdType 广告类型，允许值：
	// ALL 所有广告（默认值）
	// SEARCH 搜索广告
	// 搜索定向包仅支持落地页、应用推广、抖音号、直播间，不支持商品、电商店铺、快应用、小游戏
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// MarketingGoal 营销场景，允许值：VIDEO_AND_IMAGE 短视频/图片（默认值）， LIVE 直播
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 获取定向包 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AudiencePackages 定向包信息
	AudiencePackages []AudiencePackage `json:"audience_packages,omitempty"`
}
