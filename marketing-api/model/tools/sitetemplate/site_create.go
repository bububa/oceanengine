package sitetemplate

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SiteCreateRequest 基于模板创建站点 API Request
type SiteCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteID 站点ID，可通过【橙子建站】平台或【获取橙子建站站点列表】接口获取
	SiteID uint64 `json:"site_id,omitempty"`
	// TemplateID 模板ID，可通过【获取站点模版列表】查询模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// Name 站点名称，范围：长度 >= 1
	Name string `json:"name,omitempty"`
	// Bricks 组件列表，可按需填充部分组件，未填充部分会使用原模板内容。
	Bricks []Brick `json:"bricks,omitempty"`
}

// Encode implement PostRequest interface
func (r SiteCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SiteCreateResponse 基于模板创建站点 API Response
type SiteCreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// SiteID 站点ID
		SiteID uint64 `json:"site_id,omitempty"`
	} `json:"data,omitempty"`
}
