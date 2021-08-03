package oauth

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

// UserInfoResponse 获取授权User信息 API Response
type UserInfoResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UserInfoResponseData `json:"data,omitempty"`
}

// UserInfoResponseData userinfo
type UserInfoResponseData struct {
	// ID 用户id
	ID uint64 `json:"id,omitempty"`
	// Email 邮箱（已经脱敏处理）
	Email string `json:"email,omitempty"`
	// DisplayName 用户名
	DisplayName string `json:"display_name,omitempty"`
}
