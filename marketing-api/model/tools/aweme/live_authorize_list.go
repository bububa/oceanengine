package aweme

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// LiveAuthorizeListRequest 查询授权直播抖音达人列表
type LiveAuthorizeListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Status 直播达人授权状态列表 ，允许值：APPLYING申请中、APPLY_OVERDUE申请过期、AUTHORIZING授权中、AUTHORIZE_OVERDUE授权过期
	Status []enum.LiveAuthorizeStatus `json:"status,omitempty"`
	// Page 当前页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值：10，最大值：50
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r LiveAuthorizeListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	status, _ := json.Marshal(r.Status)
	values.Set("status", string(status))
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

// LiveAuthorizeListResponse 查询授权直播抖音达人列表 API Response
type LiveAuthorizeListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *LiveAuthorizeListResponseData `json:"data,omitempty"`
}

// LiveAuthorizeListResponseData json返回值
type LiveAuthorizeListResponseData struct {
	// PageInfo 分页信息
	PageInfo model.PageInfo `json:"page_info,omitempty"`
	// List 达人信息列表
	List []LiveAuthorize `json:"list,omitempty"`
}
