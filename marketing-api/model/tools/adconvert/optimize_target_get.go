package adconvert

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OptimizeTargetGetRequest 查询广告计划可用优化目标 API Request
type OptimizeTargetGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LandingType 广告组推广目的
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// MarketingPurpose 营销目的，可选值：UNLIMITED不限，CONVERSION行动转化， INTENTION用户意向，ACKNOWLEDGE品牌认知
	MarketingPurpose enum.MarketingPurpose `json:"marketing_purpose,omitempty"`
	// PromotionContent 投放内容，根据不同推广目的对应的不同的投放内容，详情可参考下方的【联动关系】对照表格
	// 允许值:
	// AWEME_HOME_PAGE、DOUYIN、DOWNLOAD_URL、EXTERNAL_URL、GOODS_LINK、LIVE_ROOM、MICRO_APP、NORMAL、QUICK_APP_URL、SHOP、THIRD_PARTY
	PromotionContent enum.PromotionContent `json:"promotion_content,omitempty"`
	// ExternalURL 落地页链接
	ExternalURL string `json:"external_url,omitempty"`
	// PackageName Android应用包名
	PackageName string `json:"package_name,omitempty"`
	// AppType 应用下载类型
	AppType string `json:"app_type,omitempty"`
	// ItunesURL iOS应用下载链接
	ItunesURL string `json:"itunes_url,omitempty"`
	// AppSchema 小程序app_schema
	AppSchema string `json:"app_schema,omitempty"`
	// CampaignType 广告组类型，允许值：FEED信息流广告，SEARCH搜索广告
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// ConvertType 跟踪方式
	ConvertType enum.AdConvertSource `json:"convert_type,omitempty"`
	// ConvertID 自定义转化id
	ConvertID uint64 `json:"convert_id,omitempty"`
	// DeepExternalAction 深度转化目标
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
	// ConvertName 转化名称
	ConvertName string `json:"convert_name,omitempty"`
	// DedicateType IOS14.5专属广告入参标识，可选值："UNSET"、"DEDICATED"，查询IOS14.5专属广告可用转化目标时必填"DEDICATED"
	DedicateType string `json:"dedicate_type,omitempty"`
	// LaunchTargetType 投放类型，允许值：LIVE_CONVERT 直播间转化、APP 应用下载、EXTERNAL 线索收集
	LaunchTargetType enum.LaunchTargetType `json:"launch_target_type,omitempty"`
	// Page 页数
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r OptimizeTargetGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Add("landing_type", string(r.LandingType))
	values.Add("marketing_purpose", string(r.MarketingPurpose))
	if r.PromotionContent != "" {
		values.Add("promotion_content", string(r.PromotionContent))
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
	if r.CampaignType != "" {
		values.Add("campaign_type", string(r.CampaignType))
	}
	if r.ConvertType != "" {
		values.Add("convert_type", string(r.ConvertType))
	}
	if r.ConvertID > 0 {
		values.Add("convert_id", strconv.FormatUint(r.ConvertID, 10))
	}
	if r.DeepExternalAction != "" {
		values.Add("deep_external_action", string(r.DeepExternalAction))
	}
	if r.ConvertName != "" {
		values.Add("convert_name", string(r.ConvertName))
	}
	if r.DedicateType != "" {
		values.Add("dedicate_type", string(r.DedicateType))
	}
	if r.LaunchTargetType != "" {
		values.Add("launch_target_type", string(r.LaunchTargetType))
	}
	if r.Page > 0 {
		values.Add("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Add("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OptimizeTargetGetResponse 查询广告计划可用优化目标 API Response
type OptimizeTargetGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List 转化数据列表
		List []OptimizeTarget `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
