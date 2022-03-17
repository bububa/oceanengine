package sitetemplate

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 基于站点创建模板 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteID 站点ID，可通过【橙子建站】平台或【获取橙子建站站点列表】接口获取
	SiteID uint64 `json:"site_id,omitempty"`
	// TemplateName 模板名称，默认模版名称为原站点名词
	TemplateName string `json:"template_name,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 基于站点创建模板 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *Template `json:"data,omitempty"`
}
