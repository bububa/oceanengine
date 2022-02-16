package site

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// SiteDetail 站点信息
type SiteDetail struct {
	// Bricks 具体见返回示例业务数据（新建或更新时传递的数据）
	Bricks []Brick `json:"bricks,omitempty"`
	// ID 站点ID
	ID model.FlexUint64 `json:"id,omitempty"`
	// Status 站点状态
	Status enum.SiteStatus `json:"status,omitempty"`
	// SiteType 建站类型
	SiteType enum.SiteType `json:"site_type,omitempty"`
	// Thumbnail 缩略图
	Thumbnail string `json:"thumbnail,omitempty"`
}

// Site 站点信息
type Site struct {
	// SiteID 站点ID
	SiteID model.FlexUint64 `json:"siteId,omitempty"`
	// Name 建站名称
	Name string `json:"name,omitempty"`
	// Status 站点状态
	Status enum.SiteStatus `json:"status,omitempty"`
	// SiteType 建站类型
	SiteType enum.SiteType `json:"siteType,omitempty"`
	// FunctionType 建站类别，SITE_FUNC_TYPE_NEW_MODULAR（智能建站）SITE_FUNC_TYPE_NORMAL （普通建站）
	FunctionType enum.SiteFunctionType `json:"function_type,omitempty"`
	// Thumbnail 缩略图
	Thumbnail string `json:"thumbnail,omitempty"`
}
