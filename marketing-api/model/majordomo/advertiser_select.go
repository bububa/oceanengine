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
	Data *AdvertiserSelectResponseData `json:"data,omitempty"`
}

// AdvertiserSelectResponseData json返回值
type AdvertiserSelectResponseData struct {
	// List 广告主列表
	List []AdvertiserSelectResponseList `json:"list,omitempty"`
}

// AdvertiserSelectResponseList 广告主
type AdvertiserSelectResponseList struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告主名称
	AdvertiserName string `json:"advertiser_name,omitempty"`
}
