package hotmaterialderive

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SubmitRequest 提交爆款裂变任务 API Request
type SubmitRequest struct {
	//  AdvertiserID 广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialIDs 爆款素材 ID，规则：需要有投放记录，不能是低质、低效素材等
	// 长度限制 1-50
	// 全量生效策略参数配置
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// Strategies 生效策略，默认为空全生效；支持项：
	// CLIP_REPLACE - 分镜替换
	// ROBOT_REPLACE - 人物替换
	// HOT_PRE_VIDEO - 爆款开头
	// MIX_CUT - 重新混剪
	// PRE_VIDEO_CLIP_REPLACE - 前贴扩写
	// DERIVE_FROM_CHOSEN_HOT_MID - 自有爆款套路 （目前仅千川账户支持）
	// DERIVE_FROM_INDUSTRY_HOT_PATTERN - 行业爆款套路（目前仅千川账户支持）
	Strategies []string `json:"strategies,omitempty"`
	// ExcludeStrategies 禁止策略，范围同生效策略，优先级高于生效策略（两者都有时不生效）
	ExcludeStrategies []string `json:"exclude_strategies,omitempty"`
}

// Encode implements PostRequest interface
func (r SubmitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SubmitResponse 提交爆款裂变任务 API Response
type SubmitResponse struct {
	model.BaseResponse
	Data struct {
		Tasks []Task `json:"tasks,omitempty"`
	} `json:"data,omitempty"`
}
