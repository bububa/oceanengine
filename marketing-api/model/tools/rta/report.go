package rta

import "github.com/bububa/oceanengine/marketing-api/model"

type Report struct {
	// Date 数据统计日期，格式YYYYMMDDHH
	Date string `json:"date,omitempty"`
	// VID 联合实验组唯一标识
	VID int `json:"vid,omitempty"`
	// CusVID 客户自行开分桶实验的唯一标识
	CusVID int `json:"cus_vid,omitempty"`
	// Strategy 联合实验策略，请求入参
	Strategy model.Int `json:"strategy,omitempty"`
	// Click 展现数据-点击数。当头条用户点击广告素材时，触发点击事件，该事件被认为是一次有效的广告点击
	Click int64 `json:"click,omitempty"`
	// Show 展现数据-展示数。广告展示给用户的次数。计算方式：经平台判定有效且被计费的展示次数
	Show int64 `json:"show,omitempty"`
	// Convert 转化数据-转化数。将转化数记录在转化事件发生的时间上。建议广告主考核成本时参考“转化数据（计费时间）”例如您的广告在早上8点进行了展示和点击，用户晚上19点发生了激活行为，我们会把激活数披露在晚上19点
	Convert int64 `json:"convert,omitempty"`
	// Cost 展现数据-总花费。表示广告在投放期内的预估花费金额
	Cost float64 `json:"cost,omitempty"`
	// BidCoef 返回RTA出价系数的区间值
	BidCoef string `json:"bid_coef,omitempty"`
	// WinRatio 竞胜率。竞胜率=竞胜数/参竞数，代表广告主参竞请求的胜出比例，范围 0~1
	WinRatio float64 `json:"win_ratio,omitempty"`
}
