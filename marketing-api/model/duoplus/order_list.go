package duoplus

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderListRequest 查询订单列表 API Request
type OrderListRequest struct {
	// AwemeSecUid 抖音号ID，通过已授权账户接口可以获取
	AwemeSecUid string `json:"aweme_sec_uid,omitempty"`
	// Filter 筛选条件
	Filter *OrderListFilter `json:"filter,omitempty"`
	// Page 页码，默认值：0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：1- 100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// OrderListFilter 筛选条件
type OrderListFilter struct {
	// SceneType 按营销目标过滤，允许值：:
	// LIVE ：直播
	// VIDEO ：短视频
	SceneType enum.MarketingGoal `json:"scene_type,omitempty"`
	// OrderID 订单id，按id过滤数据
	OrderID []uint64 `json:"order_id,omitempty"`
	// Status DOU+订单状态 可选值：
	// UNPAID未支付
	// AUDITING 审核中
	// OFFLINE_AUDIT审核不通过
	// TIME_NO_REACH 待开播
	// DELIVERING 投放中
	// UNDELIVERIED 投放终止
	// DELIVERIED 投放完成/结束
	// AUDIT_PAUSE 审核暂停
	// WAIT_TO_START 等待开始
	Status enum.DuoplusOrderStatus `json:"status,omitempty"`
}

// Encode implement GetRequest interface
func (r OrderListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("aweme_sec_uid", r.AwemeSecUid)
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
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

// OrderListResponse 查询订单列表 API Response
type OrderListResponse struct {
	model.BaseResponse
	Data *OrderListResult `json:"data,omitempty"`
}

type OrderListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// OrderList 订单列表
	OrderList []Order `json:"order_list,omitempty"`
}
