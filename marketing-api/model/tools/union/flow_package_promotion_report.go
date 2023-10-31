package union

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FlowPackagePromotionReportRequest 查看穿山甲2.0广告位数据 API Request
type FlowPackagePromotionReportRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filter 过滤字段
	Filter *FlowPackagePromotionReportFilter `json:"filter,omitempty"`
	// OrderField 排序字段，所有的统计指标均可参与排序
	// 默认按rit排序
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式
	// 枚举值：ASC（升序），DESC（降序）
	// 默认：DESC
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页数
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	// 默认值: 10，最大值100
	PageSize int `json:"page_size,omitempty"`
}

type FlowPackagePromotionReportFilter struct {
	// PromotionIDs 广告 id 列表，最大数量限制：1000
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
	// Rits 广告位列表，最大数量限制：1000
	Rits []uint64 `json:"rits,omitempty"`
	// LandingType 推广目的，允许值：APP 应用推广、LINK 销售线索推广、MICRO_GAME 小程序、SHOP 电商店铺推广、QUICK_APP快应用、DPA商品目录 可选值:
	// APP
	// DPA
	// LINK
	// MICRO_GAME
	// QUICK_APP
	// SHOP
	LandngType enum.LandingType `json:"landng_type,omitempty"`
	// HighCost 消耗金额上限，单位元
	HighCost model.JSONInt64 `json:"high_cost,omitempty"`
	// LowCost 消耗金额下限，单位元
	LowCost model.JSONInt64 `json:"low_cost,omitempty"`
	// StartTime 开始时间，格式为"yyyy-mm-dd"
	// 限制范围100天内
	// 默认7天前（不包括当天），即不指定起止时间获取最近7天数据
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间，格式为"yyyy-mm-dd"
	// 默认昨天，即不指定起止时间获取最近7天数据
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement GetRequest interface
func (r FlowPackagePromotionReportRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 10 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
