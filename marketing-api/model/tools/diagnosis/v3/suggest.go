package v3

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/tools/diagnosis"
)

// AdSuggestion 诊断建议
type AdSuggestion struct {
	// PromotionID 计划id
	PromotionID model.Uint64 `json:"promotion_id,omitempty"`
	// SceneList 计划对应的场景列表
	SceneList []diagnosis.SuggestScene `json:"scene_list,omitempty"`
}
