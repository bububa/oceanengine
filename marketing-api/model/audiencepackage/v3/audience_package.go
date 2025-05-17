package v3

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model/audiencepackage"
)

// AudiencePackage 定向包信息
type AudiencePackage struct {
	audiencepackage.AudienceBase
	// FilterAwemeAbnormalActive （抖音号推广特有）过滤高活跃用户; 取值：0表示不过滤，1表示过滤
	FilterAwemeAbnormalActive enum.OnOff `json:"filter_aweme_abnormal_active,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InterestActionMode 行为兴趣选择，枚举值：
	// UNLIMITED（不限），CUSTOM（自定义），RECOMMEND（系统推荐）
	InterestActionMode enum.InterestActionMode `json:"interest_action_mode,omitempty"`
	// AudiencePackageID 定向包id
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	// Name 定向包名称
	Name string `json:"name,omitempty"`
	// Description 定向包描述
	Description string `json:"description,omitempty"`
	// LandingType 定向包推广类型
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// AdType 广告类型，允许值：
	// ALL 所有广告（默认值）
	// SEARCH 搜索广告
	// 搜索定向包仅支持落地页、应用推广、抖音号、直播间，不支持商品、电商店铺、快应用、小游戏
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// MarketingGoal 营销场景：VIDEO_AND_IMAGE 短视频/图片， LIVE 直播
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// DeliveryRange 广告投放范围【附录：广告投放范围】
	DeliveryRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
	// HideIfExists 已安装用户，0表示不限，1表示过滤，2表示定向；过滤表示投放时不给安装客户展示广告，支持应用推广；定向表示投放时给安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；定向仅对Android链接生效。
	// 过滤已安装
	// UNLIMITED不限（默认值）
	// FILTER 过滤，仅安卓应用推广时支持，其他情况传入不生效
	// TARGETING 定向
	HideIfExists enum.HideIfExists `json:"hide_if_exists,omitempty"`
	// FiltetEvent 自定义过滤事件
	// 枚举值：
	// 优化目标：AD_CONVERT_EXTERNAL_ACTION
	// 激活：AD_CONVERT_TYPE_ACTIVE
	// 注册：AD_CONVERT_TYPE_ACTIVE_REGISTER
	// 付费：AD_CONVERT_TYPE_PAY
	FilterEvent []enum.AudiencePackageFiltetEvent `json:"filter_event,omitempty"`
	// HideIfConverted 过滤已转化
	// 枚举值：NO_EXCLUDE 不限制、PROMOTION 广告、PROJECT 项目、ADVERTISER 广告账户、APP 应用、CUSTOMER 客户、ORGANIZATION 组织（仅加白广告主可用）
	HideIfConverted enum.HideIfConverted `json:"hide_if_converted,omitempty"`
	// ConvertedTimeDuration 过滤时间范围，详见， 【附录-过滤时间范围】，当过滤已转化用户类型选择"公司账户"时填写，教育行业专属，非教育行业不会生效
	ConvertedTimeDuration enum.ConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	// CarrierRegionOptimize 运营商号段开关，枚举值：ON 开启 / OFF 关闭
	CarrierRegionOptimize enum.OnOff `json:"carrier_region_optimize,omitempty"`
}
