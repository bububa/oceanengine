package landinggroup

import "github.com/bububa/oceanengine/marketing-api/enum"

// LandingGroup 落地页组
type LandingGroup struct {
	// GroupID 落地页组 ID
	GroupID uint64 `json:"group_id,omitempty"`
	// GroupTitle 落地页组名称
	GroupTitle string `json:"group_title,omitempty"`
	// GroupStatus 落地页组状态
	GroupStatus enum.LandingGroupStatus `json:"group_status,omitempty"`
	// GroupFlowType 流量分配方式
	GroupFlowType enum.GroupFlowType `json:"group_flow_type,omitempty"`
	// Sites 站点列表
	Sites []Site `json:"sites,omitempty"`
}
