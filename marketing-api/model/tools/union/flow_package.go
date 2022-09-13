package union

import "github.com/bububa/oceanengine/marketing-api/enum"

// FlowPackage 流量包
type FlowPackage struct {
	// FlowPackageID 流量包ID
	FlowPackageID uint64 `json:"flow_package_id,omitempty"`
	// Name 流量包名称
	Name string `json:"name,omitempty"`
	// Status 流量包状态
	// FLOW_PACKAGE_ENABLE：已启用
	// FLOW_PACKAGE_DISABLE：已废弃
	Status enum.FlowPackageStatus `json:"status,omitempty"`
	// FlowPackageType 按照流量包类型过滤
	// 枚举值：CUSTOMIZE：自定义、FEATURED：运营推荐、 SYSTEM： 系统推荐
	FlowPackageType enum.FlowPackageType `json:"flow_package_type,omitempty"`
	// Rit 流量位ID数组
	Rit []uint64 `json:"rit,omitempty"`
}
