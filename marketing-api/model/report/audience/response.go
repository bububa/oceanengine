package audience

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Response 受众分析数据报表 API Response
type Response struct {
	model.BaseResponse
	// Data json返回值
	Data []ResponseData `json:"data,omitempty"`
}

// ResponseData 指标数据
type ResponseData struct {
	// MetricsDict 查询指标列表
	MetricsDict *MetricsDict `json:"metrics_dict,omitempty"`
	// ProvinceName 省份
	ProvinceName string `json:"province_name,omitempty"`
	// CityName 城市
	CityName string `json:"city_name,omitempty"`
	// GenderName 性别，允许值：男，女，其他
	GenderName string `json:"gender_name,omitempty"`
	// AdTagName 兴趣标签
	AdTagName string `json:"ad_tag_name,omitempty"`
	// AgeName 年龄段
	AgeName string `json:"age_name,omitempty"`
}
