package advertiser

import (
	"fmt"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type AvatarGetRequest struct {
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

func (r AvatarGetRequest) Encode() string {
	return fmt.Sprintf("advertiser_id=%d", r.AdvertiserID)
}

type AvatarGetResponse struct {
	model.BaseResponse
	Data *AvatarGetResponseData `json:"data,omitempty"`
}

type AvatarGetResponseData struct {
	AdvertiserID uint64      `json:"advertiser_id,omitempty"` // 广告主id
	AvatarStatus int         `json:"avatar_status,omitempty"` // 头像审核状态 0-未设置，1-审核中，2-审核被拒，3-审核通过
	AvatarReason string      `json:"avatar_reason,omitempty"` // 头像被拒原因
	SourceStatus int         `json:"source_status,omitempty"` // 品牌审核状态 0-未设置，1-审核中，2-审核被拒，3-审核通过,
	SourceReason string      `json:"source_reason,omitempty"` // 品牌信息被拒原因
	AvatarInfo   *AvatarInfo `json:"avatar_info,omitempty"`   // 头像信息
}

type AvatarInfo struct {
	WebUri      string `json:"web_uri,omitempty"`       // 当前头像的uri
	AuditWebUri string `json:"audit_web_uri,omitempty"` // 审核中头像的uri
	Height      int    `json:"height,omitempty"`        // 审核中头像的高度
	Width       int    `json:"width,omitempty"`         // 审核中头像的宽度
}
