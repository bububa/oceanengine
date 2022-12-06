package clue

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SmartPhoneGetRequest 建站工具——查询已有智能电话 API Request
type SmartPhoneGetRequest struct {
	// AdvertiserID 广告主id，范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码，默认为1，范围：page >= 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认10条每页，范围：page_size >= 1
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r SmartPhoneGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// SmartPhoneGetResponse 建站工具——查询已有智能电话 API Response
type SmartPhoneGetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *SmartPhoneGetResponseData `json:"data,omitempty"`
}

// SmartPhoneGetResponseData json 返回值
type SmartPhoneGetResponseData struct {
	// AdvKey 表征创建者信息的字段，历史为加密的advId
	AdvKey string `json:"adv_key,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 智能电话列表
	List []SmartPhone `json:"list,omitempty"`
}
