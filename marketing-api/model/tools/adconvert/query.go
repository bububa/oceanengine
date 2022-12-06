package adconvert

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryRequest 查询广告计划可用转化目标 API Request
type QueryRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LandingType 广告组推广目的
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// PromotionContent 投放内容，根据不同推广目的对应的不同的投放内容，详情可参考下方的【联动关系】对照表格
	// 允许值:
	// AWEME_HOME_PAGE、DOUYIN、DOWNLOAD_URL、EXTERNAL_URL、GOODS_LINK、LIVE_ROOM、MICRO_APP、NORMAL、QUICK_APP_URL、SHOP、THIRD_PARTY
	PromotionContent enum.PromotionContent `json:"promotion_content,omitempty"`
	// DeliveryRange 广告投放范围
	DeliveryRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
	// ExternalURL 落地页链接
	ExternalURL string `json:"external_url,omitempty"`
	// AppType 应用下载类型
	AppType string `json:"app_type,omitempty"`
	// PackageName Android应用包名
	PackageName string `json:"package_name,omitempty"`
	// ItunesURL iOS应用下载链接
	ItunesURL string `json:"itunes_url,omitempty"`
	// AppSchema 小程序app_schema
	AppSchema string `json:"app_schema,omitempty"`
	// AdvancedCreativeType 附加创意类型
	AdvancedCreativeType enum.AdvancedCreativeType `json:"advanced_creative_type,omitempty"`
	// MarketingScene 游戏预约场景，附加创意类型为ATTACHED_CREATIVE_GAME_SUBSCRIBE游戏预约时填写，允许值:
	// GAME_PROMOTION（游戏大推）、GAME_SUBSCRIBE（游戏预约 ）、NORMAL（普通场景）
	MarketingScene enum.MarketingScene `json:"marketing_scene,omitempty"`
}

// Encode implement GetRequest interface
func (r QueryRequest) Encode() string {
	values := util.GetUrlValues()
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Add("landing_type", string(r.LandingType))
	values.Add("promotion_content", string(r.PromotionContent))
	if r.DeliveryRange != "" {
		values.Add("delivery_range", string(r.DeliveryRange))
	}
	if r.ExternalURL != "" {
		values.Add("external_url", r.ExternalURL)
	}
	if r.AppType != "" {
		values.Add("app_type", r.AppType)
	}
	if r.PackageName != "" {
		values.Add("package_name", r.PackageName)
	}
	if r.ItunesURL != "" {
		values.Add("itunes_url", r.ItunesURL)
	}
	if r.AppSchema != "" {
		values.Add("app_schema", r.AppSchema)
	}
	if r.AdvancedCreativeType != "" {
		values.Add("advanced_creative_type", string(r.AdvancedCreativeType))
	}
	if r.MarketingScene != "" {
		values.Add("marketing_scene", string(r.MarketingScene))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryResponse 查询广告计划可用转化目标 API Response
type QueryResponse struct {
	model.BaseResponse
	Data struct {
		// List 转化数据列表
		List []AdConvert `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
