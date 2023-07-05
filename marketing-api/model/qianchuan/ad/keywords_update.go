package ad

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// KeywordsUpdateRequest 更新关键词 API Request
type KeywordsUpdateRequest struct {
	// AdvertiserID 千川广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// Keywords 更新关键词
	Keywords []UpdateKeyword `json:"keywords,omitempty"`
}

type UpdateKeyword struct {
	// ID 关键词ID
	ID uint64 `json:"id,omitempty"`
	// StatusType 关键词操作状态 可选值:
	// DELETE: 删除
	// ENABLE: 启动
	// PAUSED: 暂停
	StatusType qianchuan.KeywordStatus `json:"status_type,omitempty"`
	// MatchType 关键词匹配类型 可选值:
	// EXTENSIVE: 广泛匹配
	// PHRASE: 短语匹配
	// PRECISION: 精准匹配
	MatchType enum.KeywordMatchType `json:"match_type,omitempty"`
}

// Encode implement PostRequest interface
func (r KeywordsUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
