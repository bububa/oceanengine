package interestaction

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// ActionCategoryRequest 行为类目查询 API Request
type ActionCategoryRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ActionScene 行为场景，查询目标为行为时必填，兴趣不生效;允许值：E-COMMERCE、NEWS、APP
	ActionScene enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 行为天数，查询目标为行为时必填，兴趣不生效; 允许值：7, 15, 30, 60, 90, 180, 365
	ActionDays int `json:"action_days,omitempty"`
}

// Encode implement GetRequest interface
func (r ActionCategoryRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("action_scene", string(r.ActionScene))
	values.Set("action_days", strconv.Itoa(r.ActionDays))
	return values.Encode()
}

// ActionCategoryResponse 行为类目查询 API Response
type ActionCategoryResponse struct {
	model.BaseResponse
	// Data json返回值
	Data []Object `json:"data,omitempty"`
}
