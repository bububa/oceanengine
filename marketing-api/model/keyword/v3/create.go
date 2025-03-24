package v3

import (
	"github.com/bububa/oceanengine/marketing-api/model/keyword"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建关键词 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionID 待创建搜索关键词的广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// ProjectID 待创建搜索关键词的项目ID，仅支持搜索极速智投在项目层级设置关键词
	ProjectID uint64 `json:"project_id,omitempty"`
	// Keywords 关键词列表，一个广告最多可添加1000个
	Keywords []keyword.Keyword `json:"keywords,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
