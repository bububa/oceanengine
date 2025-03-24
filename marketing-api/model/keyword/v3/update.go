package v3

import (
	"github.com/bububa/oceanengine/marketing-api/model/keyword"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 更新关键词 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionID 待更新搜索关键词属性的广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// ProjectID 待创建搜索关键词的项目ID，仅支持搜索极速智投在项目层级设置关键词
	ProjectID uint64 `json:"project_id,omitempty"`
	// Keywords 待更新搜索关键词列表，一次最多批量更新100个关键词
	Keywords []keyword.Keyword `json:"keywords,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
