package keywordsbidratio

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 设置优词提量系数和生效维度
type CreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// WordList 优词列表，最多支持10个，字数最大限制：30，每个优词生成一个优词计划id
	// 每个账户最多支持200个优词
	WordList []string `json:"word_list,omitempty"`
	// BidRatio 优词提量系数，取值范围:[1.1,2]，仅支持一位小数
	BidRatio float64 `json:"bid_ratio,omitempty"`
	// Dimension 生效维度，可选值:
	// ADVERTISER: 账户维度
	// PROJECT: 项目维度
	Dimension string `json:"dimension,omitempty"`
	// ProjectIDs 生效的项目id列表，最多100个，当dimension=PROJECT时必填
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
