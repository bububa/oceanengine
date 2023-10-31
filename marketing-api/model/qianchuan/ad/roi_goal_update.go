package ad

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RoiGoalUpdateRequest 更新计划的支付ROI目标 API Request
type RoiGoalUpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RoiGoalUpdates 更新计划出价的列表，当前最多支持1个
	RoiGoalUpdates []RoiGoalUpdate `json:"roi_goal_updates,omitempty"`
}

// RoiGoalUpdate 更新计划出价的列表，当前最多支持1个
type RoiGoalUpdate struct {
	// AdID 需要更新ROI目标的计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// RoiGoal 计划更新之后的支付ROI目标，最多支持两位小数，0.01～100
	// 注意：
	// 按展示付费(oCPM)，根据【保障规则】提供保障福利，请谨慎修改支付ROI目标和定向，以免失去保障资格。
	RoiGoal float64 `json:"roi_goal,omitempty"`
}

// Encode implement PostRequest interface
func (r RoiGoalUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// RoiGoalUpdateResponse 更新计划的支付ROI目标 API Response
type RoiGoalUpdateResponse struct {
	model.BaseResponse
	Data struct {
		Results []RoiGoalUpdateResult `json:"results,omitempty"`
	} `json:"data,omitempty"`
}

type RoiGoalUpdateResult struct {
	// Flag 更新ROI目标结果，枚举值:
	// 1 成功
	// 0 失败
	Flat int `json:"flat,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// ErrorMessage 更新ROI目标失败的原因
	// 仅当"flag"为0 时有效
	ErrorMessage string `json:"error_message,omitempty"`
}
