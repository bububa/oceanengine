package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CountryInfoRequest 查询国家/区域信息 API Request
type CountryInfoRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Language 语言类型;
	// ZH_CN表示常用名，如“北京”
	// ZH_CN_GOV表示官方全称，如“北京市”
	Language string `json:"language,omitempty"`
}

// Encode implement GetRequest interface
func (r CountryInfoRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("language", r.Language)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CountryInfoResponse 查询国家/区域信息 API Response
type CountryInfoResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CountryInfoResponseData `json:"data,omitempty"`
}

// CountryInfoResponseData json返回值
type CountryInfoResponseData struct {
	// Districts 行政区域信息
	Districts []District `json:"districts,omitempty"`
}
