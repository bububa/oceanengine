package rta

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RtaExpLocalHourlyGetRequest 获取站内媒体RTA联合实验数据（分时t+5）API Request
type RtaExpLocalHourlyGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// RtaID RTA ID
	RtaID uint64 `json:"rta_id"`
	// StarteDate 开始时间，格式YYYYMMDDHH，示例：2022090208；注：HH取值范围为00-23；历史数据最早可回溯到8月1号0点；最大查询跨度为10天；只支持查询当日前90天内且查询当时5个小时以前的数据；若start_date=2022090208，end_date=2022090209，则获取2022年9月2日8时、9时整两小时的数据
	StartDate string `json:"start_date"`
	// EndDate 结束时间，格式YYYYMMDDHH，示例：2022090209；注：HH取值范围为00-23；历史数据最早可回溯到8月1号0点；最大查询跨度为10天；只支持查询当日前90天内且查询当时5个小时以前的数据；若start_date=2022090208，end_date=2022090209，则获取2022年9月2日8时、9时整两小时的数据
	EndDate string `json:"end_date"`
	// VID 联合实验组唯一标识，共10个实验分桶标记。若使用该数据报表，务必在入参中携带cus_vid或者vid二者之一，若为空则报错。
	// 注：联系对接销售/运营咨询获取vid
	VID int `json:"vid,omitempty"`
	// CusVID 客户自行开分桶实验的唯一标识，一共10个枚举值（1~10）。若使用该数据报表，务必在入参中携带cus_vid或者vid二者之一，若为空则报错。
	// 注：cus_vid和vid不能同时使用，同时入参；如有特殊需求，联系对接销售/运营咨询。
	CusVID int `json:"cus_vid,omitempty"`
}

func (r RtaExpLocalHourlyGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("rta_id", strconv.FormatUint(r.RtaID, 10))
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if r.VID > 0 {
		values.Set("vid", strconv.Itoa(r.VID))
	}
	if r.CusVID > 0 {
		values.Set("cus_vid", strconv.Itoa(r.CusVID))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RtaExpLocalHourlyGetResponse 获取站内媒体RTA联合实验数据（分时t+5） API Response
type RtaExpLocalHourlyGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		Data []Report `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
