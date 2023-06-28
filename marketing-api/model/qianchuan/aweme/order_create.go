package aweme

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderCreateRequest 创建随心推订单 API Request
type OrderCreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MarketingGoal 营销目标，可选值:
	// LIVE_PROM_GOODS: 直播间带货
	// VIDEO_PROM_GOODS: 短视频带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// AwemeID 推广对象的抖音id，可通过【查询可推广抖音号列表】接口获取名下可推广抖音号
	// 1. 视频加热和视频加热直播间：传递视频作者的抖音uid
	// 2. 直接加热直播间：传递主播的抖音uid
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// AwemeItemID 音视频ID，可通过【随心推投放-获取随心推可投视频】接口获得视频， 视频加热和视频加热直播间：传递视频ID；
	// 1.当marketing_goal= VIDEO_PROM_GOODS时 必填
	// 2.当marketing_goal= LIVE_PROM_GOODS &liveroom_heat_mode= BY_VIDEO 时必填，不支持图文视频
	// 注意：当为图文视频时，仅支持短视频投放-商品购买/商品支付ROI场景
	AwemeItemID uint64 `json:"aweme_item_id,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
}

// Encode implement PostRequest interface
func (r OrderCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// OrderCreateResponse 创建随心推订单 API Response
type OrderCreateResponse struct {
	model.BaseResponse
	Data *Order `json:"data,omitempty"`
}
