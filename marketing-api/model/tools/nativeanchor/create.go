package nativeanchor

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 原生锚点创建 API Request
type CreateRequest struct {
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AnchorInfo 锚点内容
	AnchorInfo *NativeAnchor `json:"anchor_info,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 原生锚点创建API Response
type CreateResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *CreateResponseData `json:"data,omitempty"`
}

// CreateResponseData 返回数据
type CreateResponseData struct {
	// AnchorID
	AnchorID string `json:"anchor_id,omitempty"`
	// AnchorType 可选值:
	// APP_GAME: 应用下载-游戏锚点
	// APP_INTERNET_SERVICE: 应用下载-网服锚点
	// APP_SHOP: 应用下载-电商锚点
	// ONLINE_SUBSCRIBE: 通用锚点
	AnchorType enum.AnchorType `json:"anchor_type,omitempty"`
}
