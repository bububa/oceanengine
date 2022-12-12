package interestaction

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InterestKeywordRequest 兴趣关键词查询 API Request
type InterestKeywordRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// QueryWords 关键词
	QueryWords string `json:"query_words,omitempty"`
}

// Encode implement GetRequest interface
func (r InterestKeywordRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("query_words", r.QueryWords)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// InterestKeywordResponse 兴趣关键词查询 API Response
type InterestKeywordResponse struct {
	model.BaseResponse
	// Data json返回值
	Data InterestKeywordResponseData `json:"data,omitempty"`
}

// InterestKeywordResponseData json返回值
type InterestKeywordResponseData struct {
	List []Object `json:"list,omitempty"`
}
