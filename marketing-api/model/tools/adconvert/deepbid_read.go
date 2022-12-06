package adconvert

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeepbidReqdRequest 查询深度优化方式 API Request
type DeepbidReadRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// DeepExternalAction 深度转化目标
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
	// DeliveryRange 投放范围
	DeliveryRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
	// ConvertID 转化id，convert_id和external_action二选一
	ConvertID uint64 `json:"convert_id,omitempty"`
	// ExternalAction 转化类型，convert_id和external_action二选一
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// FlowControlMode 竞价策略(投放方式)
	FlowControlMode enum.FlowControlMode `json:"flow_control_mode,omitempty"`
	// SmartBidType 投放场景(出价方式)
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
}

// Encode implement GetRequest interface
func (r DeepbidReadRequest) Encode() string {
	values := util.GetUrlValues()
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Add("campaign_id", strconv.FormatUint(r.CampaignID, 10))
	values.Add("deep_external_action", string(r.DeepExternalAction))
	values.Add("delivery_range", string(r.DeliveryRange))
	if r.ConvertID > 0 {
		values.Add("convert_id", strconv.FormatUint(r.ConvertID, 10))
	}
	if r.ExternalAction != "" {
		values.Add("external_action", string(r.ExternalAction))
	}
	values.Add("flow_control_mode", string(r.FlowControlMode))
	values.Add("smart_bid_type", string(r.SmartBidType))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DeepbidReadResponse 查询深度优化方式 API Response
type DeepbidReadResponse struct {
	model.BaseResponse
	Data struct {
		// SuccessList 可用的深度转化方式列表
		SuccessList []enum.DeepBidType `json:"success_list,omitempty"`
	} `json:"data,omitempty"`
}
