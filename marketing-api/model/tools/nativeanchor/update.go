package nativeanchor

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 原生锚点更新 API Request
type UpdateRequest struct {
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AnchorInfo 锚点内容
	AnchorInfo *NativeAnchor `json:"anchor_info,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 原生锚点更新API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData 返回数据
type UpdateResponseData struct {
	// AnchorID
	AnchorID string `json:"anchor_id,omitempty"`
}
