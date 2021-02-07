package audience

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

type Response struct {
	model.BaseResponse
	Data *ResponseData `json:"data,omitempty"`
}

type ResponseData struct {
	List     []ResponseList  `json:"list,omitempty"`
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type ResponseList struct {
	MetricsDict  *MetricsDict `json:"metrics_dict,omitempty"`  // 查询指标列表
	ProvinceName string       `json:"province_name,omitempty"` // 省份
	CityName     string       `json:"city_name,omitempty"`     // 城市
	GenderName   string       `json:"gender_name,omitempty"`   // 性别，允许值：男，女，其他
	AdTagName    string       `json:"ad_tag_name,omitempty"`   // 兴趣标签
	AgeName      string       `json:"age_name,omitempty"`      // 年龄段
}
