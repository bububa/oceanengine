package finance

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailGetRequest 获取财务流水信息 API Request
type DetailGetRequest struct {
	// AdvertiserID 千川广告主/代理商账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 开始时间，格式 2021-04-05，默认值为end_date-14
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 2021-04-05，默认值为当天
	EndDate string `json:"end_date,omitempty"`
	// Type 资金池类型，允许值：
	// ALL 全部（默认值）
	// DEFAULT 通用
	// BRAND 品牌
	Type qianchuan.ViewDeliveryType `json:"type,omitempty"`
	// Page 页码，默认值: 1
	// 注意：page*page_size不可大于10000
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量，默认值: 10 ，上限：200
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.StartDate != "" {
		values.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		values.Set("end_date", r.EndDate)
	}
	if r.Type != "" {
		values.Set("type", string(r.Type))
	}
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

// DetailGetResponse 获取财务流水信息 API Response
type DetailGetResponse struct {
	model.BaseResponse
	// Date 返回数据，单位为千分之一分，即174938000.00=1749.38000元
	Data *DetailGetResult `json:"data,omitempty"`
}

type DetailGetResult struct {
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	List     []Transaction   `json:"list,omitempty"`
}
