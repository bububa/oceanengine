package site

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FormsListRequest 获取落地页预约表单信息 API Request
type FormsListRequest struct {
	// AdvertiserID 广告主id，数字范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
}

// Encode implement GetRequest interface
func (r FormsListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("site_id", strconv.FormatUint(r.SiteID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// FormsListResponse 获取落地页预约表单信息 API Response
type FormsListResponse struct {
	model.BaseResponse
	Data struct {
		List []Form `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// Form 落地页表单
type Form struct {
	// FormID 落地页表单id
	FormID uint64 `json:"form_id,omitempty"`
	// FormIndex 落地页表单位置
	FormIndex int `json:"form_index,omitempty"`
	// FormName 落地页表单名字
	FormName string `json:"form_name,omitempty"`
}
