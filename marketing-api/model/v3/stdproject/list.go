package stdproject

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取标准项目列表 API Request
type ListRequest struct {
	// AdvertiserID 投放账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *ListFilter `json:"filtering,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10, 20, 50, 100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// ListFilter 过滤条件
type ListFilter struct {
	// ProjectIDs 项目IDs
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// StatusFirst 项目一级状态过滤，允许值：
	// PROJECT_STATUS_DELETE 已删除
	// PROJECT_STATUS_DONE 已完成
	// PROJECT_STATUS_DISABLE 未投放
	// PROJECT_STATUS_ENABLE 投放中
	// PROJECT_STATUS_FROZEN 已终止
	// ALL_EXCEPT_DELETE 不限
	// ALL 不限(包含已删除和已终止)
	// 默认值: ALL_EXCEPT_DELETE
	StatusFirst enum.StdProjectStatusFirst `json:"status_first,omitempty"`
	// StatusSecond 项目二级状态
	StatusSecond enum.StdProjectStatusSecond `json:"status_second,omitempty"`
	// AdType 营销类型，允许值：ALL 通投、SEARCH 搜索
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// MarketingGoal 营销场景，允许值：VIDEO_AND_IMAGE 短视频/图片、LIVE 直播
	MarketingGoal []enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// LandingType 营销目的，可选值：APP 应用、LINK 销售线索、MICRO_GAME 小程序、SHOP 电商
	LandingType []enum.LandingType `json:"landing_type,omitempty"`
	// BidType 竞价策略，允许值：CUSTOM 稳定成本、NO_BID 最大转化投放
	BidType []enum.BidType `json:"bid_type,omitempty"`
	// Platform 按平台过滤，允许值：ANDROID 安卓、HARMONY 鸿蒙、IOS IOS
	Platform []enum.AudiencePlatform `json:"platform,omitempty"`
	// StarDeliveryType 是否为星广联投项目，允许值：NOT_STAR_DELIVERY 非星广联投项目、STAR_DELIVERY 星广联投项目
	StarDeliveryType []string `json:"star_delivery_type,omitempty"`
	// BudgetMode 项目预算类型，允许值：BUDGET_MODE_INFINITE 不限、BUDGET_MODE_DAY 日预算、BUDGET_MODE_TOTAL 总预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// NativeType 投放身份，允许值：ACCOUNT 账户信息、AWEME 抖音号
	NativeType enum.NativeType `json:"native_type,omitempty"`
	// RejectReasonType 审核建议类型，允许值：NONE 无建议、SUGGEST_MODIFYING 建议修改、REVIEW_REJECT 审核不通过
	RejectReasonType string `json:"reject_reason_type,omitempty"`
	// Pricing 按计费方式过滤，可选值：
	// PRICING_CPC 按点击计费(CPC)
	// PRICING_CPM 按千次展示(CPM)
	// PRICING_OCPC 按转化优化，按点击计费
	// PRICING_OCPM 按转化优化，按展示计费
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// BlueFlowPackageSetting 是否投放蓝海流量，允许值：OFF、ON
	BlueFlowPackageSetting string `json:"blue_flow_package_setting,omitempty"`
	// ProjectCreateStartTime 项目创建开始时间，格式yyyy-mm-dd，表示过滤出当天创建的项目
	ProjectCreateStartTime string `json:"project_create_start_time,omitempty"`
	// ProjectCreateEndTime 项目创建结束时间，格式yyyy-mm-dd，表示过滤出当天创建的项目
	ProjectCreateEndTime string `json:"project_create_end_time,omitempty"`
	// ProjectModifyStartTime 项目更新开始时间，格式yyyy-mm-dd hh:mm:ss
	ProjectModifyStartTime string `json:"project_modify_start_time,omitempty"`
	// ProjectModifyEndTime 项目更新结束时间，格式yyyy-mm-dd hh:mm:ss
	ProjectModifyEndTime string `json:"project_modify_end_time,omitempty"`
	// AwemeID 抖音号ID
	AwemeID string `json:"aweme_id,omitempty"`
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
	// ProductPlatformID 商品库ID
	ProductPlatformID uint64 `json:"product_platform_id,omitempty"`
	// UniqueProductID 升级版商品库商品ID
	UniqueProductID uint64 `json:"unique_product_id,omitempty"`
	// InstanceID 字节小游戏/小程序、微信小游戏/小程序ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// BrandID 品牌ID
	BrandID uint64 `json:"brand_id,omitempty"`
	// ExternalAction 优化目标
	ExternalAction string `json:"external_action,omitempty"`
	// DeepExternalAction 深度优化目标
	DeepExternalAction string `json:"deep_external_action,omitempty"`
	// DeepBidType 深度优化方式
	DeepBidType string `json:"deep_bid_type,omitempty"`
	// GameAddictionID 关键行为ID
	GameAddictionID uint64 `json:"game_addiction_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
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

// ListResponse 获取标准项目列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResponseData `json:"data,omitempty"`
}

// ListResponseData 获取标准项目列表返回数据
type ListResponseData struct {
	// List 项目列表
	List []StdProject `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
