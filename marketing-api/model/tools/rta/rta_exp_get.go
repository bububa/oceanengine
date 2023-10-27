package rta

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RtaExpGetRequest 获取穿山甲渠道RTA联合实验数据 API Request
type RtaExpGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// RtaID RTA ID
	RtaID int `json:"rta_id"`
	// StartDate 开始日期，格式YYYYMMDD，示例：20220828；注：历史数据最早可回溯到8月1号0点；最大查询跨度为10天；只支持查询当日前90天内的数据
	StartDate string `json:"start_date"`
	// EndDate 结束日期，格式YYYYMMDD，示例：20220902；注：历史数据最早可回溯到8月1号0点；最大查询跨度为10天；只支持查询当日前90天内的数据
	EndDate string `json:"end_date"`
	// Strategy 联合实验策略。共10个实验分桶标记，每个分桶代表一种策略。允许值：0 代表基线策略 ，传入1、2、3、4、5、6、7、8、9
	Strategy int `json:"strategy,omitempty"`
}

// Encode implement GetRequest interface
func (r RtaExpGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("rta_id", strconv.Itoa(r.RtaID))
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	values.Set("strategy", strconv.Itoa(r.Strategy))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RtaExpGetResponse 获取穿山甲渠道RTA联合实验数据 API Response
type RtaExpGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		Data []Report `json:"data_array,omitempty"`
	} `json:"data,omitempty"`
}
