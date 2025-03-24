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
	// AigcDynamicCreativeSwitch
	// 是否开启AIGC动态创意，系统将根据创意生态要求生成素材探索，不传表示不开启，枚举值：
	// ON 开启
	// OFF 关闭
	// 注意：
	// 该功能仅支持行业白名单客户使用，如需使用可咨询对接销售/运营
	// 必须关联产品投放，才可设置该开关（即必须传入related_product结构体）
	// 如投放短剧商品，需要开启AIGC动态创意，请确保：
	// 关联短剧商品内包含字节小程序推广链接
	// 短剧剧目已获取原片使用权，可通过「查询短剧商品原片授权申请状态」接口查询授权申请状态
	AigcDynamicCreativeSwitch enum.OnOff `json:"aigc_dynamic_creative_switch,omitempty"`
	// SearchBidRatio 出价系数，默认系数为1，出价系数可通过【获取快投推荐出价系数】查询，小数点后最多两位,取值范围 [1,2]
	// 当符合以下所有条件时填写有效
	// 1. bid_type != NO_BID && pricing = PRICING_OCPM
	// 2. deep_bid_type = DEEP_BID_DEFAULT 无深度优化方式 /BID_PER_ACTION 每次付费
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// AudienceExtend 定向拓展, 允许值：ON:开启（默认值）， OFF:关闭
	AudienceExtend enum.OnOff `json:"audience_extend,omitempty"`
	// Keywords 搜索关键词列表
	Keywords *[]Keyword `json:"keywords,omitempty"`
	// AutoTraficExtend 智能拓流 ，允许值：ON开启； OFF关闭
	// 仅支持ad_type = SEARCH时设置
	// 自定义关键词和智能拓流二者必须开启一个：若keywords为空，智能拓流 auto_extend_traffic需为ON
	// 对于搜索极速智投项目，若设置blue_flow_keyword_name蓝海关键词，智能拓流默认值为ON，且不得设置为OFF
	AutoExtendTraffic enum.OnOff `json:"auto_extend_traffic,omitempty"`
	// BlueFlowPackage 搜索蓝海流量投放相关参数
	BlueFlowPackage *BlueFlowPackage `json:"blue_flow_package,omitempty"`
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
	// RelatedProduct 关联产品投放相关参数
	// 该参数结构为非必填，不传即视为原关联产品无改动
	RelatedProduct *RelatedProduct `json:"related_product,omitempty"`
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
