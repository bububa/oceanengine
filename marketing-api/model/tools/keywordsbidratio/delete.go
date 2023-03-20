package keywordsbidratio

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeleteRequest 删除优词提量系数和生效维度
type DeleteRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionWordIDs 要删除的优词计划ID列表，从【查询优词提量系数信息】接口获取,最多支持10个
	PromotionWordIDs []string `json:"promotion_word_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
