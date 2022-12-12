package servemarket

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderGetRequest 获取应用订单数据 API Request
type OrderGetRequest struct {
	// AppID 应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// Filtering 过滤条件
	Filtering *OrderGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 分页数量，默认20，最大支持100
	PageSize int `json:"page_size,omitempty"`
}

// OrderGetFilter 过滤条件
type OrderGetFilter struct {
	// OrderID 按订单ID过滤，数组，支持N个入参，仅拉取应用相关的订单
	OrderID []uint64 `json:"order_id,omitempty"`
	// UseUid 按使用用户ID过滤，仅支持一个入参，仅拉取应用相关的订单。
	// 通过【获取授权时登录用户信息】接口可以获取授权用户user_id，该user_id为巨量引擎的登录帐号ID（与手机号/邮箱一一映射）
	UseUid uint64 `json:"use_uid,omitempty"`
}

// Encode implement GetReqeust
func (r OrderGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("app_id", strconv.FormatUint(r.AppID, 10))
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OrderGetResponse 获取应用订单数据 API Response
type OrderGetResponse struct {
	model.BaseResponse
	Data *OrderGetResponseData `json:"data,omitempty"`
}

type OrderGetResponseData struct {
	// PageInfo 页面信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// OrderList 订单列表
	OrderList []Order `json:"order_list,omitempty"`
}
