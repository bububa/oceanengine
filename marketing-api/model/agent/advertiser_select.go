package agent

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type AdvertiserSelectRequest struct {
	AdvertiserID uint64 `json:"advertiser_id,omitempty"` // 代理商ID
	Page         int    `json:"page,omitempty"`          // 页码.默认值: 1
	PageSize     int    `json:"page_size,omitempty"`     // 页面数据量.默认值: 100
}

func (r AdvertiserSelectRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

type AdvertiserSelectResponse struct {
	model.BaseResponse
	Data *AdvertiserSelectResponseData `json:"data,omitempty"`
}

type AdvertiserSelectResponseData struct {
	List     []uint64        `json:"list,omitempty"`
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
