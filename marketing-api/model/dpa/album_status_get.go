package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AlbumStatusGetRequest 查询短剧可投状态 API Request
type AlbumStatusGetRequest struct {
	// AlbumID 短剧id
	AlbumID string `json:"album_id,omitempty"`
	// AppID 请传入app_access_token对应的app_id，可通过开放平台官网-开发者后台查询
	AppID uint64 `json:"app_id,omitempty"`
}

// Encode implements GetRequest interface
func (r AlbumStatusGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("album_id", r.AlbumID)
	values.Set("app_id", strconv.FormatUint(r.AppID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AlbumStatusGetResponse 查询短剧可投状态 API Response
type AlbumStatusGetResponse struct {
	model.BaseResponse
	Data *AlbumStatusGetResult `json:"data,omitempty"`
}

type AlbumStatusGetResult struct {
	// CanPromotion 是否可投
	CanPromption string `json:"can_promotion,omitempty"`
	// CanNotPromptionReason 不可投原因 可选值:
	// RESOURCE_NOT_EXIST 资源尚未创建或非微小短剧
	// RESOURCE_AUDITING 资源审核中
	// RESOURCE_UNDERRATING 资源可见度未达到可推广
	// RESOURCE_CREATING 资源创建中
	// RESOURCE_FAILED 资源创建失败
	CanNotPromotionReason enum.AlbumCanNotPromotionReason `json:"can_not_promotion_reason,omitempty"`
}
