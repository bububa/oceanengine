package oauth

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

type UserInfoResponse struct {
	model.BaseResponse
	Data *UserInfoResponseData `json:"data,omitempty"`
}

type UserInfoResponseData struct {
	ID          uint64 `json:"id,omitempty"`           // 用户id
	Email       string `json:"email,omitempty"`        // 邮箱（已经脱敏处理）
	DisplayName string `json:"display_name,omitempty"` // 用户名
}
