package sitetemplate

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PreviewRequest 获取模版预览链接 API Request
type PreviewRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TemplateID 模板ID，可通过【获取站点模版列表】查询模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
}

// Encode implement GetRequest interface
func (r PreviewRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("template_id", strconv.FormatUint(r.TemplateID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// PreviewResponse 获取模版预览链接 API Response
type PreviewResponse struct {
	model.BaseResponse
	Data struct {
		// PreviewURL 预览URL，模板的预览地址，有效期：20分钟
		PreviewURL string `json:"preview_url,omitempty"`
	} `json:"data,omitempty"`
}
