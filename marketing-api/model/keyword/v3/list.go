package v3

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取关键词列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤结构
	Filtering *ListFiltering `json:"filtering,omitempty"`
}

// ListFiltering 过滤结构
type ListFiltering struct {
	// PromotionID 待过滤的广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// ProjectID 待过滤的广告ID
	ProjectID uint64 `json:"project_id,omitempty"`
}

func (f ListFiltering) GetID() uint64 {
	return f.PromotionID
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filter, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filter))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
