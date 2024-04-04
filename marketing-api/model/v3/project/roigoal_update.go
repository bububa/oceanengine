package project

import "github.com/bububa/oceanengine/marketing-api/util"

// RoiGoalUpdateRequest 批量修改项目ROI系数 API Request
type RoiGoalUpdateRequest struct {
	// AdvertiserID 广告主账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改项目ROI系数，list长度限制10
	Data []RoiGoalUpdateData `json:"data,omitempty"`
}

// RoiGoalUpdateData 批量修改项目ROI系数
type RoiGoalUpdateData struct {
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// Roigoal 深度转化ROI系数，填写要求如下：
	// 应用推广、自动投放项目，ROI系数范围[0.01,5]，仅支持最多4位小数
	// 电商推广、自动投放项目，ROI系数范围(0.01,100]，仅支持最多4位小数
	// 电商推广、自动投放、周期稳投项目，ROI系数只允许调小且新建7天内每自然日最多允许成功修改1次
	RoiGoal float64 `json:"roi_goal,omitempty"`
}

// Encode impllements PostRequest interface
func (r RoiGoalUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
