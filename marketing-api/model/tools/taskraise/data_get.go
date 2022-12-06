package taskraise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DataGetRequest 查询优选起量任务数据 API Request
type DataGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ReportID 任务id
	ReportID uint64 `json:"report_id,omitempty"`
}

// Encode implement GetRequet interface
func (r DataGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("report_id", strconv.FormatUint(r.ReportID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DataGetResponse 查询优选起量任务数据 API Response
type DataGetResponse struct {
	model.BaseResponse
	Data *DataGetResponseData `json:"data,omitempty"`
}

type DataGetResponseData struct {
	// Cost 消耗数据
	Cost Stat `json:"cost,omitempty"`
	// Show 展示数据
	ShowCnt Stat `json:"show_cnt,omitempty"`
	// Click 点击数据
	Click Stat `json:"click_cnt,omitempty"`
	// Convert 转化数据
	Convert Stat `json:"convert_cnt,omitempty"`
}

type Stat struct {
	// Base 基础数据
	Base string `json:"base_value,omitempty"`
	// Increased 起量任务数据
	Increased string `json:"increased_value,omitempty"`
}
