package comment

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TermsBannedGetRequest 获取屏蔽词列表 API Request
type TermsBannedGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 为绑定的抖音号添加屏蔽词, 只允许传入1个，通过【获取绑定的抖音号】 接口获取，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeID string `json:"aweme_id,omitempty"`
	// IsApplyToAdv 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv bool `json:"is_apply_to_adv,omitempty"`
	// Page 页数 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r TermsBannedGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.AwemeID != "" {
		values.Set("aweme_id", r.AwemeID)
	}
	if r.IsApplyToAdv {
		values.Set("is_apply_to_adv", "true")
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

// TermsBannedGetResponse 获取屏蔽词列表 API Response
type TermsBannedGetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *TermsBannedGetResponseData `json:"data,omitempty"`
}

// TermsBannedGetResponseData json 返回值
type TermsBannedGetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Terms 屏蔽词列表
	Terms []string `json:"terms,omitempty"`
}
