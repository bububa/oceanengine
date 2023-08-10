package sitetemplate

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PicURLGetRequest 获取模板/站点URL API Request
type PicURLGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
	// TemplateID 模板id
	TemplateID uint64 `json:"template_id,omitempty"`
}

// Encode implement GetRequest interface
func (r PicURLGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("site_id", strconv.FormatUint(r.SiteID, 10))
	values.Set("template_id", strconv.FormatUint(r.TemplateID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// PicURLGetResponse 获取模板/站点URL API Response
type PicURLGetResponse struct {
	model.BaseResponse
	Data struct {
		// URLMap uri
		URLMap map[string]string `json:"url_map,omitempty"`
	} `json:"data,omitempty"`
}
