package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 新建全域推广计划 API Request
type CreateRequest struct {
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音号id
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// MarketingGoal 营销目标，可选值:
	// LIVE_PROM_GOODS 直播间带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// CreativeSetting 创意设置
	CreativeSetting *CreativeSetting `json:"creative_setting,omitempty"`
	// ProgrammaticCreativeMediaList 程序化创意信息
	ProgrammaticCreativeMediaList []ProgrammaticCreativeMedia `json:"programmatic_creative_media_list,omitempty"`
}

// Encode implements PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建全域推广计划 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// AdID 计划id
		AdID uint64 `json:"ad_id,omitempty"`
	} `json:"data,omitempty"`
}
