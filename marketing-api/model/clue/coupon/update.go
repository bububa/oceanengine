package coupon

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 编辑卡券
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ActivityID 青鸟营销活动ID
	ActivityID uint64 `json:"activity_id,omitempty"`
	// DeliverStart 活动开始时间，格式：%Y-%m-%d
	// 小于结束时间，最多往前调一年
	// 若小于当天则立即生效
	// 活动开始后不允许修改
	DeliverStart string `json:"deliver_start,omitempty"`
	// DeliverEnd 活动结束时间，格式：%Y-%m-%d，
	// 大于当前时间，最多往后调一年
	// deliver_end <= valid_end
	// valid_end为优惠券的有效期结束时间
	// 活动开始后不允许修改
	DeliverEnd string `json:"deliver_end,omitempty"`
	// GlobalLimit 总库存量限制
	// 活动开始后不允许修改
	GlobalLimit *LimitSetting `json:"global_limit,omitempty"`
	// UserLimit 单个用户的库存量限制
	// 活动开始后不允许修改
	UserLimit *LimitSetting `json:"user_limit,omitempty"`
	// Status NORMAL：上线; OFFLINE： 下线
	Status CouponStatus `json:"status,omitempty"`
}

// Encode implemnent PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
