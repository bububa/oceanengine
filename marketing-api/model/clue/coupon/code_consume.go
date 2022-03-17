package coupon

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CodeConsumeRequest 核销券码
type CodeConsumeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CodeID 券码ID，c端获取
	CodeID string `json:"code_id,omitempty"`
	// Code 券码code，c端获取
	// code_id和code至少设置一个，两个都设置code_id优先
	// 当设置code时，必须设置activity_id、coupon_id和resource_id中的一个
	Code string `json:"code,omitempty"`
	// ActivityID 线索通活动ID，可从查询券码记录接口获取
	ActivityID uint64 `json:"activity_id,omitempty"`
	// CouponID 卡券ID，可从查询券码记录接口获取
	CouponID uint64 `json:"coupon_id,omitempty"`
	// ResourceID 优惠券ID，可从查询券码记录接口获取
	ResourceID uint64 `json:"resource_id,omitempty"`
	// Employee 核销员信息
	Employee *Employee `json:"employee,omitempty"`
	// Extra 核销补充信息，即自定义信息
	Extra map[string]string `json:"extra,omitempty"`
}

// Encode implemnent PostRequest interface
func (r CodeConsumeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
