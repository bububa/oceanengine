package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 更新项目 API Request
type UpdateRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// SearchBidRatio 出价系数，默认系数为1，出价系数可通过【获取快投推荐出价系数】查询，小数点后最多两位,取值范围 [1,2]
	// 当符合以下所有条件时填写有效
	// 1. bid_type != NO_BID && pricing = PRICING_OCPM
	// 2. deep_bid_type = DEEP_BID_DEFAULT 无深度优化方式 /BID_PER_ACTION 每次付费
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// AudienceExtend 定向拓展, 允许值：ON:开启（默认值）， OFF:关闭
	AudienceExtend enum.OnOff `json:"audience_extend,omitempty"`
	// Keywords 搜索关键词列表
	Keywords *[]Keyword `json:"keywords,omitempty"`
	// DownloadMode 优先从系统应用商店下载（下载模式），枚举值：APP_STORE_DELIVERY 优先商店下载、 DEFAULT 默认下载
	DownloadMode enum.DownloadMode `json:"download_mode,omitempty"`
	// OpenURL Deeplink直达链接
	OpenURL string `json:"open_url,omitempty"`
	// UlinkURL ulink直达链接
	UlinkURL string `json:"ulink_url,omitempty"`
	// DpaCategories 商品投放范围，分类列表，由【DPA商品广告-获取DPA分类】 得到
	// 个数限制 [0, 1000]
	// 不传和传空数组即为不限商品投放范围
	// DPA推广目的下有效
	DpaCategories []uint64 `json:"dpa_categories,omitempty"`
	// DpaProductTarget 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。获取商品库元信息-商品广告-商业开放平台
	// 数组长度限制：最大5条
	// DPA推广目的下有效
	DpaProductTarget []DpaProductTarget `json:"dpa_product_target,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// TrackURLSetting 监测链接设置
	TrackURLSetting *TrackURLSetting `json:"track_url_setting,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
