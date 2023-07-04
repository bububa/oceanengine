package appmanagement

import "github.com/bububa/oceanengine/marketing-api/enum"

// AndroidBasicPackage 安卓应用母包
type AndroidBasicPackage struct {
	// AppName 应用名
	AppName string `json:"app_name,omitempty"`
	// VersionName 版本名
	VersionName string `json:"version_name,omitempty"`
	// VersionCode 版本编码
	VersionCode string `json:"version_code,omitempty"`
	// Status 状态，枚举值：
	// ALL 不限、AUDIT_ACCEPTED 审核通过、AUDIT_DOING 审核中、AUDIT_REJECTED 审核拒绝、ENABLE 可用
	Status enum.AndroidBasicPackageStatus `json:"status,omitempty"`
	// AuditID 审核id，审核通过时id显示为0
	AuditID uint64 `json:"audit_id,omitempty"`
	// AuditMessage 审核失败信息，审核通过显示空
	AuditMessage string `json:"audit_message,omitempty"`
}
