package coupon

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CodeGetRequest 查询券码记录 API Request
type CodeGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ActivityID 活动ID，可从获取卡券详情接口获取
	ActivityID uint64 `json:"activity_id,omitempty"`
	// CouponID 卡券活动ID，可从获取卡券详情接口获取
	CouponID uint64 `json:"coupon_id,omitempty"`
	// ResourceID 优惠券ID，可从获取卡券详情接口获取
	ResourceID uint64 `json:"resource_id,omitempty"`
	// StartTime 起始时间，格式：%Y-%m-%d，默认7天前，即不指定起止时间获取最近7天数据
	StartTime string `json:"start_time,omitempty"`
	// EndTime 截止时间，格式：%Y-%m-%d，默认当天，即不指定起止时间获取最近7天数据
	EndTime string `json:"end_time,omitempty"`
	// Status 券码状态
	// VALID:有效中
	// EXPIRED:已过期
	// USED:已使用
	// ABANDONED:已废弃
	Status CodeStatus `json:"status,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值10，不超过50
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r CodeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("activity_id", strconv.FormatUint(r.ActivityID, 10))
	if r.CouponID > 0 {
		values.Set("coupon_id", strconv.FormatUint(r.CouponID, 10))
	}
	if r.ResourceID > 0 {
		values.Set("resource_id", strconv.FormatUint(r.ResourceID, 10))
	}
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.Status != "" {
		values.Set("status", string(r.Status))
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

// CodeGetResponse 查询券码记录 API Response
type CodeGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CodeGetResponseData `json:"data,omitempty"`
}

// CodeGetResponseData json返回值
type CodeGetResponseData struct {
	// PageInfo 分页数据
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 券码记录列表
	List []Code `json:"list,omitempty"`
}
