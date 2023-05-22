package v3

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SuggestRequest 搜索快投关键词推荐 API Request
type SuggestRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// QueryList 以词推词，最多只传入10个，长度要求30字内 ,可传入行业名称获取行业推词
	QueryList []string `json:"query_list,omitempty"`
	// PromotionMaterial 广告素材
	PromotionMaterial *PromotionMaterial `json:"promotion_material,omitempty"`
}

// PromotionMaterial 广告素材
type PromotionMaterial struct {
	// Abstracts 广告摘要
	Abstracts []string `json:"abstracts,omitempty"`
	// QuantityWords 优词
	QuantityWords []string `json:"quantity_words,omitempty"`
	// AppName 应用名
	AppName string `json:"app_name,omitempty"`
	// Source 来源
	Source string `json:"source,omitempty"`
}

// Encode implement PostRequest interface
func (r SuggestRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
