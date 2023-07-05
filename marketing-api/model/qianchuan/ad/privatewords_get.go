package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PrivatewordsGetRequest 获取否定词列表 API Request
type PrivatewordsGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 筛选条件
	Filtering *PrivatewordsGetFilter `json:"filtering,omitempty"`
}

type PrivatewordsGetFilter struct {
	// AdID 待获取否定词的计划ID
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement GetRequest interface
func (r PrivatewordsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// PrivatewordsGetResponse 获取否定词列表 API Response
type PrivatewordsGetResponse struct {
	model.BaseResponse
	Data *PrivateWords `json:"data,omitempty"`
}
