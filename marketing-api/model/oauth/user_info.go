package oauth

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

// UserInfoResponse 获取授权User信息 API Response
type UserInfoResponse struct {
	// Data json返回值
	Data *UserInfoResponseData `json:"data,omitempty"`
	model.BaseResponse
}

// UserInfoResponseData userinfo
type UserInfoResponseData struct {
	// Email 邮箱（已经脱敏处理）
	Email string `json:"email,omitempty"`
	// DisplayName 用户名
	DisplayName string `json:"display_name,omitempty"`
	// TokenScopeList 权限点list
	TokenScopeList []uint64 `json:"token_scope_list,omitempty"`
	// TokenApiList 当前token可操作的api接口列表
	TokenApiList []string `json:"token_api_list,omitempty"`
	// ID 用户id
	ID uint64 `json:"id,omitempty"`
	// AppID 授权的应用id
	AppID uint64 `json:"app_id,omitempty"`
	// MaterialAuditStatus 是否敏感物料授权, true 已敏感物料授权, false 未敏感物料授权
	MaterialAuditStatus bool `json:"material_audit_status,omitempty"`
}
