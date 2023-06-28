package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InterestActionInterestKeywordRequest 获取随心推兴趣标签 API Request
type InterestActionInterestKeywordRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r *InterestActionInterestKeywordRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// InterestActionInterestKeywordResponse 获取随心推兴趣标签 API Response
type InterestActionInterestKeywordResponse struct {
	model.BaseResponse
	Data struct {
		List []InterestKeyword `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
