package star

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReportOrderUserDistributionGetRequest 获取订单投后受众报表 API Request
type ReportOrderUserDistributionGetRequest struct {
	// StarID 星图id，星图客户授权后，通过“获取已授权账户”接口，查询到账号角色为”6-星图账号“的账户id，即为星图id
	StarID uint64 `json:"star_id,omitempty"`
	// OrderID 订单id，可通过“获取星图客户任务订单列表”获取
	OrderID uint64 `json:"order_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ReportOrderUserDistributionGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_id", strconv.FormatUint(r.StarID, 10))
	values.Set("order_id", strconv.FormatUint(r.OrderID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ReportOrderUserDistributionGetResponse 获取订单投后受众报表 API Response
type ReportOrderUserDistributionGetResponse struct {
	// Data json返回值
	Data *ReportOrderUserDistributionGetResult `json:"data,omitempty"`
	model.BaseResponse
}

// ReportOrderUserDistributionGetResult 返回值
type ReportOrderUserDistributionGetResult struct {
	// UpdateTime 数据更新时间，格式%Y-%m-%d %H:%M:%S
	UpdateTime string `json:"update_time,omitempty"`
	// Activity 活跃度分布
	Activity []KV `json:"activity,omitempty"`
	// Age 年龄分布
	Age []KV `json:"age,omitempty"`
	// City 城市分布
	City []KV `json:"city,omitempty"`
	// Device 设备分布
	Device []KV `json:"device,omitempty"`
	// Gender 性别分布
	Gender []KV `json:"gender,omitempty"`
	// Interest 兴趣分布
	Interest []KV `json:"interest,omitempty"`
	// Province 省份分布
	Province []KV `json:"province,omitempty"`
}

// KV .
type KV struct {
	// Key .
	Key string `json:"key,omitempty"`
	// Value 分布值
	Value int64 `json:"value,omitempty"`
}
