package audiencepackage

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建定向包 API Request
type CreateRequest struct {
	AudienceBase
	// FilterAwemeAbnormalActive （抖音号推广特有）过滤高活跃用户; 取值：0表示不过滤，1表示过滤
	FilterAwemeAbnormalActive *int `json:"filter_aweme_abnormal_active,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 定向包名称
	Name string `json:"name,omitempty"`
	// Description 定向包描述
	Description string `json:"description,omitempty"`
	// LandingType 定向包推广类型
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// DeliveryRange 广告投放范围【附录：广告投放范围】
	DeliveryRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
	// AdType 广告类型，允许值：
	// ALL 所有广告（默认值）
	// SEARCH 搜索广告
	// 搜索定向包仅支持落地页、应用推广、抖音号、直播间，不支持商品、电商店铺、快应用、小游戏
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// MarketingGoal 营销场景，允许值：VIDEO_AND_IMAGE 短视频/图片（默认值）， LIVE 直播
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// HideIfExists 已安装用户，0表示不限，1表示过滤，2表示定向；过滤表示投放时不给安装客户展示广告，支持应用推广；定向表示投放时给安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；定向仅对Android链接生效。
	HideIfExists int `json:"hide_if_exists,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建定向包 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	} `json:"data,omitempty"`
}
