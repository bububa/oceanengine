package nativeanchor

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeleteRequest 删除原生锚点 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AnchorID 原生锚点ID
	AnchorID string `json:"anchor_id,omitempty"`
	// AnchorType 原生锚点类型
	AnchorType enum.AnchorType `json:"anchor_type,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
