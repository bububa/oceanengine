package privativeword

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取否定词列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 搜索否定词查询条件
	Filtering *GetFiltering `json:"filtering,omitempty"`
}

// GetFiltering 搜索否定词查询条件
type GetFiltering struct {
	// AdIDs 广告计划ids, 最多100个, ad_ids 与 campaign_ids 同时传入，取并集操作
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// CampaignIDs 广告组ids, 最多100个, ad_ids 与 campaign_ids 同时传入，取并集操作
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 获取否定词列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// AdsPrivative 广告计划否定词列表
	AdsPrivative []AdWord `json:"ads_privative,omitempty"`
	// CampaignsPrivative 广告组否定词列表
	CampaignsPrivative []AdWord `json:"campaigns_privative,omitempty"`
	// AdErrorList 请求失败的广告计划id
	AdErrorList []string `json:"ad_error_list,omitempty"`
	// CampaignErrorList 请求失败的广告组id
	CampaignErrorList []string `json:"campaign_error_list,omitempty"`
}
