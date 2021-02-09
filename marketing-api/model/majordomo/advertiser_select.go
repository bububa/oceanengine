package majordomo

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type AdvertiserSelectRequest struct {
	AdvertiserID uint64 `json:"advertiser_id,omitempty"` // 账户管家id
}

func (r AdvertiserSelectRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}

type AdvertiserSelectResponse struct {
	model.BaseResponse
	Data *AdvertiserSelectResponseData `json:"data,omitempty"`
}

type AdvertiserSelectResponseData struct {
	List []AdvertiserSelectResponseList `json:"list,omitempty"`
}

type AdvertiserSelectResponseList struct {
	AdvertiserID   uint64 `json:"advertiser_id,omitempty"`
	AdvertiserName string `json:"advertiser_name,omitempty"`
}
