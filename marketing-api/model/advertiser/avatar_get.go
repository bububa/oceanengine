package advertiser

import (
	"fmt"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// AvatarGetRequest 获取广告主头像信息 API Request
type AvatarGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AvatarGetRequest) Encode() string {
	return fmt.Sprintf("advertiser_id=%d", r.AdvertiserID)
}

// AvatarGetResponse 获取广告主头像信息 API Response
type AvatarGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AvatarGetResponseData `json:"data,omitempty"`
}

// AvatarGetResponseData 获取广告主头像信息json返回值
type AvatarGetResponseData struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AvatarStatus 头像审核状态 0-未设置，1-审核中，2-审核被拒，3-审核通过
	AvatarStatus int `json:"avatar_status,omitempty"`
	// AvatarReason 头像被拒原因
	AvatarReason string `json:"avatar_reason,omitempty"`
	// SourceStatus 品牌审核状态 0-未设置，1-审核中，2-审核被拒，3-审核通过,
	SourceStatus int `json:"source_status,omitempty"`
	// SourceReason 品牌信息被拒原因
	SourceReason string `json:"source_reason,omitempty"`
	// AvatarInfo 头像信息
	AvatarInfo *AvatarInfo `json:"avatar_info,omitempty"`
}

// AvatarInfo 头像信息
type AvatarInfo struct {
	// WebUri 当前头像的uri
	WebUri string `json:"web_uri,omitempty"`
	// AuditWebUri 审核中头像的uri
	AuditWebUri string `json:"audit_web_uri,omitempty"`
	// Height 审核中头像的高度
	Height int `json:"height,omitempty"`
	// Width 审核中头像的宽度
	Width int `json:"width,omitempty"`
}
