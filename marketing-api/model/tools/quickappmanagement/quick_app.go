package quickappmanagement

import "github.com/bububa/oceanengine/marketing-api/enum"

// QuickApp 快应用信息
type QuickApp struct {
	// QuickAppID 快应用id
	QuickAppID uint64 `json:"quick_app_id,omitempty"`
	// Name 快应用名称
	Name string `json:"name,omitempty"`
	// PackageName 快应用包名
	PackageName string `json:"package_name,omitempty"`
	// HomepageURL 快应用主页链接
	HomepageURL string `json:"homepage_url,omitempty"`
	// Status 快应用状态；可选值:
	// AUDIT_ACCEPTED: 审核通过 （已上架、已发布）
	// AUDIT_DOING: 审核中
	// AUDIT_REJECTED: 审核拒绝 （审核失败)
	// AUDIT_SEND_REJECTED: 送审失败
	// REMOVED: 已下架
	Status enum.QuickAppStatus `json:"status,omitempty"`
	// CreateTime 更新时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 更新时间，格式','=Y-%m-%d %H:%M:%S'
	UpdateTime string `json:"update_time,omitempty"`
}
