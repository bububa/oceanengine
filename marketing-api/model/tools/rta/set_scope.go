package rta

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SetScopeRequest 设置账户下RTA策略生效范围 API Request
type SetScopeRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RtaID 预期设置的rta策略id
	RtaID uint64 `json:"rta_id,omitempty"`
	// TargetType 生效维度
	// 默认值:ADV广告账户
	// 允许值:ADV广告账户，CAMPAIGN广告组，PROJECT项目（体验版）
	TargetType enum.RtaTargetType `json:"target_type,omitempty"`
	// TargetIDs 生效列表，当target_type = CAMPAIGN 或 PROJECT 有效
	// 当target_type = CAMPAIGN，传入广告组id
	// 当target_type = PROJECT，传入项目id
	TargetIDs []uint64 `json:"target_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r SetScopeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
