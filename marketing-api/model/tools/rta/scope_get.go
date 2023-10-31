package rta

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ScopeGetRequest 获取RTA策略绑定信息列表 API Request
type ScopeGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RtaID RTA策略ID
	RtaID uint64 `json:"rta_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ScopeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("rta_id", strconv.FormatUint(r.RtaID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ScopeGetResponse 获取RTA策略绑定信息列表 API Response
type ScopeGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 指定rta策略下所绑定的信息列表
		List []Scope `json:"list,omitempty"`
	}
}

// Scope 指定rta策略下所绑定的信息
type Scope struct {
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// ProjectName 项目名称
	ProjectName string `json:"project_name,omitempty"`
	// CampaignID 广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 广告组名称
	CampaignName string `json:"campaign_name,omitempty"`
}
