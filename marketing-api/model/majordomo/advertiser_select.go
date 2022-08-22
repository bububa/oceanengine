package majordomo

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// AdvertiserSelectRequest 广告主列表（管家） API Request
type AdvertiserSelectRequest struct {
	// AdvertiserID 账户管家id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AdvertiserSelectRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}

// AdvertiserSelectResponse 广告主列表（管家） API Response
type AdvertiserSelectResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List 广告主列表
		List []Advertiser `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
