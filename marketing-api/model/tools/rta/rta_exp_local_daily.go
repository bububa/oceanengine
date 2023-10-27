package rta

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RtaExpLocalDailyGetRequest 获取站内媒体RTA联合实验数据 API Request
type RtaExpLocalDailyGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// RtaID RTA ID
	RtaID uint64 `json:"rta_id"`
	// StartDate 开始日期，格式YYYYMMDD，示例：20220828；注：历史数据最早可回溯到8月1号0点；最大查询跨度为10天；只支持查询当日前90天内的数据
	StartDate string `json:"start_date"`
	// EndDate 结束日期，格式YYYYMMDD，示例：20220902；注：历史数据最早可回溯到8月1号0点；最大查询跨度为10天；只支持查询当日前90天内的数据
	EndDate string `json:"end_date"`
	// VID 联合实验组唯一标识，共10个实验分桶标记。若使用该数据报表，务必在入参中携带cus_vid或者vid二者之一，若为空则报错。
	// 注：联系对接销售/运营咨询获取vid
	VID int `json:"vid,omitempty"`
	// CusVID 客户自行开分桶实验的唯一标识，一共10个枚举值（1~10）。若使用该数据报表，务必在入参中携带cus_vid或者vid二者之一，若为空则报错。
	// 注：cus_vid和vid不能同时使用，同时入参；如有特殊需求，联系对接销售/运营咨询。
	CusVID int `json:"cus_vid,omitempty"`
	// FIltering 过滤条件
	Filtering *RtaExpLocalDailyGetFilter `json:"filtering,omitempty"`
}

type RtaExpLocalDailyGetFilter struct {
	// BidCoef 出价系数取数区间。枚举值：COEF_0（表示系数为0）、COEF_BETWEEN_0_0.1（表示系数区间为(0,0.1)）、COEF_BETWEEN_0.1_0.2（表示系数区间为[0.1,0.2)）、COEF_BETWEEN_0.2_0.3（表示系数区间为[0.2,0.3)）、COEF_BETWEEN_0.3_0.4（表示系数区间为[0.3,0.4)）、COEF_BETWEEN_0.4_0.5（表示系数区间为[0.4,0.5)）、COEF_BETWEEN_0.5_0.6（表示系数区间为[0.5,0.6)）、COEF_BETWEEN_0.6_0.7（表示系数区间为[0.6,0.7)）、COEF_BETWEEN_0.7_0.8（表示系数区间为[0.7,0.8)）、COEF_BETWEEN_0.8_0.9（表示系数区间为[0.8,0.9)）、COEF_BETWEEN_0.9_1.0（表示系数区间为[0.9,1.0)）、COEF_BETWEEN_1.0_1.1（表示系数区间为[1.0,1.1)）、COEF_BETWEEN_1.1_1.2（表示系数区间为[1.1,1.2)）、COEF_BETWEEN_1.2_1.3（表示系数区间为[1.2,1.3)）、COEF_BETWEEN_1.3_1.4（表示系数区间为[1.3,1.4)）、COEF_BETWEEN_1.4_1.5（表示系数区间为[1.4,1.5)）、COEF_BETWEEN_1.5_1.6（表示系数区间为[1.5,1.6)）、COEF_BETWEEN_1.6_1.7（表示系数区间为[1.6,1.7)）、COEF_BETWEEN_1.7_1.8（表示系数区间为[1.7,1.8)）、COEF_BETWEEN_1.8_1.9（表示系数区间为[1.8,1.9)）、COEF_BETWEEN_1.9_2.0（表示系数区间为[1.9,2.0)）、COEF_BETWEEN_2.0_2.1（表示系数区间为[2.0,2.1)）、COEF_BETWEEN_2.1_2.2（表示系数区间为[2.1,2.2)）、COEF_BETWEEN_2.2_2.3（表示系数区间为[2.2,2.3)）、COEF_BETWEEN_2.3_2.4（表示系数区间为[2.3,2.4)）、COEF_BETWEEN_2.4_2.5（表示系数区间为[2.4,2.5)）、COEF_BETWEEN_2.5_2.6（表示系数区间为[2.5,2.6)）、COEF_BETWEEN_2.6_2.7（表示系数区间为[2.6,2.7)）、COEF_BETWEEN_2.7_2.8（表示系数区间为[2.7,2.8)）、COEF_BETWEEN_2.8_2.9（表示系数区间为[2.8,2.9)）、COEF_BETWEEN_2.9_3.0（表示系数区间为[2.9,3.0)）、COEF_BETWEEN_3.0_3.5（表示系数区间为[3.0,3.5)）、COEF_BETWEEN_3.5_4.0（表示系数区间为[3.5,4.0)）、COEF_BETWEEN_4.0_4.5（表示系数区间为[4.0,4.5)）、COEF_BETWEEN_4.5_5.0（表示系数区间为[4.5,5.0)）、COEF_ABOVE_5.0（表示系数大于等于5）
	BidCoef []string `json:"bid_coef"`
}

func (r RtaExpLocalDailyGetRequest) Encode() string {
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
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RtaExpLocalDailyGetResponse 获取站内媒体RTA联合实验数据 API Response
type RtaExpLocalDailyGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		Data []Report `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
