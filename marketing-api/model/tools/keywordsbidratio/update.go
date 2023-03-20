package keywordsbidratio

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 设置优词提量系数和生效维度
type UpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionWordID 优词计划ID，从【查询优词提量系数信息】接口获取
	PromotionWordID string `json:"promotion_word_id,omitempty"`
	// BidRatio 优词提量系数，取值范围:[1.1,2]，仅支持一位小数
	BidRatio float64 `json:"bid_ratio,omitempty"`
	// ProjectIDs 生效的项目id列表，最多100个，当dimension=PROJECT时必填
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
