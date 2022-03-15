package tools

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// AdminInfoRequest 获取行政信息 API Request
type AdminInfoRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Codes 行政区域编码列表
	Codes []string `json:"codes,omitempty"`
	// Language 语言类型;
	// ZH_CN表示常用名，如“北京”
	// ZH_CN_GOV表示官方全称，如“北京市”
	Language string `json:"language,omitempty"`
	// SubDistrict 行政区域层级。
	// NONE 当前层级
	// ONE_LEVEL下一级区域
	// TWO_LEVEL下二级区域
	// THREE_LEVEL下三级区域
	// FOUR_LEVEL下四级区域
	SubDistrict string `json:"sub_district,omitempty"`
}

// Encode implement GetRequest interface
func (r AdminInfoRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	codes, _ := json.Marshal(r.Codes)
	values.Set("codes", string(codes))
	values.Set("language", r.Language)
	values.Set("sub_district", r.SubDistrict)
	return values.Encode()
}

// AdminInfoResponse 获取行政信息 API Response
type AdminInfoResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdminInfoResponseData `json:"data,omitempty"`
}

// AdminInfoResponseData json返回值
type AdminInfoResponseData struct {
	// Version 行政信息版本号
	Version string `json:"version,omitempty"`
	// Districts 行政层级信息
	Districts []District `json:"districts,omitempty"`
}
