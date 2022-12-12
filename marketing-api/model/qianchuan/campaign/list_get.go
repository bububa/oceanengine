package campaign

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListGetRequest 获取广告组 API Request
type ListGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件，若此字段不传，或传空则视为无限制条件
	Filtering *ListGetFiltering `json:"filtering,omitempty"`
	// Page 当前页码: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小 默认值: 10， 取值范围：1-1000
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ListGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filter", string(filtering))
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

// ListGetFiltering 过滤条件
type ListGetFiltering struct {
	// IDs 广告组ID列表，目前只支持一个
	IDs []uint64 `json:"ids,omitempty"`
	// Name 广告组名称关键字，长度为1-30个字符，其中1个中文字符算2位
	Name string `json:"name,omitempty"`
	// MarketingGoal 广告组营销目标，允许值：VIDEO_PROM_GOODS：短视频带货、LIVE_PROM_GOODS：直播带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// Status 广告组状态，允许值：ALL：所有包含已删除、ENABLE：启用、DISABLE：暂停、DELETE：已删除。不传入即默认返回“所有不包含已删除”
	Status string `json:"status,omitempty"`
}

// ListGetResponse 广告组列表获取 API Response
type ListGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *ListGetResponseData `json:"data,omitempty"`
}

// ListGetResponseData json返回值
type ListGetResponseData struct {
	// PageInfo 分页信息
	PageInfo model.PageInfo `json:"page_info,omitempty"`
	// List 广告组列表
	List []Campaign `json:"list,omitempty"`
}
