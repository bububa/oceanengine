package interestaction

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// InterestCategoryRequest 兴趣类目查询 API Request
type InterestCategoryRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r InterestCategoryRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}

// InterestCategoryResponse 兴趣类目查询 API Response
type InterestCategoryResponse struct {
	model.BaseResponse
	// Data json返回值
	Data []Object `json:"data,omitempty"`
}
